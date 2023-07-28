package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/alarmnotifications"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmNotificationWorkflow(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	smnClient, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	topic := createTopic(t, smnClient)
	defer deleteTopic(t, smnClient, topic)

	t.Logf("Attempting to get Alarm Notification")
	alarmNotification, err := alarmnotifications.List(client).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Got Alarm Notification: %s", alarmNotification.ID)

	t.Logf("Attempting to update Alarm Notification: %s", alarmNotification.ID)
	enabled := true
	updateOpts := alarmnotifications.UpdateOpts{
		Enabled:       &enabled,
		TopicURN:      &topic,
		SendFrequency: 5,
		Times:         200,
		Threat:        []string{"xss", "sqli"},
	}

	_, err = alarmnotifications.Update(client, alarmNotification.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated Alarm Notification: %s", alarmNotification.ID)
	defer func() {
		t.Logf("Attempting to disable Alarm Notification: %s", alarmNotification.ID)
		disabled := false
		emptyTopic := ""
		updateOpts := alarmnotifications.UpdateOpts{
			Enabled:       &disabled,
			TopicURN:      &emptyTopic,
			SendFrequency: 5,
			Times:         1,
			Threat:        []string{"all"},
		}
		alarm, err := alarmnotifications.Update(client, alarmNotification.ID, updateOpts).Extract()
		th.AssertNoErr(t, err)
		t.Logf("Disabled Alarm Notification: %s", alarm.ID)
	}()

	newAlarmNotification, err := alarmnotifications.List(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, enabled, newAlarmNotification.Enabled)
	th.AssertEquals(t, updateOpts.SendFrequency, newAlarmNotification.SendFrequency)
	th.AssertEquals(t, updateOpts.Times, newAlarmNotification.Times)
}

func createTopic(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create SMN topic")
	opts := topics.CreateOps{
		Name: tools.RandomString("topic-", 3),
	}
	topic, err := topics.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)
	return topic.TopicUrn
}

func deleteTopic(t *testing.T, client *golangsdk.ServiceClient, topicURN string) {
	t.Logf("Attempting to delete SMN topic: %s", topicURN)
	err := topics.Delete(client, topicURN).Err
	th.AssertNoErr(t, err)
	t.Logf("Deleted SMN topic: %s", topicURN)
}
