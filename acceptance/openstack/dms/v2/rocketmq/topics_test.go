package rocketmq

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const topicNotFound = "DMS.00050004"

func TestRocketMQTopics(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	t.Run("LifeCycle", func(t *testing.T) {
		testRocketMQTopicLifeCycle(t, client, instanceID)
	})
	t.Run("BatchDelete", func(t *testing.T) {
		testRocketMQTopicBatchDelete(t, client, instanceID)
	})
}

func testRocketMQTopicLifeCycle(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	topicName := createRocketMQTopic(t, client, instanceID)

	topic, err := topics.Get(client, instanceID, topicName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, topicName, topic.Name)
	th.AssertEquals(t, "NORMAL", topic.MessageType)
	th.AssertEquals(t, "all", topic.Permission)
	th.AssertEquals(t, "", topic.Description)

	listed, err := topics.List(client, instanceID, topics.ListOpts{Limit: 50})
	th.AssertNoErr(t, err)
	var found bool
	for _, tp := range listed.Topics {
		if tp.Name == topicName {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	t.Logf("Attempting to update RocketMQ topic: %s", topicName)
	err = topics.Update(client, instanceID, topicName, topics.UpdateOpts{
		Description: pointerto.String("updated description"),
	})
	th.AssertNoErr(t, err)

	topic, err = topics.Get(client, instanceID, topicName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated description", topic.Description)

	status, err := topics.GetStatus(client, instanceID, topicName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, int64(0), status.MinOffset)
	th.AssertEquals(t, int64(0), status.MaxOffset)

	// The consumer group list is only populated by consumers subscribed to
	// the topic, which the test has no client for.
	consumers, err := topics.ListGroups(client, instanceID, topicName, topics.ListGroupsOpts{Limit: 10})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(consumers.Groups))

	t.Logf("Attempting to delete RocketMQ topic: %s", topicName)
	th.AssertNoErr(t, topics.Delete(client, instanceID, topicName))
	_, err = topics.Get(client, instanceID, topicName)
	th.AssertEquals(t, true, hasErrorCode(err, topicNotFound))
}

func testRocketMQTopicBatchDelete(t *testing.T, client *golangsdk.ServiceClient, instanceID string) {
	names := []string{
		createRocketMQTopic(t, client, instanceID),
		createRocketMQTopic(t, client, instanceID),
	}

	t.Logf("Attempting to batch delete RocketMQ topics: %v", names)
	job, err := topics.BatchDelete(client, instanceID, topics.BatchDeleteOpts{Topics: names})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, job.JobID != "")

	for _, name := range names {
		th.AssertNoErr(t, golangsdk.WaitFor(300, func() (bool, error) {
			_, err := topics.Get(client, instanceID, name)
			if hasErrorCode(err, topicNotFound) {
				return true, nil
			}
			return false, err
		}))
	}
}

func createRocketMQTopic(t *testing.T, client *golangsdk.ServiceClient, instanceID string) string {
	t.Logf("Attempting to create RocketMQ topic")

	topic, err := topics.Create(client, instanceID, topics.CreateOpts{
		Name:        tools.RandomString("topic-acc-", 8),
		QueueNum:    3,
		MessageType: "NORMAL",
	})
	th.AssertNoErr(t, err)
	t.Logf("RocketMQ topic created: %s", topic.ID)

	t.Cleanup(func() {
		err := topics.Delete(client, instanceID, topic.ID)
		if err != nil && !hasErrorCode(err, topicNotFound) {
			t.Errorf("failed to delete RocketMQ topic %s: %s", topic.ID, err)
		}
	})

	return topic.ID
}
