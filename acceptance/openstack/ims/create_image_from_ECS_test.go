package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	image1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateImageFromECS(t *testing.T) {
	client1, client2 := getClient(t)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	t.Cleanup(func() { openstack.DeleteCloudServer(t, computeClient, ecs.ID) })

	fromDisk, err := images.CreateImageFromDisk(client2, images.CreateImageFromDiskOpts{
		Name:      tools.RandomString("ims-test-", 3),
		VolumeId:  ecs.VolumeAttached[0].ID,
		OsVersion: "Debian GNU/Linux 10.0.0 64bit",
	})
	th.AssertNoErr(t, err)
	jobEntities(t, client1, client2, fromDisk)

	fromECS, err := images.CreateImageFromECS(client2, images.CreateImageFromECSOpts{
		Name:       tools.RandomString("ims-test-", 3),
		InstanceId: ecs.ID,
	})
	th.AssertNoErr(t, err)

	image := jobEntities(t, client1, client2, fromECS)

	obsClient, bucketName := newBucket(t)

	export, err := image1.ExportImage(client1, image1.ExportImageOpts{
		ImageId:    image.ImageId,
		BucketUrl:  bucketName + ":" + image.ImageName,
		FileFormat: "zvhd",
	})
	t.Cleanup(func() {
		_, err = obsClient.DeleteObject(&obs.DeleteObjectInput{
			Bucket: bucketName,
			Key:    image.ImageName,
		})
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	err = others.WaitForJob(client1, *export, 1000)
	th.AssertNoErr(t, err)

	// quick, err := images.ImportImageQuick(client2, images.ImportImageQuickOpts{
	// 	Name:      tools.RandomString("ims-test-", 3),
	// 	OsVersion: "Debian GNU/Linux 10.0.0 64bit",
	// 	ImageUrl:  bucketName + ":" + image.ImageName,
	// 	MinDisk:   100,
	// })
	// th.AssertNoErr(t, err)
	// jobEntities(t, client1, client2, quick)
}
