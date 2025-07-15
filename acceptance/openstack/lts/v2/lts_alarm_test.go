package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/alarm"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsAlarmList(t *testing.T) {
	clientSMN, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOpts{
		Name: topicName,
	}
	topic, err := topics.Create(clientSMN, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	t.Cleanup(func() {
		t.Logf("Attempting to delete SMN topic: %s", topic.TopicUrn)
		err := topics.Delete(clientSMN, topic.TopicUrn)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN topic: %s", topic.TopicUrn)
	})

	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list SMN topics by LTS api")
	listResp, err := alarm.ListTopic(
		client,
		alarm.ListOpts{
			Limit:  pointerto.Int(100),
			Offset: pointerto.Int(0),
		})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestLtsAlarmHistoryList(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	listResp, err := alarm.ListHistory(
		client,
		client.DomainID,
		alarm.ListHistoryQueryOpts{
			Type: "active_alert",
		},
		alarm.ListHistoryBodyOpts{
			WhetherCustomField: pointerto.Bool(true),
			StartTime:          1644108607,
			EndTime:            1744108607,
		})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, listResp)
}

func TestLtsAlarmsLifecycle(t *testing.T) {
	clientV2, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	t.Logf("Attempting to Create Log Group")
	group, err := groups.Create(clientV2, groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(clientV2, group)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(clientV2, streams.CreateOpts{
		GroupId:       group,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(clientV2, streams.DeleteOpts{
			GroupId:  group,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	clientSMN, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOpts{
		Name: topicName,
	}
	topic, err := topics.Create(clientSMN, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete SMN topic: %s", topic.TopicUrn)
		err := topics.Delete(clientSMN, topic.TopicUrn)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN topic: %s", topic.TopicUrn)
	})

	t.Logf("Attempting to Create LTS Alarm Keyword Rule")
	aname := tools.RandomString("keyword", 3)
	optsAlarm := alarm.CreateOpts{
		Name:        aname,
		Description: "description",
		DomainId:    clientV2.DomainID,
		Details: []alarm.Details{
			{
				LogStreamId:         stream,
				LogGroupName:        "",
				LogGroupId:          group,
				LogStreamName:       "",
				Keyword:             "test",
				Condition:           ">",
				Number:              100,
				SearchTimeRange:     10,
				SearchTimeRangeUnit: "minute",
			},
		},
		Frequency: &alarm.Frequency{
			Type:          "FIXED_RATE",
			CronExpr:      "",
			HourOfDay:     0,
			DayOfWeek:     0,
			FixedRate:     10,
			FixedRateUnit: "minute",
		},
		Severity:              "Critical",
		Send:                  pointerto.Bool(true),
		NotificationFrequency: 5,
		AlarmActionRuleName:   "my_rule",
		NotificationSave: &alarm.NotificationSave{
			Language: "en-us",
			Timezone: "xx/xx",
			UserName: "test",
			Topics: []alarm.TopicsCreate{
				{
					Name:        topicName,
					TopicUrn:    topic.TopicUrn,
					PushPolicy:  0,
					DisplayName: "",
				},
			},
			// TemplateName: "my_template",
		},
	}
	rule, err := alarm.CreateKeywordRule(clientV2, optsAlarm)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete LTS Alarm Keyword Rule: %s", rule)
		err := alarm.DeleteKeywordRule(clientV2, rule)
		th.AssertNoErr(t, err)
		t.Logf("Deleted LTS Alarm Keyword Rule: %s", rule)
	})

	rules, err := alarm.ListKeywordRules(clientV2)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(rules))

	t.Logf("Attempting to Update LTS Alarm Keyword Rule: %s", rule)
	aliasUp := tools.RandomString("keyword-update", 3)
	optsAlarmUpdate := alarm.UpdateOpts{
		ID:          rule,
		Name:        aname,
		Alias:       aliasUp,
		Description: "description",
		DomainId:    clientV2.DomainID,
		Details: []alarm.Details{
			{
				LogStreamId:         stream,
				LogGroupName:        "",
				LogGroupId:          group,
				LogStreamName:       "",
				Keyword:             "test",
				Condition:           ">",
				Number:              100,
				SearchTimeRange:     10,
				SearchTimeRangeUnit: "minute",
			},
		},
		Frequency: &alarm.Frequency{
			Type:          "FIXED_RATE",
			CronExpr:      "",
			HourOfDay:     0,
			DayOfWeek:     0,
			FixedRate:     10,
			FixedRateUnit: "minute",
		},
		Severity:              "Critical",
		Send:                  pointerto.Bool(true),
		NotificationFrequency: 5,
		AlarmActionRuleName:   "my_rule",
		NotificationSave: &alarm.NotificationSave{
			Language: "en-us",
			Timezone: "xx/xx",
			UserName: "test",
			Topics: []alarm.TopicsCreate{
				{
					Name:        topicName,
					TopicUrn:    topic.TopicUrn,
					PushPolicy:  0,
					DisplayName: "",
				},
			},
			// TemplateName: "my_template",
		},
	}
	ruleUpdate, err := alarm.UpdateKeywordRule(clientV2, optsAlarmUpdate)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, aliasUp, ruleUpdate.Alias)
}
