package v2

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/repositories"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRepositoryWorkflow(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	// setup org
	orgName := fmt.Sprintf("repo-test-%d", tools.RandomInt(0, 0xf))
	err = organizations.Create(client, organizations.CreateOpts{Namespace: orgName}).ExtractErr()
	th.AssertNoErr(t, err)
	defer func() {
		th.AssertNoErr(t, organizations.Delete(client, orgName).ExtractErr())
	}()
	//

	repoName := "magic-test-repo"
	createOpts := repositories.CreateOpts{
		Repository:  repoName,
		Category:    "linux",
		Description: "Test linux repository",
		IsPublic:    true,
	}
	err = repositories.Create(client, orgName, createOpts).ExtractErr()
	th.AssertNoErr(t, err)

	defer func() {
		th.AssertNoErr(t, repositories.Delete(client, orgName, repoName).ExtractErr())
	}()

	repo, err := repositories.Get(client, orgName, repoName).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, repoName, repo.Name)
	th.AssertEquals(t, createOpts.Category, repo.Category)
	th.AssertEquals(t, createOpts.IsPublic, repo.IsPublic)

	zero := 0

	found := false
	err = repositories.List(client, repositories.ListOpts{
		Offset: &zero,
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

	pages, err := repositories.List(client, nil).AllPages()
	th.AssertNoErr(t, err)
	_, err = repositories.ExtractRepositories(pages)
	th.AssertNoErr(t, err)

	updateOpts := repositories.UpdateOpts{
		Category:    "other",
		Description: "Updated description",
		IsPublic:    false,
	}
	err = repositories.Update(client, orgName, repoName, updateOpts).ExtractErr()
	th.AssertNoErr(t, err)
	updated, err := repositories.Get(client, orgName, repoName).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Description, updated.Description)
	th.AssertEquals(t, updateOpts.Category, updated.Category)
	th.AssertEquals(t, updateOpts.IsPublic, updated.IsPublic)
}
