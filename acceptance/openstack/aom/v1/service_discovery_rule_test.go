package v1

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/aom/v1/monitoring"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServiceDiscoveryRuleLifecycle(t *testing.T) {
	client, err := clients.NewAOMV1Client()
	th.AssertNoErr(t, err)

	// RandomString may return uppercase characters, but rule names allow
	// only lowercase letters, digits and hyphens.
	ruleName := strings.ToLower(tools.RandomString("tf-svc-disc-", 4))

	enable := true
	createOpts := monitoring.AddOrModifyServiceDiscoveryRuleOpts{
		AppRules: []monitoring.AppRule{
			{
				Enable:    &enable,
				Name:      ruleName,
				EventName: "aom_inventory_rules_event",
				Hostid:    []string{},
				Projectid: client.ProjectID,
				Desc:      "created by acceptance test",
				Spec: monitoring.AppRulesSpec{
					AppType:       "Java",
					DetectLog:     "false",
					IsDefaultRule: "false",
					IsDetect:      "false",
					LogFileFix:    []string{"log", "trace", "out"},
					Priority:      9999,
					DiscoveryRule: []monitoring.DiscoveryRule{
						{
							CheckType:    "cmdLine",
							CheckMode:    "contain",
							CheckContent: []string{"acceptance-test-process"},
						},
					},
					NameRule: monitoring.NameRule{
						AppNameRule: []monitoring.AppNameRule{
							{
								NameType: "str",
								Args:     []string{"-svc"},
							},
						},
						ApplicationNameRule: []monitoring.ApplicationNameRule{
							{
								NameType: "str",
								Args:     []string{"-app"},
							},
						},
					},
				},
			},
		},
	}

	t.Logf("Attempting to create service discovery rule: %s", ruleName)
	createResp, err := monitoring.AddOrModifyServiceDiscoveryRule(client, createOpts)
	th.AssertNoErr(t, err)

	if len(createResp.Results) == 0 {
		t.Fatal("no rule returned in add/modify response")
	}
	ruleId := createResp.Results[0]["id"]
	if ruleId == "" {
		t.Fatal("created rule has no ID in add/modify response")
	}

	t.Cleanup(func() {
		t.Logf("Attempting to delete service discovery rule: %s (%s)", ruleName, ruleId)
		_, err := monitoring.DeleteServiceDiscoveryRule(client, []string{ruleId})
		th.AssertNoErr(t, err)
		t.Logf("Deleted service discovery rule: %s", ruleName)
	})

	t.Logf("Attempting to list all service discovery rules")
	rules, err := monitoring.ListServiceDiscoveryRules(client, monitoring.ListServiceDiscoveryRulesOpts{})
	th.AssertNoErr(t, err)

	created, found := findServiceDiscoveryRule(rules, ruleName)
	if !found {
		t.Fatal("created service discovery rule not found in list")
	}
	th.AssertEquals(t, ruleName, created.Name)
	th.AssertEquals(t, "created by acceptance test", created.Desc)
	th.AssertEquals(t, true, *created.Enable)
	th.AssertEquals(t, createOpts.AppRules[0].Spec.Priority, created.Spec.Priority)
	th.AssertEquals(t, createOpts.AppRules[0].Spec.DetectLog, created.Spec.DetectLog)

	t.Logf("Attempting to list service discovery rule by ID: %s", ruleId)
	rulesById, err := monitoring.ListServiceDiscoveryRules(client, monitoring.ListServiceDiscoveryRulesOpts{Id: ruleId})
	th.AssertNoErr(t, err)
	if len(rulesById) != 1 {
		t.Fatalf("expected 1 rule when filtering by ID, got %d", len(rulesById))
	}
	th.AssertEquals(t, ruleId, rulesById[0].Id)

	t.Logf("Attempting to update service discovery rule: %s", ruleName)
	updateOpts := createOpts
	updateOpts.AppRules[0].Id = ruleId
	updateOpts.AppRules[0].CreateTime = created.CreateTime
	updateOpts.AppRules[0].Desc = "updated by acceptance test"
	updateResp, err := monitoring.AddOrModifyServiceDiscoveryRule(client, updateOpts)
	th.AssertNoErr(t, err)
	if len(updateResp.Results) == 0 {
		t.Fatal("no rule returned in update response")
	}

	rules, err = monitoring.ListServiceDiscoveryRules(client, monitoring.ListServiceDiscoveryRulesOpts{Id: ruleId})
	th.AssertNoErr(t, err)

	updated, found := findServiceDiscoveryRule(rules, ruleName)
	if !found {
		t.Fatal("updated service discovery rule not found in list")
	}
	th.AssertEquals(t, "updated by acceptance test", updated.Desc)
}

func findServiceDiscoveryRule(rules []monitoring.AppRule, name string) (monitoring.AppRule, bool) {
	for _, rule := range rules {
		if rule.Name == name {
			return rule, true
		}
	}
	return monitoring.AppRule{}, false
}
