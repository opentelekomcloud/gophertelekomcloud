package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	rules "github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/datamasking_rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDatamaskingRuleWorkflow(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	pID := prepareAndRemovePolicy(t, client)
	opts := rules.CreateOpts{
		Path:     "/*",
		Category: "params",
		Index:    "name",
	}
	r, err := rules.Create(client, pID, opts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, pID, r.Id).Err)
	})

	th.AssertEquals(t, opts.Path, r.Path)
	th.AssertEquals(t, pID, r.PolicyID)

	uOpts := rules.UpdateOpts{
		Path:     "/test/*",
		Category: opts.Category,
		Index:    opts.Index,
	}
	_, err = rules.Update(client, pID, r.Id, uOpts).Extract()
	th.AssertNoErr(t, err)

	r, err = rules.Get(client, pID, r.Id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, pID, r.PolicyID)
	th.AssertEquals(t, uOpts.Path, r.Path)
}
