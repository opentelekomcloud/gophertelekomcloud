package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	VpcV3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCV3Listing(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	if region != "eu-ch2" {
		t.Skip("Currently VPC V3 only works in SWISS region")
	}
	client, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	listOpts := VpcV3.ListOpts{}
	vpcsList, err := VpcV3.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, rtb := range vpcsList {
		tools.PrintResource(t, rtb)
	}
}

func TestVPCV3Lifecycle(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	if region != "eu-ch2" {
		t.Skip("Currently VPC V3 only works in SWISS region")
	}
	clientV1, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	clientV3, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("vpc-acc-", 3)
	createOpts := vpcs.CreateOpts{
		Name:        name,
		Description: "some interesting description",
		CIDR:        "192.168.0.0/16",
	}

	vpc, err := vpcs.Create(clientV1, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpc.EnableSharedSnat, false)
	th.AssertEquals(t, vpc.Description, "some interesting description")
	th.AssertEquals(t, vpc.Name, name)

	t.Cleanup(func() {
		err = vpcs.Delete(clientV1, vpc.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	cidrOpts := VpcV3.CidrOpts{
		Vpc: &VpcV3.AddExtendCidrOption{
			ExtendCidrs: []string{"23.8.0.0/16"}},
	}
	vpcSecCidr, err := VpcV3.AddSecondaryCidr(clientV3, vpc.ID, cidrOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpcSecCidr.SecondaryCidrs[0], "23.8.0.0/16")

	vpcV3Get, err := VpcV3.Get(clientV3, vpc.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpcV3Get.SecondaryCidrs[0], "23.8.0.0/16")

	vpcSecCidrRm, err := VpcV3.RemoveSecondaryCidr(clientV3, vpc.ID, cidrOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(vpcSecCidrRm.SecondaryCidrs), 0)
}
