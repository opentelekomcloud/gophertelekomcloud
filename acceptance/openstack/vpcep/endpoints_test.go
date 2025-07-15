package vpcep

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/endpoints"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createService(t *testing.T, client *golangsdk.ServiceClient, elbPortID string) string {
	iFalse := false
	createOpts := services.CreateOpts{
		PortID:      elbPortID,
		ServiceName: tools.RandomString("svc-", 5),
		VpcId:       routerID,
		ServerType:  services.ServerTypeLB,
		ServiceType: services.ServiceTypeInterface,
		Ports: []services.PortMapping{
			{
				ClientPort: 80,
				ServerPort: 8080,
			},
		},
		ApprovalEnabled: &iFalse,
	}
	svc, err := services.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = services.WaitForServiceStatus(client, svc.ID, services.StatusAvailable, 30)
	th.AssertNoErr(t, err)
	return svc.ID
}

func TestEndpointLifecycle(t *testing.T) {
	if routerID == "" || networkID == "" || subnetID == "" {
		t.Skip("OS_ROUTER_ID/VPC_ID, OS_SUBNET_ID and OS_NETWORK_ID variables need to be set")
	}
	t.Parallel()
	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)

	elb := createELB(t)
	defer deleteELB(t, elb.ID)

	t.Logf("Attempting to CREATE VPCEP service with port: %s", elb.VipPortID)
	srvID := createService(t, client, elb.VipPortID)
	t.Cleanup(func() {
		t.Logf("Attempting to DELETE VPCEP Service: %s", srvID)
		th.AssertNoErr(t, services.Delete(client, srvID))
		th.AssertNoErr(t, services.WaitForServiceStatus(client, srvID, services.StatusDeleted, 30))
	})

	t.Logf("Attempting to CREATE VPCEP Endpoint for service: %s", srvID)
	opts := endpoints.CreateOpts{
		NetworkID: networkID,
		ServiceID: srvID,
		VpcId:     routerID,
		EnableDNS: true,
		PortIP:    openstack.ValidIP(t, networkID),
		Tags:      []tags.ResourceTag{{Key: "fizz", Value: "buzz"}},
	}
	created, err := endpoints.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, endpoints.StatusCreating, created.Status)

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE VPCEP Endpoint: %s", created.ID)
		th.AssertNoErr(t, endpoints.Delete(client, created.ID))
		th.AssertNoErr(t, endpoints.WaitForEndpointStatus(client, created.ID, "", 30))
	})

	th.AssertNoErr(t, endpoints.WaitForEndpointStatus(client, created.ID, endpoints.StatusAccepted, 30))

	batchUpdate := endpoints.BatchUpdateReq{
		Permissions: []string{
			"iam:domain::698f9bf85ca9437a9b2f41132ab3aa0e",
		},
		Action: "add",
	}
	t.Logf("Attempting to UPDATE VPCEP Endpoint whitelists: %s", created.ServiceID)
	whiteList, err := endpoints.BatchUpdateWhitelist(client, created.ServiceID, batchUpdate)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, whiteList.Permissions[0], batchUpdate.Permissions[0])

	t.Logf("Attempting to OBTAIN VPCEP Endpoint whitelists: %s", created.ServiceID)
	getWhitelist, err := endpoints.GetWhitelist(client, created.ServiceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getWhitelist.Permissions[0].Permission, batchUpdate.Permissions[0])

	t.Logf("Attempting to OBTAIN VPCEP Endpoint: %s", created.ID)
	got, err := endpoints.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, opts.PortIP, got.IP)

	t.Log("Attempting to LIST VPCEP Endpoints")
	eps, err := endpoints.List(client, endpoints.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.ID, eps[0].ID)

	t.Log("Attempting to OBTAIN basic VPCEP Endpoint service information")
	info, err := services.GetBasicInfo(client, services.BasicOpts{Name: "com.t-systems.otc.eu-de.apig"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "interface", info.ServiceType)
}
