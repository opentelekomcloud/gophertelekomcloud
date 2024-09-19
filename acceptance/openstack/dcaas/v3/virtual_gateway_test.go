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
	if os.Getenv("OS_VPC_ID") == "" {
		t.Skip("OS_VPC_ID necessary for this test")
	}
	client, err := clients.NewDCaaSV3Client()
	th.AssertNoErr(t, err)

	// Create a virtual gateway
	name := strings.ToLower(tools.RandomString("test-virtual-gateway", 5))
	createOpts := virtual_gateway.CreateOpts{
		Name:         name,
		VpcId:        os.Getenv("OS_VPC_ID"),
		Description:  "test-virtual-gateway-v3",
		LocalEpGroup: []string{""},
		Tags: []tags.ResourceTag{
			{
				Key:   "TestKey",
				Value: "TestValue",
			},
		},
	}

	created, err := virtual_gateway.Create(client, createOpts)
	th.AssertNoErr(t, err)

	_, err = virtual_gateway.Get(client, created.ID)
	th.AssertNoErr(t, err)

	updateOpts := virtual_gateway.UpdateOpts{
		Name:        tools.RandomString(name, 3),
		Description: pointerto.String("test-virtual-gateway-v3-updated"),
	}
	_ = virtual_gateway.Update(client, created.ID, updateOpts)
	th.AssertNoErr(t, err)

	opts := virtual_gateway.ListOpts{}
	_, err = virtual_gateway.List(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = virtual_gateway.Delete(client, created.ID)
		th.AssertNoErr(t, err)
	})
}
