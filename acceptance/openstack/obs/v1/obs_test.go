package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	bucketName = "obs-test-gophertelekomcloud"
)

func TestObsBucketLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
	}

	_, err = client.CreateBucket(createOpts)
	th.AssertNoErr(t, err)

	_, err = client.DeleteBucket(bucketName)
	th.AssertNoErr(t, err)
}

func TestObsObjectLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	createOpts := &obs.CreateBucketInput{
		Bucket: bucketName,
	}

	_, err = client.CreateBucket(createOpts)
	th.AssertNoErr(t, err)

	defer func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	}()

	objectName := tools.RandomString("test-obs-", 5)

	objectOpts := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    objectName,
			},
		},
	}
	_, err = client.PutObject(objectOpts)
	th.AssertNoErr(t, err)

	_, err = client.DeleteObject(&obs.DeleteObjectInput{
		Bucket: bucketName,
		Key:    objectName,
	})
}
