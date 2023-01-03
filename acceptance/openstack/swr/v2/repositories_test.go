package v2

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/repositories"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRepositoryWorkflow(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	// setup org
	orgName := fmt.Sprintf("repo-test-%d", tools.RandomInt(0, 0xf))
	dep := dependencies{t: t, client: client}
	dep.createOrganization(orgName)
	t.Cleanup(func() {
		dep.deleteOrganization(orgName)
	})

	repoName := "magic-test-repo"
	createOpts := repositories.CreateOpts{
		Namespace:   orgName,
		Repository:  repoName,
		Category:    "linux",
		Description: "Test linux repository",
		IsPublic:    true,
	}
	err = repositories.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, repositories.Delete(client, orgName, repoName))
	})

	repo, err := repositories.Get(client, orgName, repoName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, repoName, repo.Name)
	th.AssertEquals(t, createOpts.Category, repo.Category)
	th.AssertEquals(t, createOpts.IsPublic, repo.IsPublic)

	found := false
	err = repositories.List(client, repositories.ListOpts{
		Offset: pointerto.Int(0),
		Limit:  4,
	}).EachPage(func(p pagination.Page) (bool, error) {
		rps, err := repositories.ExtractRepositories(p)
		if err != nil {
			return false, err
		}
		for _, r := range rps {
			if r.Name == repoName {
				found = true
				return false, nil
			}
		}
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, found)

	pages, err := repositories.List(client, repositories.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	_, err = repositories.ExtractRepositories(pages)
	th.AssertNoErr(t, err)

	updateOpts := repositories.UpdateOpts{
		Namespace:   orgName,
		Repository:  repoName,
		Category:    "other",
		Description: "Updated description",
		IsPublic:    false,
	}
	err = repositories.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	updated, err := repositories.Get(client, orgName, repoName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Description, updated.Description)
	th.AssertEquals(t, updateOpts.Category, updated.Category)
	th.AssertEquals(t, updateOpts.IsPublic, updated.IsPublic)
}
