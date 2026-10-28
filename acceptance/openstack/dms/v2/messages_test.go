package v2

import (
	"strconv"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/messages"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMessageQuery(t *testing.T) {
	name := kafkaInstanceName(t)

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := existingKafkaInstance(t, client, name).InstanceID

	topicName := tools.RandomString("dms-msg-", 8)
	_, err = topics.Create(client, instanceID, topics.CreateOpts{
		Name:             topicName,
		Partition:        1,
		Replication:      1,
		RetentionTime:    1,
		SyncReplication:  false,
		SyncMessageFlush: false,
		Description:      "message query acceptance test topic",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		deleteTopic(t, client, instanceID, topicName)
	})

	err = golangsdk.WaitFor(60, func() (bool, error) {
		found, err := findTopic(client, instanceID, topicName)
		return found != nil, err
	})
	th.AssertNoErr(t, err)

	beginning, err := messages.ShowPartitionBeginningMessage(client, instanceID, topicName, 0)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, beginning)
	th.AssertEquals(t, topicName, beginning.Topic)
	th.AssertEquals(t, 0, beginning.Partition)

	end, err := messages.ShowPartitionEndMessage(client, instanceID, topicName, 0)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, end)
	th.AssertEquals(t, topicName, end.Topic)
	th.AssertEquals(t, 0, end.Partition)
	th.AssertEquals(t, beginning.Offset, end.Offset)

	byPeriod, err := messages.ShowMessages(client, instanceID, topicName, messages.ShowMessagesOpts{
		Limit: 10,
		// The API rejects offset 0 (DMS.00400830), although the docs allow values >= 0.
		Offset:    pointerto.Int(1),
		Partition: "-1",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, byPeriod)
	th.AssertEquals(t, 0, len(byPeriod.Messages))

	now := time.Now()
	byInstance, err := messages.ShowInstanceMessages(client, instanceID, messages.ShowInstanceMessagesOpts{
		Topic:     topicName,
		Asc:       pointerto.Bool(false),
		StartTime: strconv.FormatInt(now.Add(-time.Hour).UnixMilli(), 10),
		EndTime:   strconv.FormatInt(now.UnixMilli(), 10),
		Limit:     "10",
		Offset:    "0",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, byInstance)
	th.AssertEquals(t, int64(0), byInstance.Total)
	th.AssertEquals(t, 0, len(byInstance.Messages))

	byOffset, err := messages.ShowPartitionMessage(client, instanceID, topicName, 0, messages.ShowPartitionMessageOpts{
		MessageOffset: strconv.Itoa(beginning.Offset),
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, byOffset)
	th.AssertEquals(t, 0, len(byOffset.Message))
}
