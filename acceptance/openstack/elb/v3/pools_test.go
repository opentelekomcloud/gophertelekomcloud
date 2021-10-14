package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPoolList(t *testing.T) {
	client, err := clients.NewElbV3Client(t)
	th.AssertNoErr(t, err)

	listOpts := pools.ListOpts{}
	poolPages, err := pools.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	poolList, err := pools.ExtractPools(poolPages)
	th.AssertNoErr(t, err)

	for _, pool := range poolList {
		tools.PrintResource(t, pool)
	}
}

func TestPoolLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client(t)
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

}
