package lbaas_v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
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
	lbaasMonitor := createLbaasMonitor(t, client, loadBalancerPool.ID)
	th.AssertNoErr(t, err)
	defer deleteLbaasMonitor(t, client, lbaasMonitor.ID)
	th.AssertEquals(t, false, lbaasMonitor.AdminStateUp)
	th.AssertEquals(t, "", lbaasMonitor.DomainName)

	tools.PrintResource(t, lbaasMonitor)

	updateLbaasMonitor(t, client, lbaasMonitor.ID)

	newLbaasMonitor, err := monitors.Get(client, lbaasMonitor.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, newLbaasMonitor.AdminStateUp)
	th.AssertEquals(t, "www.test.com", newLbaasMonitor.DomainName)
	tools.PrintResource(t, newLbaasMonitor)
}

func createLbaasMonitor(t *testing.T, client *golangsdk.ServiceClient, lbaasPoolID string) *monitors.Monitor {
	monitorName := tools.RandomString("create-monitor-", 3)
	t.Logf("Attempting to create LbaasV2 monitor")

	adminState := false
	createOpts := monitors.CreateOpts{
		PoolID:        lbaasPoolID,
		Type:          "HTTP",
		Delay:         15,
		Timeout:       10,
		MaxRetries:    10,
		Name:          monitorName,
		URLPath:       "/status.php",
		ExpectedCodes: "200",
		AdminStateUp:  &adminState,
	}
	lbaasMonitor, err := monitors.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created LbaasV2 monitor: %s", lbaasMonitor.ID)

	return lbaasMonitor
}

func deleteLbaasMonitor(t *testing.T, client *golangsdk.ServiceClient, lbaasMonitorID string) {
	t.Logf("Attempting to delete LbaasV2 monitor: %s", lbaasMonitorID)

	err := monitors.Delete(client, lbaasMonitorID).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("LbaasV2 monitor is deleted: %s", lbaasMonitorID)
}

func updateLbaasMonitor(t *testing.T, client *golangsdk.ServiceClient, lbaasMonitorID string) {
	t.Logf("Attempting to update LbaasV2 monitor")

	monitorNewName := tools.RandomString("update-monitor-", 3)
	adminStateUp := true

	updateOpts := monitors.UpdateOpts{
		Name:         monitorNewName,
		AdminStateUp: &adminStateUp,
		DomainName:   "www.test.com",
	}

	_, err := monitors.Update(client, lbaasMonitorID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("LbaasV2 monitor successfully updated: %s", lbaasMonitorID)
}
