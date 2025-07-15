package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/publish"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/templates"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMessagePublishWorkflow(t *testing.T) {
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

	defer func() {
		t.Logf("Attempting to delete SMN topic: %s", topic.TopicUrn)
		err := topics.Delete(client, topic.TopicUrn)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN topic: %s", topic.TopicUrn)
	}()

	t.Logf("Attempting to publish message in text format")

	msgOpts := publish.PublishOpts{
		TopicUrn:   topic.TopicUrn,
		Subject:    "test message v2",
		Message:    "This is a test message v2",
		TimeToLive: "3600",
	}

	_, err = publish.Publish(client, msgOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to publish using message structure")

	msgOpts2 := publish.PublishOpts{
		TopicUrn:         topic.TopicUrn,
		Subject:          "test message v2",
		MessageStructure: "{\"default\": \"xxx\", \"APNS\": \"{\\\"aps\\\":{\\\"alert\\\":{\\\"title\\\":\\\"xxx\\\",\\\"body\\\":\\\"xxx\\\"}}}\"}",
		TimeToLive:       "3600",
	}

	_, err = publish.Publish(client, msgOpts2)
	th.AssertNoErr(t, err)

	rName := tools.RandomString("tf-", 8)

	t.Logf("Attempting to create SMN template")
	createOpts := templates.CreateOpts{
		Name:     rName,
		Content:  "Test message for topic: {topic_id}.",
		Protocol: "default",
	}

	template, err := templates.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN template: %s", template.MessageTemplateID)
	defer func() {
		t.Logf("Attempting to delete SMN template: %s", template.MessageTemplateID)
		err := templates.Delete(client, template.MessageTemplateID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN template: %s", template.MessageTemplateID)
	}()

	t.Logf("Attempting to publish using template name")

	msgOpts3 := publish.PublishOpts{
		TopicUrn:            topic.TopicUrn,
		Subject:             "test message v2",
		MessageTemplateName: rName,
		Tags: map[string]string{
			"topic_id": "topic_id3332",
		},
	}

	_, err = publish.Publish(client, msgOpts3)
	th.AssertNoErr(t, err)
}
