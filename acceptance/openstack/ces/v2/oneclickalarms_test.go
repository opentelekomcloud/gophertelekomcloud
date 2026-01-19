package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/oneclickalarms"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestOneClickAlarmsList(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to list one-click alarms")
	alarmsList, err := oneclickalarms.List(client)
	th.AssertNoErr(t, err)
	t.Logf("Found %d one-click alarms", len(alarmsList))
}

func TestOneClickAlarmsCRUD(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to enable one-click monitoring for ECS")
	createOpts := oneclickalarms.CreateOpts{
		OneClickAlarmId: "ECSSystemOneClickAlarm",
		DimensionNames: oneclickalarms.DimensionNames{
			Metric: []string{"instance_id"},
			Event:  []string{},
		},
		NotificationEnabled: false,
	}

	oneClickAlarmId, err := oneclickalarms.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Enabled one-click monitoring: %s", oneClickAlarmId)

	t.Cleanup(func() {
		t.Log("Attempting to disable one-click monitoring")
		_, err := oneclickalarms.BatchDelete(client, oneclickalarms.BatchDeleteOpts{
			OneClickAlarmIds: []string{oneClickAlarmId},
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to list one-click alarms to verify creation")
	alarmsList, err := oneclickalarms.List(client)
	th.AssertNoErr(t, err)

	found := false
	for _, alarm := range alarmsList {
		if alarm.OneClickAlarmId == oneClickAlarmId && alarm.Enabled {
			found = true
			t.Logf("Found enabled one-click alarm: %s, Namespace: %s", alarm.OneClickAlarmId, alarm.Namespace)
			break
		}
	}
	th.AssertEquals(t, found, true)

	t.Log("Attempting to get alarm rules for one-click monitoring")
	alarmRules, err := oneclickalarms.ListAlarmRules(client, oneClickAlarmId)
	th.AssertNoErr(t, err)
	t.Logf("Found %d alarm rules", len(alarmRules))
	th.AssertEquals(t, len(alarmRules) > 0, true)

	alarmRule := alarmRules[0]
	t.Logf("First alarm rule: ID=%s, Name=%s, Enabled=%v",
		alarmRule.AlarmId, alarmRule.Name, alarmRule.Enabled)

	t.Log("Attempting to disable alarm rule")
	_, err = oneclickalarms.BatchEnableAlarmRules(client, oneClickAlarmId, oneclickalarms.BatchEnableAlarmRulesOpts{
		AlarmIds:     []string{alarmRule.AlarmId},
		AlarmEnabled: false,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to enable alarm rule")
	_, err = oneclickalarms.BatchEnableAlarmRules(client, oneClickAlarmId, oneclickalarms.BatchEnableAlarmRulesOpts{
		AlarmIds:     []string{alarmRule.AlarmId},
		AlarmEnabled: true,
	})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(alarmRule.Policies) > 0, true)
	policy := alarmRule.Policies[0]
	t.Logf("First policy: ID=%s, MetricName=%s, Enabled=%v",
		policy.AlarmPolicyId, policy.MetricName, policy.Enabled)

	t.Log("Attempting to disable alarm policy")
	_, err = oneclickalarms.BatchEnablePolicies(client, oneClickAlarmId, alarmRule.AlarmId, oneclickalarms.BatchEnablePoliciesOpts{
		AlarmPolicyIds: []string{policy.AlarmPolicyId},
		Enabled:        false,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to enable alarm policy")
	_, err = oneclickalarms.BatchEnablePolicies(client, oneClickAlarmId, alarmRule.AlarmId, oneclickalarms.BatchEnablePoliciesOpts{
		AlarmPolicyIds: []string{policy.AlarmPolicyId},
		Enabled:        true,
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to update notifications")
	err = oneclickalarms.UpdateNotifications(client, oneClickAlarmId, oneclickalarms.UpdateNotificationsOpts{
		NotificationEnabled:   false,
		NotificationBeginTime: "00:00",
		NotificationEndTime:   "23:59",
	})
	th.AssertNoErr(t, err)
}
