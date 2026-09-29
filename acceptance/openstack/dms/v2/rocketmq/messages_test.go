package rocketmq

import (
	"strconv"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/messages"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRocketMQMessages(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)
	topicName := createRocketMQTopic(t, client, instanceID)
	start := time.Now().Add(-time.Minute)

	t.Logf("Attempting to send a message to RocketMQ topic: %s", topicName)
	sent, err := messages.Send(client, instanceID, messages.SendOpts{
		Topic: topicName,
		Body:  "acceptance test message",
		PropertyList: []messages.Property{
			{Name: "KEYS", Value: "acc-key"},
			{Name: "TAGS", Value: "acc-tag"},
		},
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, sent)
	th.AssertEquals(t, topicName, sent.Topic)
	th.AssertEquals(t, "acceptance test message", sent.Body)
	th.AssertEquals(t, true, sent.MsgID != "")

	byID, err := messages.List(client, instanceID, messages.ListOpts{
		Topic: topicName,
		MsgID: sent.MsgID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, byID)
	th.AssertEquals(t, 1, len(byID.Messages))
	th.AssertEquals(t, sent.MsgID, byID.Messages[0].MsgID)
	th.AssertEquals(t, "acceptance test message", byID.Messages[0].Body)

	byTime, err := messages.List(client, instanceID, messages.ListOpts{
		Topic:     topicName,
		Limit:     10,
		StartTime: strconv.FormatInt(start.UnixMilli(), 10),
		EndTime:   strconv.FormatInt(time.Now().Add(time.Minute).UnixMilli(), 10),
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, byTime.Total)
	th.AssertEquals(t, sent.MsgID, byTime.Messages[0].MsgID)

	// The trace is written asynchronously and is empty right after sending.
	var trace *messages.ListTraceResponse
	th.AssertNoErr(t, golangsdk.WaitFor(120, func() (bool, error) {
		trace, err = messages.ListTrace(client, instanceID, messages.ListTraceOpts{
			MsgID: sent.MsgID,
			Limit: 10,
		})
		if err != nil {
			return false, err
		}
		return trace.Total > 0, nil
	}))
	tools.PrintResource(t, trace)
	th.AssertEquals(t, "Pub", trace.Trace[0].TraceType)
	th.AssertEquals(t, sent.MsgID, trace.Trace[0].MsgID)
	th.AssertEquals(t, topicName, trace.Trace[0].Topic)

	groupName := createRocketMQGroup(t, client, instanceID)

	// Dead letter messages can only be resent and consumption can only be
	// verified while a consumer is connected to the group, which the test
	// has no client for.
	t.Logf("Attempting to resend dead letter messages of RocketMQ consumer group: %s", groupName)
	_, err = messages.ResendDeadLetter(client, instanceID, messages.ResendDeadLetterOpts{
		Topic:     "%DLQ%" + groupName,
		MsgIDList: []string{sent.MsgID},
	})
	th.AssertEquals(t, true, hasErrorCode(err, groupNotOnline))

	t.Logf("Attempting to verify consumption of RocketMQ consumer group: %s", groupName)
	_, err = messages.VerifyConsumption(client, instanceID, messages.VerifyConsumptionOpts{
		Group:     groupName,
		Topic:     topicName,
		ClientID:  "127.0.0.1@acc-test",
		MsgIDList: []string{sent.MsgID},
	})
	th.AssertEquals(t, true, hasErrorCode(err, groupNotOnline))
}
