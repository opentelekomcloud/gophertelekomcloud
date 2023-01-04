package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/domains"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAccessDomainWorkflow(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	// setup deps
	orgName := "domaintest"
	dep := dependencies{t: t, client: client}
	dep.createOrganization(orgName)
	t.Cleanup(func() { dep.deleteOrganization(orgName) })

	repoName := "repodomain"
	dep.createRepository(orgName, repoName)
	t.Cleanup(func() { dep.deleteRepository(orgName, repoName) })

	domainToShare := clients.EnvOS.GetEnv("DOMAIN_NAME_2")
	if domainToShare == "" {
		t.Skipf("OS_DOMAIN_NAME_2 env var is missing but SWR domain test requires it")
	}
	opts := domains.CreateOpts{
		Namespace:    orgName,
		Repository:   repoName,
		AccessDomain: domainToShare,
		Permit:       "read",
		Deadline:     "forever",
	}
	th.AssertNoErr(t, domains.Create(client, opts))

	t.Cleanup(func() {
		err = domains.Delete(client, domains.GetOpts{
			Namespace:    orgName,
			Repository:   repoName,
			AccessDomain: domainToShare,
		})
		th.AssertNoErr(t, err)
	})

	accessDomains, err := domains.List(client, domains.ListOpts{
		Namespace:  orgName,
		Repository: repoName,
	})
	th.AssertNoErr(t, err)

	found := false
	for _, d := range accessDomains {
		if d.AccessDomain == domainToShare {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("access domain %s is not found in the list", domainToShare)
	}

	updateOpts := domains.UpdateOpts{
		Namespace:    orgName,
		Repository:   repoName,
		AccessDomain: domainToShare,
		Permit:       "read", // only read premission is possible
		Deadline:     "2024-01-01T00:00:00.000Z",
		Description:  "Updated description",
	}
	th.AssertNoErr(t, domains.Update(client, updateOpts))

	domain, err := domains.Get(client, domains.GetOpts{
		Namespace:    orgName,
		Repository:   repoName,
		AccessDomain: domainToShare,
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, domainToShare, domain.AccessDomain)
	th.CheckEquals(t, updateOpts.Permit, domain.Permit)
	th.CheckEquals(t, true, domain.Status)
	th.CheckEquals(t, updateOpts.Description, domain.Description)
}
