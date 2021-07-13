package v1

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createOBSBucket(t *testing.T) string {
	t.Logf("Attempting to create OBS bucket")
	client, err := clients.NewOBSClientWithoutHeader()
	th.AssertNoErr(t, err)
	bucketName := strings.ToLower(tools.RandomString("obs-cts-test", 5))

	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
		ACL:    "public-read",
	}

	_, err = client.CreateBucket(createOpts)
	th.AssertNoErr(t, err)

	t.Logf("Created OBS Bucket: %s", bucketName)

	return bucketName
}

func deleteOBSBucket(t *testing.T, bucketName string) {
	t.Logf("Attempting to delete OBS bucket: %s", bucketName)
	client, err := clients.NewOBSClientWithoutHeader()
	th.AssertNoErr(t, err)

	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
	t.Logf("Deleted OBS Bucket: %s", bucketName)
}

func createSMNTopic(t *testing.T) *topics.Topic {
	t.Logf("Attempting to create SMNv2 topic")
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	smnTopicName := strings.ToLower(tools.RandomString("smn-cts-test", 5))

	createOpts := topics.CreateOps{
		Name: smnTopicName,
	}

	smnTopic, err := topics.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created SMNv2 Topic: %s", smnTopic.TopicUrn)
	return smnTopic
}

func deleteSMNTopic(t *testing.T, topicUrn string) {
	t.Logf("Attempting to delete SMNv2 topic: %s", topicUrn)
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, topics.Delete(client, topicUrn).ExtractErr())
	t.Logf("Deleted SMNv2 Topic: %s", topicUrn)
}
