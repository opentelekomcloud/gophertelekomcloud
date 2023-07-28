package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	rules "github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/webtamperprotection_rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWebTamperRuleWorkflow(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	pID := prepareAndRemovePolicy(t, client)
	opts := rules.CreateOpts{
		Hostname: "example.com",
		Path:     "/*",
	}
	r, err := rules.Create(client, pID, opts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, pID, r.Id).Err)
	})

	th.AssertEquals(t, r.Path, opts.Path)
	th.AssertEquals(t, r.PolicyID, pID)
}
