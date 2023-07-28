package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	rules "github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/ccattackprotection_rules"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCCAttackProtectionRuleWorkflow(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	pID := prepareAndRemovePolicy(t, client)
	opts := rules.CreateOpts{
		Path:        "/admin*",
		LimitNum:    pointerto.Int(2),
		LimitPeriod: pointerto.Int(30),
		LockTime:    pointerto.Int(1200),
		TagType:     "cookie",
		TagIndex:    "sessionid",
		TagCondition: rules.TagCondition{
			Category: "Referer",
			Contents: []string{"http://www.example.com/path"},
		},
		Action: rules.Action{
			Category: "block",
			Detail: rules.Detail{
				Response: rules.Response{
					ContentType: "application/json",
					Content:     `{"error":"forbidden"}`,
				},
			},
		},
	}
	r, err := rules.Create(client, pID, opts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, rules.Delete(client, pID, r.Id).Err)
	})

	th.AssertEquals(t, r.PolicyID, pID)
}

func prepareAndRemovePolicy(t *testing.T, client *golangsdk.ServiceClient) string {
	p := preparePolicy(t, client)
	t.Cleanup(func() {
		th.AssertNoErr(t, policies.Delete(client, p.Id).Err)
	})
	return p.Id
}
