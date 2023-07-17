//go:build acceptance
// +build acceptance

package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/projects"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestProjectsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	listOpts := projects.ListOpts{
		Name: os.Getenv("OS_PROJECT_NAME"),
	}

	allPages, err := projects.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	for _, project := range allProjects {
		tools.PrintResource(t, project)
	}
}

func TestProjectsGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)
	opts := projects.ListOpts{
		Name: os.Getenv("OS_PROJECT_NAME"),
	}
	allPages, err := projects.List(client, opts).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	project := allProjects[0]
	p, err := projects.Get(client, project.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, p)
}

func TestProjectsCRUD(t *testing.T) {
	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	project, err := CreateProject(t, client, nil)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteProject(t, client, project.ID)
	})

	tools.PrintResource(t, project)

	updateOpts := projects.UpdateOpts{
		Description: "Updated",
	}

	updatedProject, err := projects.Update(client, project.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedProject)
}

func TestProjectsNested(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	projectMain, err := CreateProject(t, client, nil)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteProject(t, client, projectMain.ID)
	})
	tools.PrintResource(t, projectMain)

	createOpts := projects.CreateOpts{
		ParentID: projectMain.ID,
	}

	project, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteProject(t, client, projectMain.ID)
	})

	tools.PrintResource(t, project)
}
