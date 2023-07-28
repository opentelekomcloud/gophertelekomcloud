package vpcep

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
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
	th.AssertNoErr(t, loadbalancers.Delete(client, id).Err)
}

func TestListPublicServices(t *testing.T) {
	t.Parallel()

	client, err := clients.NewVPCEndpointV1Client()
	th.AssertNoErr(t, err)

	pages, err := services.ListPublic(client, nil).AllPages()
	th.AssertNoErr(t, err)

	public, err := services.ExtractPublicServices(pages)
	th.AssertNoErr(t, err)
	if len(public) == 0 {
		t.Fatal("Empty public service list")
	}
	th.AssertEquals(t, "OTC", public[0].Owner)
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

	createOpts := &services.CreateOpts{
		PortID:      elb.VipPortID,
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
	}
	svc, err := services.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	err = services.WaitForServiceStatus(client, svc.ID, services.StatusAvailable, 30)
	th.AssertNoErr(t, err)

	defer func() {
		err := services.Delete(client, svc.ID).Err
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, services.WaitForServiceStatus(client, svc.ID, services.StatusDeleted, 30))
	}()

	pages, err := services.List(client, &services.ListOpts{
		ID: svc.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	svcs, err := services.ExtractServices(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(svcs))
	th.AssertEquals(t, svc.ID, svcs[0].ID)

	got, err := services.Get(client, svc.ID).Extract()
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
	updated, err := services.Update(client, svc.ID, uOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, fmt.Sprintf("%s.%s.%s", client.RegionID, uOpts.ServiceName, svc.ID), updated.ServiceName)
}
