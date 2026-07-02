package v1

import (
	"fmt"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/projects"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/providers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/resources"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/versions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnterpriseProjectsListAndGet(t *testing.T) {
	t.Skip("enterprise project tests not for CI")
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	allProjects, err := projects.List(client, projects.ListOpts{})
	th.AssertNoErr(t, err)

	enterpriseProjectID := clients.EnvOS.GetEnv("ENTERPRISE_PROJECT_ID")
	if enterpriseProjectID == "" {
		if len(allProjects.EnterpriseProjects) == 0 {
			t.Skip("no enterprise projects are available")
		}
		enterpriseProjectID = allProjects.EnterpriseProjects[0].ID
	}

	p, err := projects.Get(client, enterpriseProjectID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, enterpriseProjectID, p.ID)
}

func TestEnterpriseProjectsLifecycle(t *testing.T) {
	t.Skip("enterprise project tests not for CI")
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	suffix := fmt.Sprintf("%d", time.Now().UnixNano()%100000)
	createName := "eps-test-" + suffix
	updateName := "eps-test-upd-" + suffix

	// Create
	created, err := projects.Create(client, projects.CreateOpts{
		Name:        createName,
		Description: "Created by gophertelekomcloud acceptance test",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createName, created.Name)
	t.Logf("Created enterprise project: %s (%s)", created.Name, created.ID)

	// Ensure cleanup
	defer func() {
		_ = projects.Action(client, created.ID, projects.ActionOpts{Action: "disable"})
	}()

	// Update
	updated, err := projects.Update(client, created.ID, projects.UpdateOpts{
		Name:        updateName,
		Description: "Updated by acceptance test",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateName, updated.Name)

	// List with name filter
	filtered, err := projects.List(client, projects.ListOpts{
		Name: updateName,
	})
	th.AssertNoErr(t, err)
	if filtered.TotalCount < 1 {
		t.Fatal("expected at least 1 project in filtered list")
	}

	// Disable (Action)
	err = projects.Action(client, created.ID, projects.ActionOpts{Action: "disable"})
	th.AssertNoErr(t, err)
	t.Log("Disabled enterprise project")

	// Verify disabled
	disabled, err := projects.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, disabled.Status)

	// Enable (Action)
	err = projects.Action(client, created.ID, projects.ActionOpts{Action: "enable"})
	th.AssertNoErr(t, err)
	t.Log("Enabled enterprise project")

	// Verify enabled
	enabled, err := projects.Get(client, created.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, enabled.Status)
}

func TestVersions(t *testing.T) {
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	allVersions, err := versions.List(client)
	th.AssertNoErr(t, err)

	if len(allVersions) == 0 {
		t.Fatal("expected at least one API version")
	}
	t.Logf("Found %d API version(s)", len(allVersions))
	for _, v := range allVersions {
		t.Logf("  %s (status: %s)", v.ID, v.Status)
	}
}

func TestProviders(t *testing.T) {
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	allProviders, err := providers.List(client, providers.ListOpts{})
	th.AssertNoErr(t, err)

	if len(allProviders) == 0 {
		t.Fatal("expected at least one provider")
	}
	t.Logf("Found %d provider(s)", len(allProviders))
	for _, p := range allProviders[:3] {
		t.Logf("  %s (%s)", p.Provider, p.ProviderI18nDisplay)
	}
}

func TestResourcesFilter(t *testing.T) {
	t.Skip("enterprise project tests not for CI")
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	projectID := clients.EnvOS.GetEnv("PROJECT_ID")
	if projectID == "" {
		t.Skip("OS_PROJECT_ID is required for resource filtering")
	}

	result, err := resources.Filter(client, "0", resources.FilterOpts{
		Projects:      []string{projectID},
		ResourceTypes: []string{"ecs"},
		Limit:         5,
	})
	th.AssertNoErr(t, err)
	t.Logf("Found %d resource(s) in default project (total: %d)", len(result.Resources), result.TotalCount)
}
