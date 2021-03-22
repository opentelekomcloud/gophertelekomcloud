package lbaas_v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLbaasV2MonitorsList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	listOpts := monitors.ListOpts{}
	allPages, err := monitors.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	lbaasMonitors, err := monitors.ExtractMonitors(allPages)
	th.AssertNoErr(t, err)

	for _, monitor := range lbaasMonitors {
		tools.PrintResource(t, monitor)
	}
}

func TestLbaasV2MonitorLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create lbaasV2 Load Balancer
	loadBalancer, err := createLbaasLoadBalancer(t, client)
	th.AssertNoErr(t, err)
	defer deleteLbaasLoadBalancer(t, client, loadBalancer.ID)

	// Create lbaasV2 pool
	loadBalancerPool, err := createLbaasPool(t, client, loadBalancer.ID)
	th.AssertNoErr(t, err)
	defer deleteLbaasPool(t, client, loadBalancerPool.ID)

	// Create lbaasV2 monitor
	lbaasMonitor, err := createLbaasMonitor(t, client, loadBalancerPool.ID)
	th.AssertNoErr(t, err)
	defer deleteLbaasMonitor(t, client, lbaasMonitor.ID)

	tools.PrintResource(t, lbaasMonitor)

	err = updateLbaasMonitor(t, client, lbaasMonitor.ID)
	th.AssertNoErr(t, err)

	newLbaasMonitor, err := monitors.Get(client, lbaasMonitor.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newLbaasMonitor)
}
