package v1_1

import (
	"errors"
	"fmt"
	"io"
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
	uploadFile(t, clientOBS, fileName, strings.NewReader("testData"))

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

}

func uploadFile(t *testing.T, client *obs.ObsClient, fileName string, data io.Reader) {
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
		Body: data,
	})
	th.AssertNoErr(t, err)
}

func cleanupBucket(t *testing.T, client *obs.ObsClient) {
	objs, err := client.ListObjects(&obs.ListObjectsInput{
		Bucket: bucketName,
	})
	th.AssertNoErr(t, err)

	toDelete := make([]obs.ObjectToDelete, 0, len(objs.Contents))
	for _, obj := range objs.Contents {
		toDelete = append(toDelete, obs.ObjectToDelete{
			Key: obj.Key,
		})
	}

	_, err = client.DeleteObjects(&obs.DeleteObjectsInput{
		Bucket:  bucketName,
		Objects: toDelete,
	})
	th.AssertNoErr(t, err)

	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
}
