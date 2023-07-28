package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/monitors"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMonitorList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	listOpts := monitors.ListOpts{}
	monitorPages, err := monitors.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	monitorList, err := monitors.ExtractMonitors(monitorPages)
	th.AssertNoErr(t, err)

	for _, monitor := range monitorList {
		tools.PrintResource(t, monitor)
	}
}

func TestMonitorLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	loadbalancerID := createLoadBalancer(t, client)
	defer deleteLoadbalancer(t, client, loadbalancerID)

	poolID := createPool(t, client, loadbalancerID)
	defer deletePool(t, client, poolID)

	t.Logf("Attempting to Create ELBv3 Monitor")
	monitorName := tools.RandomString("create-monitor-", 3)
	createOpts := monitors.CreateOpts{
		PoolID:        poolID,
		Type:          monitors.TypeHTTP,
		Delay:         1,
		Timeout:       30,
		MaxRetries:    3,
		HTTPMethod:    "OPTIONS",
		ExpectedCodes: "200-299",
		Name:          monitorName,
	}
	monitor, err := monitors.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		t.Logf("Attempting to Delete ELBv3 Monitor: %s", monitor.ID)
		err := monitors.Delete(client, monitor.ID).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Monitor: %s", monitor.ID)
	}()

	th.AssertEquals(t, createOpts.Name, monitor.Name)
	th.AssertEquals(t, createOpts.Type, monitor.Type)
	th.AssertEquals(t, createOpts.MaxRetries, monitor.MaxRetries)
	t.Logf("Created ELBv3 Monitor: %s", monitor.ID)

	t.Logf("Attempting to Update ELBv3 Monitor")
	monitorName = tools.RandomString("update-monitor-", 3)
	updateOpts := monitors.UpdateOpts{
		Delay:          3,
		Timeout:        35,
		MaxRetries:     5,
		MaxRetriesDown: 3,
		Name:           monitorName,
	}
	_, err = monitors.Update(client, monitor.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Monitor: %s", monitor.ID)

	newMonitor, err := monitors.Get(client, monitor.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newMonitor.Name)
	th.AssertEquals(t, updateOpts.Timeout, newMonitor.Timeout)
	th.AssertEquals(t, updateOpts.Delay, newMonitor.Delay)
	th.AssertEquals(t, updateOpts.MaxRetries, newMonitor.MaxRetries)
	th.AssertEquals(t, createOpts.HTTPMethod, newMonitor.HTTPMethod)
}
