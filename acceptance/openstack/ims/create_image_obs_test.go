package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	images1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateImageFromOBS(t *testing.T) {
	t.Skip("long run test only for manual purpose")
	client1, client2 := getClient(t)

	bucketName, objectName := makeOBS(t)

	fromOBS, err := images.CreateImageFromOBS(client2, images.CreateImageFromOBSOpts{
		Name:     tools.RandomString("ims-test-", 5),
		OsType:   "Linux",
		ImageUrl: bucketName + ":" + objectName,
		MinDisk:  1,
		Tags:     []string{"rancher"},
	})
	th.AssertNoErr(t, err)

	image := jobEntities(t, client1, client2, fromOBS)

	copied, err := others.CopyImageInRegion(client1, others.CopyImageInRegionOpts{
		ImageId: image.ImageId,
		Name:    tools.RandomString("ims-test-", 5),
	})
	th.AssertNoErr(t, err)
	jobEntities(t, client1, client2, copied)

	quota, err := others.ShowImageQuota(client1)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, quota)
}

func TestCreateDataImage(t *testing.T) {
	client1, client2 := getClient(t)

	bucketName, objectName := makeOBS(t)

	fromOBS, err := images1.CreateDataImage(client1, images1.CreateDataImageOpts{
		Name:     tools.RandomString("ims-test-", 5),
		ImageUrl: bucketName + ":" + objectName,
		MinDisk:  1,
	})
	th.AssertNoErr(t, err)

	image := jobEntities(t, client1, client2, fromOBS)

	updated, err := images.UpdateImage(client2, image.ImageId, []images.UpdateImageOpts{{
		Op:    "replace",
		Path:  "/name",
		Value: tools.RandomString("DataImage-test-", 5),
	}})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, updated)
}
