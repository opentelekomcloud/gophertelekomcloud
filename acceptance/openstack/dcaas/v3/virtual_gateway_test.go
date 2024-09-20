package v3

import (
	"os"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcaas/v3/virtual-gateway"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVirtualGatewayListing(t *testing.T) {
	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	opts := virtual_gateway.ListOpts{}
	vgwList, err := virtual_gateway.List(client, opts)
	th.AssertNoErr(t, err)

	for _, vgw := range vgwList {
		tools.PrintResource(t, vgw)
	}
}

func TestVirtualGatewayLifecycle(t *testing.T) {
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID necessary for this test")
	}
	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	clientNet, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)
	vpc, err := vpcs.Get(clientNet, vpcID).Extract()
	th.AssertNoErr(t, err)
	// Create a virtual gateway
	name := strings.ToLower(tools.RandomString("acc-virtual-gateway-v3-", 5))
	createOpts := virtual_gateway.CreateOpts{
		Name:         name,
		VpcId:        vpcID,
		Description:  "acc-virtual-gateway-v3",
		LocalEpGroup: []string{vpc.CIDR},
		Tags: []tags.ResourceTag{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		},
	}

	created, err := virtual_gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	vgw, err := virtual_gateway.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, vgw.Name)

	nameUpdated := strings.ToLower(tools.RandomString("acc-virtual-gateway-v3-up", 5))
	updateOpts := virtual_gateway.UpdateOpts{
		Name:        nameUpdated,
		Description: pointerto.String("acc-virtual-gateway-v3-updated"),
	}
	updated, err := virtual_gateway.Update(client, created.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, nameUpdated, updated.Name)

	opts := virtual_gateway.ListOpts{}
	gateways, err := virtual_gateway.List(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(gateways))

	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
