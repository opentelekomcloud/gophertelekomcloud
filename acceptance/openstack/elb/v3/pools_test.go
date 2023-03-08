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
	defer deleteLoadbalancer(t, client, loadbalancerID)

	poolID := createPool(t, client, loadbalancerID)
	defer deletePool(t, client, poolID)

	t.Logf("Attempting to update ELBv3 Pool: %s", poolID)
	poolName := tools.RandomString("update-pool-", 3)
	emptyDescription := ""
	updateOpts := pools.UpdateOpts{
		Name:                     &poolName,
		Description:              &emptyDescription,
		LBMethod:                 "ROUND_ROBIN",
		DeletionProtectionEnable: pointerto.Bool(false),
	}
	_, err = pools.Update(client, poolID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Pool: %s", poolID)

	newPool, err := pools.Get(client, poolID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *updateOpts.Name, newPool.Name)
	th.AssertEquals(t, emptyDescription, newPool.Description)
	th.AssertEquals(t, updateOpts.LBMethod, newPool.LBMethod)
	th.AssertEquals(t, false, newPool.DeletionProtectionEnable)
}
