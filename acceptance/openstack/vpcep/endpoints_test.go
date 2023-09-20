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
	createOpts := &services.CreateOpts{
		PortID:      elbPortID,
		ServiceName: tools.RandomString("svc-", 5),
		RouterID:    routerID,
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
	svc, err := services.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = services.WaitForServiceStatus(client, svc.ID, services.StatusAvailable, 30)
	th.AssertNoErr(t, err)
	return svc.ID
}

func deleteService(t *testing.T, client *golangsdk.ServiceClient, id string) {
	th.AssertNoErr(t, services.Delete(client, id).ExtractErr())
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

	srvID := createService(t, client, elb.VipPortID)
	defer deleteService(t, client, srvID)

	opts := endpoints.CreateOpts{
		NetworkID: networkID,
		ServiceID: srvID,
		RouterID:  routerID,
		EnableDNS: true,
		PortIP:    openstack.ValidIP(t, networkID),
		Tags:      []tags.ResourceTag{{Key: "fizz", Value: "buzz"}},
	}
	created, err := endpoints.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, endpoints.StatusCreating, created.Status)

	defer func() {
		th.AssertNoErr(t, endpoints.Delete(client, created.ID).ExtractErr())
	}()

	th.AssertNoErr(t, endpoints.WaitForEndpointStatus(client, created.ID, endpoints.StatusAccepted, 30))

	batchUpdate := endpoints.BatchUpdateReq{
		Permissions: []string{
			"iam:domain::698f9bf85ca9437a9b2f41132ab3aa0e",
		},
		Action: "add",
	}

	whiteList, err := endpoints.BatchUpdateWhitelist(client, created.ServiceID, batchUpdate)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, whiteList.Permissions[0], batchUpdate.Permissions[0])

	getWhitelist, err := endpoints.GetWhitelist(client, created.ServiceID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getWhitelist.Permissions[0].Permission, batchUpdate.Permissions[0])

	got, err := endpoints.Get(client, created.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, opts.PortIP, got.IP)

	pages, err := endpoints.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	eps, err := endpoints.ExtractEndpoints(pages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, created.ID, eps[0].ID)
}
