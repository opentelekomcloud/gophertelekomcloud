package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLoadBalancerList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	listOpts := loadbalancers.ListOpts{}
	loadbalancerPages, err := loadbalancers.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	loadbalancerList, err := loadbalancers.ExtractLoadbalancers(loadbalancerPages)
	th.AssertNoErr(t, err)

	for _, lb := range loadbalancerList {
		tools.PrintResource(t, lb)
	}
}

func TestLoadBalancerLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	t.Logf("Attempting to update ELBv3 LoadBalancer: %s", loadbalancerID)
	lbName := tools.RandomString("update-lb-", 3)
	emptyDescription := ""
	updateOpts := loadbalancers.UpdateOpts{
		Name:        lbName,
		Description: &emptyDescription,
	}

	_, err = loadbalancers.Update(client, loadbalancerID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 LoadBalancer: %s", loadbalancerID)

	newLoadbalancer, err := loadbalancers.Get(client, loadbalancerID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newLoadbalancer.Name)
	th.AssertEquals(t, emptyDescription, newLoadbalancer.Description)
}
