package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVpcList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	listOpts := vpcs.ListOpts{}
	vpcList, err := vpcs.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, vpc := range vpcList {
		tools.PrintResource(t, vpc)
	}
}

func TestVpcLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	// Create a vpc
	vpc := createVpc(t, client)
	th.AssertNoErr(t, err)
	defer deleteVpc(t, client, vpc.ID)

	tools.PrintResource(t, vpc)

	updateOpts := &vpcs.UpdateOpts{
		Name: tools.RandomString("acc-vpc-upd-", 3),
	}

	_, err = vpcs.Update(client, vpc.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newVpc, err := vpcs.Get(client, vpc.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newVpc)
}
