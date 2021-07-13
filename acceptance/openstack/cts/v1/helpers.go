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
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)
	bucketName := strings.ToLower(tools.RandomString("obs-cts-test", 5))

	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
		ACL:    "public-read",
	}

	_, err = client.CreateBucket(createOpts)
	th.AssertNoErr(t, err)

	return bucketName
}

func deleteOBSBucket(t *testing.T, bucketName string) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
}

func createSMNTopic(t *testing.T) *topics.Topic {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	smnTopicName := strings.ToLower(tools.RandomString("smn-cts-test", 5))

	createOpts := topics.CreateOps{
		Name: smnTopicName,
	}

	smnTopic, err := topics.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	return smnTopic
}

func deleteSMNTopic(t *testing.T, topicUrn string) {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, topics.Delete(client, topicUrn).ExtractErr())
}
