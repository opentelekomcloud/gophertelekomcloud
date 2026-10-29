package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// createSmnTopic creates an SMN topic required by the AOM alarm action rule
// and returns its URN. Deletion is registered via t.Cleanup.
func createSmnTopic(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	topic, err := topics.Create(client, topics.CreateOpts{
		Name: topicName,
	})
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	t.Cleanup(func() {
		t.Logf("Attempting to delete SMN topic: %s", topic.TopicUrn)
		err := topics.Delete(client, topic.TopicUrn)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN topic: %s", topic.TopicUrn)
	})

	return topic.TopicUrn
}

// topicNameFromUrn extracts the topic name from a topic URN
// (urn:smn:<region>:<project_id>:<topic_name>).
func topicNameFromUrn(topicUrn string) string {
	idx := strings.LastIndex(topicUrn, ":")
	if idx < 0 {
		return topicUrn
	}
	return topicUrn[idx+1:]
}
