package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createTopic(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOps{
		Name: topicName,
	}
	topic, err := topics.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	tagsClient, err := clients.NewSmnV2TagsClient()
	th.AssertNoErr(t, err)

	tagsClient.MoreHeaders = map[string]string{
		"X-SMN-RESOURCEID-TYPE": "name",
	}
	tagOpts := []tags.ResourceTag{
		{
			Key:   "muh",
			Value: "lala",
		},
		{
			Key:   "kuh",
			Value: "lala",
		},
	}
	t.Logf("Attempting to create SMN topic tags: %s", tagOpts)
	err = tags.Create(tagsClient, "smn_topic", topicName, tagOpts).ExtractErr()

	listTags, err := tags.Get(tagsClient, "smn_topic", topicName).Extract()
	th.AssertNoErr(t, err)
	if len(listTags) < 0 {
		t.Fatal("empty SMN topic tags list")
	}
	t.Logf("SMN topic tags: %s", listTags)

	tagOptsUpdate := []tags.ResourceTag{
		{
			Key:   "kuh",
			Value: "lala",
		},
	}

	t.Logf("Attempting to delete SMN topic tag")
	err = tags.Delete(tagsClient, "smn_topic", topicName, tagOptsUpdate).ExtractErr()
	th.AssertNoErr(t, err)

	return topic.TopicUrn
}

func deleteTopic(t *testing.T, client *golangsdk.ServiceClient, topicURN string) {
	t.Logf("Attempting to delete SMN topic: %s", topicURN)
	err := topics.Delete(client, topicURN).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted SMN topic: %s", topicURN)
}
