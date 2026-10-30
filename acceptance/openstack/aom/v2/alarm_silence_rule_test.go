package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/aom/v2/alarm"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSilenceRuleLifecycle(t *testing.T) {
	client, err := clients.NewAOMV2Client()
	th.AssertNoErr(t, err)

	ruleName := tools.RandomString("tf_silence_rule_", 4)

	createOpts := alarm.CreateSilenceRuleOpts{
		Name:     ruleName,
		Desc:     "created by acceptance test",
		Timezone: "UTC",
		Match: [][]alarm.SilenceMatch{
			{
				{
					Key:     "event_severity",
					Operate: "EQUALS",
					Value:   []string{"Info"},
				},
			},
		},
		MuteConfig: alarm.SilenceMuteConfig{
			StartsAt: 0,
			EndsAt:   86399,
			Type:     "DAILY",
		},
	}

	t.Logf("Attempting to create silence rule: %s", ruleName)
	err = alarm.CreateSilenceRule(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete silence rule: %s", ruleName)
		err := alarm.DeleteSilenceRule(client, []alarm.DeleteSilenceRuleName{{Name: ruleName}})
		th.AssertNoErr(t, err)
		t.Logf("Deleted silence rule: %s", ruleName)
	})

	t.Logf("Attempting to list silence rules")
	rules, err := alarm.ListSilenceRules(client)
	th.AssertNoErr(t, err)

	created, found := findSilenceRule(rules, ruleName)
	if !found {
		t.Fatal("created silence rule not found in list")
	}
	th.AssertEquals(t, createOpts.Desc, created.Desc)
	th.AssertEquals(t, createOpts.MuteConfig.Type, created.MuteConfig.Type)
	th.AssertEquals(t, createOpts.MuteConfig.StartsAt, created.MuteConfig.StartsAt)
	th.AssertEquals(t, createOpts.MuteConfig.EndsAt, created.MuteConfig.EndsAt)
	tools.PrintResource(t, created)

	t.Logf("Attempting to update silence rule: %s", ruleName)
	updateOpts := alarm.UpdateSilenceRuleOpts{
		Name:     ruleName,
		Desc:     "updated by acceptance test",
		Timezone: createOpts.Timezone,
		Match:    createOpts.Match,
		MuteConfig: alarm.SilenceMuteConfig{
			StartsAt: createOpts.MuteConfig.StartsAt,
			EndsAt:   createOpts.MuteConfig.EndsAt,
			Type:     createOpts.MuteConfig.Type,
		},
	}
	err = alarm.UpdateSilenceRule(client, updateOpts)
	th.AssertNoErr(t, err)

	rules, err = alarm.ListSilenceRules(client)
	th.AssertNoErr(t, err)

	updated, found := findSilenceRule(rules, ruleName)
	if !found {
		t.Fatal("updated silence rule not found in list")
	}
	th.AssertEquals(t, updateOpts.Desc, updated.Desc)
}

func findSilenceRule(rules []alarm.SilenceRule, name string) (alarm.SilenceRule, bool) {
	for _, rule := range rules {
		if rule.Name == name {
			return rule, true
		}
	}
	return alarm.SilenceRule{}, false
}
