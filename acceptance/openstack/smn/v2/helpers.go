package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createTopic(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOpts{
		Name: topicName,
	}
	topic, err := topics.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	return topic.TopicUrn
}

func deleteTopic(t *testing.T, client *golangsdk.ServiceClient, topicURN string) {
	t.Logf("Attempting to delete SMN topic: %s", topicURN)
	err := topics.Delete(client, topicURN)
	th.AssertNoErr(t, err)
	t.Logf("Deleted SMN topic: %s", topicURN)
}
