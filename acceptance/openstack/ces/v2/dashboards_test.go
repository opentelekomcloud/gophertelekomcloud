package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/dashboards"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDashboardsCRUD(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create dashboard")
	dashboardId, err := dashboards.Create(client, dashboards.CreateOpts{
		DashboardName: "test-dashboard-acc",
	})
	th.AssertNoErr(t, err)
	t.Logf("Created dashboard: %s", dashboardId)

	t.Log("Attempting to copy dashboard")
	copyDashboardId, err := dashboards.Create(client, dashboards.CreateOpts{
		DashboardName: "test-dashboard-copy-acc",
		DashboardId:   dashboardId,
	})
	th.AssertNoErr(t, err)
	t.Logf("Copied dashboard: %s", copyDashboardId)

	t.Cleanup(func() {
		t.Log("Attempting to batch delete dashboards")
		results, err := dashboards.BatchDelete(client, dashboards.BatchDeleteOpts{
			DashboardIds: []string{dashboardId, copyDashboardId},
		})
		th.AssertNoErr(t, err)
		for _, result := range results {
			t.Logf("Delete result: ID=%s, Status=%s", result.DashboardId, result.RetStatus)
		}
	})

	t.Log("Attempting to list dashboards")
	dashboardsList, err := dashboards.List(client, dashboards.ListOpts{})
	th.AssertNoErr(t, err)
	t.Logf("Found %d dashboards", len(dashboardsList))

	found := 0
	for _, d := range dashboardsList {
		if d.DashboardId == dashboardId || d.DashboardId == copyDashboardId {
			found++
			t.Logf("Dashboard: ID=%s, Name=%s, Creator=%s",
				d.DashboardId, d.DashboardName, d.CreatorName)
		}
	}
	th.AssertEquals(t, found, 2)

	t.Log("Attempting to modify dashboard")
	err = dashboards.Update(client, dashboardId, dashboards.UpdateOpts{
		DashboardName: "test-dashboard-updated-acc",
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify dashboard update")
	dashboardsList, err = dashboards.List(client, dashboards.ListOpts{
		DashboardId: dashboardId,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(dashboardsList), 1)
	th.AssertEquals(t, dashboardsList[0].DashboardName, "test-dashboard-updated-acc")
}
