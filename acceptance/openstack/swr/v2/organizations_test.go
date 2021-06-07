package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestOrganizationWorkflow(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	name := "test-org"
	opts := organizations.CreateOpts{Namespace: name}
	err = organizations.Create(client, opts).ExtractErr()
	th.AssertNoErr(t, err)
	defer func() {
		th.AssertNoErr(t, organizations.Delete(client, name).ExtractErr())
	}()

	org, err := organizations.Get(client, name).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, org.Name)

	pages, err := organizations.List(client, nil).AllPages()
	th.AssertNoErr(t, err)
	orgs, err := organizations.ExtractOrganizations(pages)
	th.AssertNoErr(t, err)
	found := false
	for _, o := range orgs {
		if o.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("can't find organization '%s' in the list", name)
	}
}
