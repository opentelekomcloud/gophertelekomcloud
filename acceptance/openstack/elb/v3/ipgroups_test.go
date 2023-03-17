package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/ipgroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestIpGroupList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	listOpts := ipgroups.ListOpts{}
	ipgroupsPages, err := ipgroups.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	ipgroupsList, err := ipgroups.ExtractIpGroups(ipgroupsPages)
	th.AssertNoErr(t, err)

	for _, lb := range ipgroupsList {
		tools.PrintResource(t, lb)
	}
}
