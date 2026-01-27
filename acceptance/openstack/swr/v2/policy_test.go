package v2

import (
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/policy"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/repositories"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	// setup org
	orgName := "test-acc-org-swr-pol"
	dep := dependencies{t: t, client: client}
	dep.createOrganization(orgName)
	t.Cleanup(func() {
		dep.deleteOrganization(orgName)
	})

	repoName := "test-acc-repo-swr-pol"
	dep.createRepository(orgName, repoName)
	t.Cleanup(func() {
		th.AssertNoErr(t, repositories.Delete(client, orgName, repoName))
	})

	createOpts := policy.CreateOpts{
		Algorithm: "or",
		Rules: []policy.Rule{
			{
				Template: "date_rule",
				Params: map[string]string{
					"days": "30",
				},
				TagSelectors: []policy.TagSelector{
					{
						Kind:    "label",
						Pattern: "v5",
					},
				},
			},
		},
	}

	id, err := policy.Create(client, orgName, repoName, createOpts)
	th.AssertNoErr(t, err)
	policyId := strconv.Itoa(id)
	t.Cleanup(func() {
		th.AssertNoErr(t, policy.Delete(client, orgName, repoName, policyId))
	})

	retentionPolicy, err := policy.Get(client, orgName, repoName, policyId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "date_rule", retentionPolicy.Rules[0].Template)

	found := false
	policies, err := policy.List(client, orgName, repoName)
	for _, pol := range policies {
		if pol.ID == id {
			found = true
		}
	}
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, found)

	updateOpts := policy.UpdateOpts{
		Algorithm: "or",
		Rules: []policy.Rule{
			{
				Template: "date_rule",
				Params: map[string]string{
					"days": "45",
				},
				TagSelectors: []policy.TagSelector{
					{
						Kind:    "label",
						Pattern: "v5",
					},
				},
			},
		},
	}
	err = policy.Update(client, orgName, repoName, policyId, updateOpts)
	th.AssertNoErr(t, err)

	updated, err := policy.Get(client, orgName, repoName, policyId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "45", updated.Rules[0].Params["days"])
}
