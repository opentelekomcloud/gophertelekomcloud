package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
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
	t.Cleanup(func() {
		deleteLoadbalancer(t, client, loadbalancerID)
	})

	poolID := createPool(t, client, loadbalancerID)
	t.Cleanup(func() { deletePool(t, client, poolID) })

	t.Logf("Attempting to Create ELBv3 Monitor")
	monitorName := tools.RandomString("create-monitor-", 3)
	createOpts := monitors.CreateOpts{
		PoolId:        poolID,
		Type:          "HTTP",
		Delay:         pointerto.Int(1),
		Timeout:       pointerto.Int(30),
		MaxRetries:    pointerto.Int(3),
		HttpMethod:    "OPTIONS",
		ExpectedCodes: "200-299",
		Name:          monitorName,
	}
	monitor, err := monitors.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to Delete ELBv3 Monitor: %s", monitor.Id)
		err := monitors.Delete(client, monitor.Id)
		th.AssertNoErr(t, err)
		t.Logf("Deleted ELBv3 Monitor: %s", monitor.Id)
	})

	th.AssertEquals(t, createOpts.Name, monitor.Name)
	th.AssertEquals(t, createOpts.Type, monitor.Type)
	th.AssertEquals(t, createOpts.MaxRetries, monitor.MaxRetries)
	t.Logf("Created ELBv3 Monitor: %s", monitor.Id)

	t.Logf("Attempting to Update ELBv3 Monitor")
	monitorName = tools.RandomString("update-monitor-", 3)
	updateOpts := monitors.UpdateOpts{
		Delay:          pointerto.Int(3),
		Timeout:        pointerto.Int(35),
		MaxRetries:     pointerto.Int(5),
		MaxRetriesDown: pointerto.Int(3),
		Name:           monitorName,
	}
	_, err = monitors.Update(client, monitor.Id, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 Monitor: %s", monitor.Id)

	newMonitor, err := monitors.Get(client, monitor.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newMonitor.Name)
	th.AssertEquals(t, updateOpts.Timeout, newMonitor.Timeout)
	th.AssertEquals(t, updateOpts.Delay, newMonitor.Delay)
	th.AssertEquals(t, updateOpts.MaxRetries, newMonitor.MaxRetries)
	th.AssertEquals(t, createOpts.HttpMethod, newMonitor.HttpMethod)
}
