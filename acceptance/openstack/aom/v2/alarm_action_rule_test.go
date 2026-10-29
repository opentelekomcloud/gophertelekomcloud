package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/aom/v2/alarm"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestActionRuleLifecycle(t *testing.T) {
	client, err := clients.NewAOMV2Client()
	th.AssertNoErr(t, err)

	clientSMN, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	topicUrn := createSmnTopic(t, clientSMN)

	cloud, err := clients.EnvOS.Cloud()
	th.AssertNoErr(t, err)

	ruleName := tools.RandomString("tf_action_rule_", 4)

	createOpts := alarm.CreateActionRuleOpts{
		RuleName:             ruleName,
		ProjectId:            client.ProjectID,
		UserName:             cloud.AuthInfo.Username,
		Desc:                 "created by acceptance test",
		Type:                 "1",
		NotificationTemplate: "aom.built-in.template.en",
		SmnTopics: []alarm.SmnTopic{
			{
				Name:       topicNameFromUrn(topicUrn),
				TopicUrn:   topicUrn,
				PushPolicy: 0,
			},
		},
	}

	t.Logf("Attempting to create alarm action rule: %s", ruleName)
	err = alarm.CreateActionRule(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete alarm action rule: %s", ruleName)
		err := alarm.DeleteActionRule(client, []string{ruleName})
		th.AssertNoErr(t, err)
		t.Logf("Deleted alarm action rule: %s", ruleName)
	})

	t.Logf("Attempting to list alarm action rules")
	rules, err := alarm.ListActionRules(client)
	th.AssertNoErr(t, err)

	created, found := findActionRule(rules, ruleName)
	if !found {
		t.Fatal("created alarm action rule not found in list")
	}
	th.AssertEquals(t, createOpts.NotificationTemplate, created.NotificationTemplate)
	th.AssertEquals(t, createOpts.Desc, created.Desc)
	tools.PrintResource(t, created)

	t.Logf("Attempting to update alarm action rule: %s", ruleName)
	updateOpts := alarm.UpdateActionRuleOpts{
		RuleName:             ruleName,
		ProjectId:            client.ProjectID,
		UserName:             cloud.AuthInfo.Username,
		Desc:                 "updated by acceptance test",
		Type:                 createOpts.Type,
		NotificationTemplate: createOpts.NotificationTemplate,
		SmnTopics:            createOpts.SmnTopics,
	}
	err = alarm.UpdateActionRule(client, updateOpts)
	th.AssertNoErr(t, err)

	rules, err = alarm.ListActionRules(client)
	th.AssertNoErr(t, err)

	updated, found := findActionRule(rules, ruleName)
	if !found {
		t.Fatal("updated alarm action rule not found in list")
	}
	th.AssertEquals(t, updateOpts.Desc, updated.Desc)
}

func findActionRule(rules []alarm.ActionRule, name string) (alarm.ActionRule, bool) {
	for _, rule := range rules {
		if rule.RuleName == name {
			return rule, true
		}
	}
	return alarm.ActionRule{}, false
}
