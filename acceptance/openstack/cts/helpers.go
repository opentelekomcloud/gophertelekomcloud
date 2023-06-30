package cts

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateOBSBucket(t *testing.T) string {
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

func DeleteOBSBucket(t *testing.T, bucketName string) {
	t.Logf("Attempting to delete OBS bucket: %s", bucketName)
	client, err := clients.NewOBSClientWithoutHeader()
	th.AssertNoErr(t, err)

	input := &obs.ListObjectsInput{}
	input.Bucket = bucketName

	objectsList, err := client.ListObjects(input)
	th.AssertNoErr(t, err)

	if len(objectsList.Contents) > 0 {
		objects := make([]obs.ObjectToDelete, len(objectsList.Contents))
		for i, content := range objectsList.Contents {
			objects[i].Key = content.Key
		}
		deleteOpts := &obs.DeleteObjectsInput{
			Bucket:  bucketName,
			Objects: objects,
		}
		_, err = client.DeleteObjects(deleteOpts)
		th.AssertNoErr(t, err)
		t.Logf("Deleted OBS Bucket objects: %s", objects)
	}

	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
	t.Logf("Deleted OBS Bucket: %s", bucketName)
}
