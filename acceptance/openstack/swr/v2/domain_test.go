package v2

import (
	"strings"
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
	defer dep.deleteOrganization(orgName)

	repoName := "repodomain"
	dep.createRepository(orgName, repoName)
	defer dep.deleteRepository(orgName, repoName)

	domainToShare := clients.EnvOS.GetEnv("DOMAIN_NAME_2")
	if domainToShare == "" {
		t.Skipf("OS_DOMAIN_NAME_2 env var is missing but SWR domain test requires it")
	}
	opts := domains.CreateOpts{
		AccessDomain: domainToShare,
		Permit:       "read",
		Deadline:     "forever",
	}
	th.AssertNoErr(t, domains.Create(client, orgName, repoName, opts).ExtractErr())

	defer func() {
		err = domains.Delete(client, orgName, repoName, domainToShare).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	pages, err := domains.List(client, orgName, repoName).AllPages()
	th.AssertNoErr(t, err)

	accessDomains, err := domains.ExtractAccessDomains(pages)
	th.AssertNoErr(t, err)

	found := false
	for _, d := range accessDomains {
		if d.AccessDomain == strings.ToLower(domainToShare) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("access domain %s is not found in the list", domainToShare)
	}

	updateOpts := domains.UpdateOpts{
		Permit:      "read", // only read premission is possible
		Deadline:    "2022-01-01T00:00:00.000Z",
		Description: "Updated description",
	}
	th.AssertNoErr(t, domains.Update(client, orgName, repoName, domainToShare, updateOpts).ExtractErr())

	domain, err := domains.Get(client, orgName, repoName, domainToShare).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, strings.ToLower(domainToShare), domain.AccessDomain)
	th.CheckEquals(t, updateOpts.Permit, domain.Permit)
	th.CheckEquals(t, true, domain.Status)
	th.CheckEquals(t, updateOpts.Description, domain.Description)
}
