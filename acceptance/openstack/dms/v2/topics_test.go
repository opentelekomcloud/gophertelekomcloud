package v2

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsTopicsLifecycle(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := getKafkaInstanceID(t, client)

	// Single-broker instances accept only replication 1, which requires sync_replication=false.
	topicName := tools.RandomString("dms-topic-", 8)
	_, err = topics.Create(client, instanceID, topics.CreateOpts{
		Name:             topicName,
		Partition:        3,
		Replication:      1,
		RetentionTime:    100,
		SyncReplication:  false,
		SyncMessageFlush: true,
		Description:      "acceptance test topic",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		res := deleteTopic(t, client, instanceID, topicName)
		th.AssertEquals(t, topicName, res.Topics[0].Name)

		err := golangsdk.WaitFor(60, func() (bool, error) {
			found, err := findTopic(client, instanceID, topicName)
			return found == nil, err
		})
		th.AssertNoErr(t, err)
	})

	var topic *topics.Topic
	err = golangsdk.WaitFor(60, func() (bool, error) {
		found, err := findTopic(client, instanceID, topicName)
		if err != nil {
			return false, err
		}
		topic = found
		return topic != nil, nil
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, topic)

	th.AssertEquals(t, topicName, topic.Name)
	th.AssertEquals(t, 3, topic.Partition)
	th.AssertEquals(t, 1, topic.Replication)
	th.AssertEquals(t, 100, topic.RetentionTime)
	th.AssertEquals(t, false, topic.SyncReplication)
	th.AssertEquals(t, true, topic.SyncMessageFlush)
	th.AssertEquals(t, "acceptance test topic", topic.Description)
	th.AssertEquals(t, 0, topic.TopicType)
	th.AssertEquals(t, true, topic.CreatedAt > 0)

	detail, err := topics.Get(client, instanceID, topicName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, detail)

	th.AssertEquals(t, topicName, detail.Name)
	th.AssertEquals(t, 3, len(detail.Partitions))
	for _, p := range detail.Partitions {
		th.AssertEquals(t, 1, len(p.Replicas))
		th.AssertEquals(t, p.Leader, p.Replicas[0].Broker)
		th.AssertEquals(t, true, p.Replicas[0].Leader)
		th.AssertEquals(t, true, p.Replicas[0].InSync)
	}

	partitions, err := topics.ListPartitions(client, instanceID, topicName, topics.ListPartitionsOpts{
		Offset: pointerto.Int(0),
		Limit:  pointerto.Int(10),
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, partitions)

	th.AssertEquals(t, 3, partitions.Total)
	th.AssertEquals(t, 3, len(partitions.Partitions))
	for _, p := range partitions.Partitions {
		th.AssertEquals(t, int64(0), p.MessageCount)
	}

	// No client produces to the new topic, so the list is expected to be empty.
	producers, err := topics.ListProducers(client, instanceID, topicName, topics.ListProducersOpts{
		Offset: pointerto.Int(0),
		Limit:  pointerto.Int(10),
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, producers)

	th.AssertEquals(t, 0, producers.Total)
	th.AssertEquals(t, 0, len(producers.Producers))

	// Explicit false for sync_message_flush checks that pointer fields are sent.
	err = topics.Update(client, instanceID, topics.UpdateOpts{
		Topics: []topics.UpdateItem{{
			Name:             topicName,
			RetentionTime:    pointerto.Int(200),
			SyncMessageFlush: pointerto.Bool(false),
			Description:      "updated acceptance test topic",
		}},
	})
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(60, func() (bool, error) {
		found, err := findTopic(client, instanceID, topicName)
		if err != nil || found == nil {
			return false, err
		}
		topic = found
		return topic.RetentionTime == 200 && !topic.SyncMessageFlush, nil
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, topic)

	th.AssertEquals(t, 200, topic.RetentionTime)
	th.AssertEquals(t, false, topic.SyncMessageFlush)
	th.AssertEquals(t, "updated acceptance test topic", topic.Description)
	th.AssertEquals(t, 3, topic.Partition)
	th.AssertEquals(t, 1, topic.Replication)
	th.AssertEquals(t, false, topic.SyncReplication)
}

// getKafkaInstanceID resolves the existing Kafka instance from OS_DMS_KAFKA_INSTANCE_NAME.
func getKafkaInstanceID(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Helper()

	name := clients.EnvOS.GetEnv("DMS_KAFKA_INSTANCE_NAME")
	if name == "" {
		t.Skip("OS_DMS_KAFKA_INSTANCE_NAME env var is not set")
	}

	res, err := lifecycle.List(client, lifecycle.ListOpts{
		Engine:         dmsEngine,
		Name:           name,
		ExactMatchName: "true",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(res.Instances))

	return res.Instances[0].InstanceID
}

// findTopic pages through the topic list and returns the topic with the given name, or nil.
func findTopic(client *golangsdk.ServiceClient, instanceID, name string) (*topics.Topic, error) {
	const limit = 50

	for offset := 0; ; offset += limit {
		res, err := topics.List(client, instanceID, topics.ListOpts{
			Offset: strconv.Itoa(offset),
			Limit:  strconv.Itoa(limit),
		})
		if err != nil {
			return nil, err
		}
		for i := range res.Topics {
			if res.Topics[i].Name == name {
				return &res.Topics[i], nil
			}
		}
		if len(res.Topics) < limit || offset+limit >= res.Total {
			return nil, nil
		}
	}
}
