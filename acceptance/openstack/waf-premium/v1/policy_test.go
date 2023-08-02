package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWafPremiumPolicyWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)
	th.AssertEquals(t, policy.Action.Category, "log")
	th.AssertEquals(t, policy.Level, 2)
	th.AssertEquals(t, *policy.Options.WebAttack, true)
	th.AssertEquals(t, *policy.Options.Common, true)
	th.AssertEquals(t, *policy.Options.AntiCrawler, false)
	th.AssertEquals(t, *policy.Options.CrawlerEngine, false)
	th.AssertEquals(t, *policy.Options.CrawlerScript, false)
	th.AssertEquals(t, *policy.Options.CrawlerOther, false)
	th.AssertEquals(t, *policy.Options.WebShell, false)
	th.AssertEquals(t, *policy.Options.Cc, true)
	th.AssertEquals(t, *policy.Options.Custom, true)
	th.AssertEquals(t, *policy.Options.WhiteblackIp, true)
	th.AssertEquals(t, *policy.Options.GeoIp, true)
	th.AssertEquals(t, *policy.Options.Ignore, true)
	th.AssertEquals(t, *policy.Options.Privacy, true)
	th.AssertEquals(t, *policy.Options.AntiTamper, true)
	th.AssertEquals(t, *policy.Options.AntiLeakage, false)
	th.AssertEquals(t, *policy.Options.FollowedAction, false)
	th.AssertEquals(t, *policy.Options.BotEnable, true)
	th.AssertEquals(t, *policy.Options.Crawler, true)
	th.AssertEquals(t, *policy.Options.Precise, false)

	// Not supported in SWISS
	// th.AssertEquals(t, *policy.Options.ModulexEnabled, false)
	// th.AssertEquals(t, *policy.ModulexOptions.GlobalRateEnabled, true)
	// th.AssertEquals(t, policy.ModulexOptions.GlobalRateMode, "log")
	// th.AssertEquals(t, *policy.ModulexOptions.PreciseRulesEnabled, true)
	// th.AssertEquals(t, policy.ModulexOptions.PreciseRulesMode, "log")
	// th.AssertEquals(t, policy.ModulexOptions.PreciseRulesManagedMode, "auto")
	// th.AssertEquals(t, policy.ModulexOptions.PreciseRulesAgingMode, "auto")
	// th.AssertEquals(t, policy.ModulexOptions.PreciseRulesRetention, 600)
	// th.AssertEquals(t, *policy.ModulexOptions.CcRulesEnabled, true)
	// th.AssertEquals(t, policy.ModulexOptions.CcRulesMode, "log")
	// th.AssertEquals(t, policy.ModulexOptions.CcRulesManagedMode, "auto")
	// th.AssertEquals(t, policy.ModulexOptions.CcRulesAgingMode, "manual")
	// th.AssertEquals(t, policy.ModulexOptions.CcRulesRetention, 600)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	t.Logf("Attempting to List WAF Premium policy")
	certificatesList, err := policies.List(client, policies.ListOpts{
		Name: policyName,
	})
	th.AssertNoErr(t, err)

	if len(certificatesList) < 1 {
		t.Fatal("empty WAF Premium policy list")
	}

	t.Logf("Attempting to Get WAF Premium policy: %s", policy.ID)
	p, err := policies.Get(client, policy.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, p.ID, policy.ID)

	t.Logf("Attempting to Update WAF Premium policy: %s", policy.ID)
	updateOpts := policies.UpdateOpts{
		Name: policyName + "-updated",
		Action: &policies.PolicyAction{
			Category: "block",
		},
		Options: &policies.PolicyOption{
			WebAttack: pointerto.Bool(false),
		},
		Level:         1,
		FullDetection: pointerto.Bool(true),
	}
	updated, err := policies.Update(client, policy.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updated.Level, 1)
	th.AssertEquals(t, *updated.FullDetection, true)
	th.AssertEquals(t, updated.Action.Category, "block")
	th.AssertEquals(t, *updated.Options.WebAttack, false)
}
