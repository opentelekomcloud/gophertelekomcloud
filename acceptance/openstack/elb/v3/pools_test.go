package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPoolList(t *testing.T) {
	client, err := clients.NewElbV3Client()
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
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, loadbalancerID) })

	poolID := createPool(t, client, loadbalancerID)
	t.Cleanup(func() {
		deletePool(t, client, poolID)
	})

	t.Logf("Attempting to update ELBv3 Pool: %s", poolID)
	updateOpts := pools.UpdateOpts{
		Name:                           tools.RandomString("update-pool-", 3),
		Description:                    "",
		LbAlgorithm:                    "ROUND_ROBIN",
		MemberDeletionProtectionEnable: pointerto.Bool(false),
	}
	_, err = pools.Update(client, poolID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Pool: %s", poolID)

	newPool, err := pools.Get(client, poolID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newPool.Name)
	th.AssertEquals(t, "", newPool.Description)
	th.AssertEquals(t, updateOpts.LbAlgorithm, newPool.LbAlgorithm)
	th.AssertEquals(t, false, newPool.MemberDeletionProtectionEnable)
}
