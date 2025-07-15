package vpcep

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/endpoints"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/quota"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var (
	routerID  = clients.EnvOS.GetEnv("ROUTER_ID", "VPC_ID")
	networkID = clients.EnvOS.GetEnv("NETWORK_ID")
	subnetID  = clients.EnvOS.GetEnv("SUBNET_ID")
)

func createELB(t *testing.T) *loadbalancers.LoadBalancer {
	client, err := clients.NewElbV2Client()
	th.AssertNoErr(t, err)
	t.Log("Attempting to CREATE ELB for service")
	lb, err := loadbalancers.Create(client, loadbalancers.CreateOpts{
		Name:        tools.RandomString("svc-lb-", 5),
		VipSubnetID: subnetID,
	}).Extract()
	th.AssertNoErr(t, err)
	return lb
}

func deleteELB(t *testing.T, id string) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)
	t.Logf("Attempting to DELETE ELB instance: %s", id)
	th.AssertNoErr(t, loadbalancers.Delete(client, id).ExtractErr())
}

func TestListPublicServices(t *testing.T) {
	t.Parallel()

	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)
	t.Log("Attempting to LIST Public services")
	public, err := services.ListPublic(client, services.ListPublicOpts{})
	th.AssertNoErr(t, err)
	if len(public) == 0 {
		t.Fatal("Empty public service list")
	}
	th.AssertEquals(t, "OTC", public[0].Owner)
}

func TestGetQuota(t *testing.T) {
	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)
	t.Log("Attempting to Get VPCEP endpoint_service quota")
	q, err := quota.Get(client, "endpoint_service")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "endpoint_service", q.Quotas.Resources[0].Type)
}

func TestServicesWorkflow(t *testing.T) {
	if routerID == "" || networkID == "" || subnetID == "" {
		t.Skip("OS_ROUTER_ID/VPC_ID, OS_SUBNET_ID and OS_NETWORK_ID variables need to be set")
	}

	t.Parallel()

	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)

	elb := createELB(t)
	defer deleteELB(t, elb.ID)

	createOpts := services.CreateOpts{
		PortID:      elb.VipPortID,
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
	}
	t.Log("Attempting to CREATE VPCEP service")
	svc, err := services.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = services.WaitForServiceStatus(client, svc.ID, services.StatusAvailable, 30)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE VPCEP Endpoint service: %s", svc.ID)
		th.AssertNoErr(t, services.Delete(client, svc.ID))
		th.AssertNoErr(t, services.WaitForServiceStatus(client, svc.ID, services.StatusDeleted, 30))
	})

	t.Log("Attempting to LIST VPCEP Endpoint services")
	svcs, err := services.List(client, services.ListOpts{
		ID: svc.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(svcs))
	th.AssertEquals(t, svc.ID, svcs[0].ID)

	t.Logf("Attempting to OBTAIN VPCEP Endpoint service: %s", svc.ID)
	got, err := services.Get(client, svc.ID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, svc.VIPPortID, got.VIPPortID)
	th.AssertEquals(t, svc.ApprovalEnabled, got.ApprovalEnabled)
	th.AssertEquals(t, svc.CreatedAt, got.CreatedAt)
	th.AssertEquals(t, 0, svc.ConnectionCount)

	iFalse := false
	uOpts := services.UpdateOpts{
		ApprovalEnabled: &iFalse,
		ServiceName:     tools.RandomString("edited-", 5),
	}
	t.Logf("Attempting to UPDATE VPCEP Endpoint service: %s", svc.ID)
	updated, err := services.Update(client, svc.ID, uOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, fmt.Sprintf("%s.%s.%s", client.RegionID, uOpts.ServiceName, svc.ID), updated.ServiceName)

	t.Logf("Attempting to CREATE VPCEP Endpoint for service: %s", svc.ID)
	opts := endpoints.CreateOpts{
		NetworkID: networkID,
		ServiceID: svc.ID,
		VpcId:     routerID,
		EnableDNS: true,
		PortIP:    openstack.ValidIP(t, networkID),
		Tags:      []tags.ResourceTag{{Key: "fizz", Value: "buzz"}},
	}
	ep, err := endpoints.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, endpoints.StatusCreating, ep.Status)
	th.AssertNoErr(t, endpoints.WaitForEndpointStatus(client, ep.ID, endpoints.StatusAccepted, 30))

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE VPCEP Endpoint: %s", ep.ID)
		th.AssertNoErr(t, endpoints.Delete(client, ep.ID))
		th.AssertNoErr(t, endpoints.WaitForEndpointStatus(client, ep.ID, "", 30))
	})

	t.Logf("Attempting to ACCEPT VPCEP Endpoint to service: %s", ep.ID)
	aOpts := services.ActionOpts{
		Action:    "receive",
		Endpoints: []string{ep.ID},
	}
	c, err := services.Action(client, svc.ID, aOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "accepted", c[0].Status)

	t.Logf("Attempting to OBTAIN VPCEP Endpoint service connections: %s", svc.ID)
	connections, err := services.ListConnections(client, svc.ID, services.ListConnectionsOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(connections))

	t.Logf("Attempting to REJECT VPCEP Endpoint to service: %s", ep.ID)
	rejectOpts := services.ActionOpts{
		Action:    "reject",
		Endpoints: []string{ep.ID},
	}
	rej, err := services.Action(client, svc.ID, rejectOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "rejected", rej[0].Status)
}
