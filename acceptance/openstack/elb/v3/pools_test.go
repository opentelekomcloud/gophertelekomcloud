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

	t.Logf("Attempting to create ELBv3 Pool")
	poolName := tools.RandomString("create-pool-", 3)
	createOpts := pools.CreateOpts{
		LBMethod:       "LEAST_CONNECTIONS",
		Protocol:       "HTTPS",
		LoadbalancerID: loadbalancerID,
		Name:           poolName,
		Description:    "some interesting description",
	}

	pool, err := pools.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created ELBv3 Pool: %s", pool.ID)

	defer func() {
		t.Logf("Attempting to delete ELBv3 Pool: %s", pool.ID)
		err := pools.Delete(client, pool.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Pool: %s", pool.ID)
	}()

	th.AssertEquals(t, createOpts.Name, pool.Name)
	th.AssertEquals(t, createOpts.Description, pool.Description)
	th.AssertEquals(t, createOpts.LBMethod, pool.LBMethod)

	t.Logf("Attempting to update ELBv3 Pool: %s", pool.ID)
	poolName = tools.RandomString("update-pool-", 3)
	emptyDescription := ""
	updateOpts := pools.UpdateOpts{
		Name:        poolName,
		Description: &emptyDescription,
		LBMethod:    "ROUND_ROBIN",
	}
	_, err = pools.Update(client, pool.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Pool: %s", pool.ID)

	newPool, err := pools.Get(client, pool.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newPool.Name)
	th.AssertEquals(t, emptyDescription, newPool.Description)
	th.AssertEquals(t, updateOpts.LBMethod, newPool.LBMethod)

}
