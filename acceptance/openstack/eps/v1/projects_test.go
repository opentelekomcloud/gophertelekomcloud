package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/projects"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnterpriseProjectsListAndGet(t *testing.T) {
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	allPages, err := projects.List(client, projects.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	enterpriseProjectID := clients.EnvOS.GetEnv("OS_ENTERPRISE_PROJECT_ID")
	if enterpriseProjectID == "" {
		if len(allProjects.EnterpriseProjects) == 0 {
			t.Skip("no enterprise projects are available")
		}
		enterpriseProjectID = allProjects.EnterpriseProjects[0].ID
	}

	p, err := projects.Get(client, enterpriseProjectID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, enterpriseProjectID, p.ID)
}

func TestEnterpriseProjectsLifecycle(t *testing.T) {
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	// Create
	created, err := projects.Create(client, projects.CreateOpts{
		Name:        "eps-sdk-test",
		Description: "Created by gophertelekomcloud acceptance test",
	}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "eps-sdk-test", created.Name)
	t.Logf("Created enterprise project: %s (%s)", created.Name, created.ID)

	// Update
	updated, err := projects.Update(client, created.ID, projects.UpdateOpts{
		Name:        "eps-sdk-test-updated",
		Description: "Updated by acceptance test",
	}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "eps-sdk-test-updated", updated.Name)

	// List with name filter
	allPages, err := projects.List(client, projects.ListOpts{
		Name: "eps-sdk-test-updated",
	}).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)
	if allProjects.TotalCount < 1 {
		t.Fatal("expected at least 1 project in filtered list")
	}

	// Disable (Action)
	err = projects.Action(client, created.ID, projects.ActionOpts{
		Action: "disable",
	}).ExtractErr()
	th.AssertNoErr(t, err)
	t.Log("Disabled enterprise project")

	// Verify disabled
	disabled, err := projects.Get(client, created.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, disabled.Status)

	// Enable (Action)
	err = projects.Action(client, created.ID, projects.ActionOpts{
		Action: "enable",
	}).ExtractErr()
	th.AssertNoErr(t, err)
	t.Log("Enabled enterprise project")

	// Verify enabled
	enabled, err := projects.Get(client, created.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, enabled.Status)
}
