package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/alarm"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsAlarmList(t *testing.T) {
	clientSMN, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOps{
		Name: topicName,
	}
	topic, err := topics.Create(clientSMN, opts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	t.Cleanup(func() {
		t.Logf("Attempting to delete SMN topic: %s", topic.TopicUrn)
		err := topics.Delete(clientSMN, topic.TopicUrn).ExtractErr()
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
