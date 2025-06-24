package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTopicsWorkflow(t *testing.T) {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create SMN topic")
	topicName := tools.RandomString("topic-", 3)
	opts := topics.CreateOpts{
		Name: topicName,
	}

	topic, err := topics.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN topic: %s", topic.TopicUrn)

	defer deleteTopic(t, client, topic.TopicUrn)

	displayName := tools.RandomString("topic-name-", 3)

	t.Logf("Attempting to update SMN topic display name")

	_, err = topics.Update(client, topics.UpdateOpts{
		Id:          topic.TopicUrn,
		DisplayName: displayName,
	})

	th.AssertNoErr(t, err)

	t.Logf("Updated display name")

	t.Logf("Attempting to retrieve SMN topic")

	getResp, err := topics.Get(client, topic.TopicUrn)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getResp.DisplayName, displayName)

	t.Logf("Retrieved SMN topic")

	t.Logf("Attempting to list all SMN topics")

	listResp, err := topics.List(client, topics.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(listResp) > 0)

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
	th.AssertNoErr(t, err)

	defer func() {
		tagOptsUpdate := []tags.ResourceTag{
			{
				Key:   "kuh",
				Value: "lala",
			},
		}

		t.Logf("Attempting to delete SMN topic tag")
		err = tags.Delete(tagsClient, "smn_topic", topicName, tagOptsUpdate).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	listTags, err := tags.Get(tagsClient, "smn_topic", topicName).Extract()
	th.AssertNoErr(t, err)
	if len(listTags) < 1 {
		t.Fatal("empty SMN topic tags list")
	}
	t.Logf("SMN topic tags: %s", listTags)

}
