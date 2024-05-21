package v1_1

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/resource"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const resourceName = "testResource"
const bucketName = "dataart-test-bucket"
const fileName = "testFile.txt"

func TestDataArtsResourcesLifecycle(t *testing.T) {
	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	workspace := ""

	clientOBS, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	prepareTestBucket(t, clientOBS)
	t.Cleanup(func() {
		cleanupBucket(t, clientOBS)
	})

	t.Log("create a resource")

	createOpts := resource.Resource{
		Name:     resourceName,
		Type:     "file",
		Location: fmt.Sprintf("obs://%s/%s", bucketName, fileName),
	}

	r, err := resource.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule resource cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete resource: %s", resourceName)
		err := resource.Delete(client, r.ResourceId, workspace)
		th.AssertNoErr(t, err)
		t.Logf("resource is deleted: %s", resourceName)
	})

	t.Log("get resource")

	storedResource, err := resource.Get(client, r.ResourceId, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedResource)

	createOpts.Name = "anyNewName"

	err = resource.Update(client, r.ResourceId, createOpts)
	th.AssertNoErr(t, err)

	t.Log("should wait 5 seconds")
	time.Sleep(5 * time.Second)

	t.Log("get resource")

	storedResource, err = resource.Get(client, r.ResourceId, workspace)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedResource)
	th.CheckEquals(t, "anyNewName", storedResource.Name)
}

func prepareTestBucket(t *testing.T, client *obs.ObsClient) {

	_, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{
		Bucket: bucketName,
	})

	if err != nil {
		if !errors.As(err, &obs.ObsError{}) {
			th.AssertNoErr(t, err)
		}

		_, err = client.CreateBucket(&obs.CreateBucketInput{
			Bucket: bucketName,
		})
		th.AssertNoErr(t, err)
	}

	uploadFile(t, client)
}

func uploadFile(t *testing.T, client *obs.ObsClient) {
	_, err := client.GetObjectMetadata(&obs.GetObjectMetadataInput{
		Bucket: bucketName,
		Key:    fileName,
	})

	if err != nil {
		if !errors.As(err, &obs.ObsError{}) {
			th.AssertNoErr(t, err)
		}
	}

	_, err = client.PutObject(&obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    fileName,
			},
		},
		Body: strings.NewReader("testData"),
	})
	th.AssertNoErr(t, err)
}

func cleanupBucket(t *testing.T, client *obs.ObsClient) {
	_, err := client.DeleteObject(&obs.DeleteObjectInput{
		Bucket: bucketName,
		Key:    fileName,
	})
	th.AssertNoErr(t, err)
	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
}
