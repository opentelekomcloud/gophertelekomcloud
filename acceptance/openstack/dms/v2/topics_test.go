package v2

import (
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDmsTopicsList(t *testing.T) {
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
	t.Cleanup(func() { deleteTopic(t, client, instanceID, topicName) })

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
