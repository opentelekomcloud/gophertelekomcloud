package ims

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateImageFromOBS(t *testing.T) {
	client1, client2 := getClient(t)

	obsClient, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("ims-sdk-test", 5))

	_, err = obsClient.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	img, err := downloadIMG(t)

	objectName := tools.RandomString("ims-test-", 5)

	_, err = obsClient.PutFile(&obs.PutFileInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    objectName,
			},
		},
		SourceFile: img.Name(),
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteObject(&obs.DeleteObjectInput{
			Bucket: bucketName,
			Key:    objectName,
		})
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	imgName := tools.RandomString("ims-test-", 5)

	fromOBS, err := images.CreateImageFromOBS(client2, images.CreateImageFromOBSOpts{
		Name:     imgName,
		OsType:   "Linux",
		ImageUrl: bucketName + ":" + objectName,
		MinDisk:  1,
		Tags:     []string{"rancher"},
	})
	th.AssertNoErr(t, err)

	err = others.WaitForJob(client1, *fromOBS, 1000)
	th.AssertNoErr(t, err)

	job, err := others.ShowJob(client1, *fromOBS)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = images.DeleteImage(client2, images.DeleteImageOpts{ImageId: job.Entities.ImageId})
		th.AssertNoErr(t, err)
	})
}
