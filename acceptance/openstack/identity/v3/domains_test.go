//go:build acceptance
// +build acceptance

package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/domains"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDomainsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	var iTrue bool = true
	listOpts := domains.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := domains.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allDomains, err := domains.ExtractDomains(allPages)
	th.AssertNoErr(t, err)

	for _, domain := range allDomains {
		tools.PrintResource(t, domain)
	}
}

func TestDomainsGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := domains.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allDomains, err := domains.ExtractDomains(allPages)
	th.AssertNoErr(t, err)

	domain := allDomains[0]
	p, err := domains.Get(client, domain.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, p)
}
