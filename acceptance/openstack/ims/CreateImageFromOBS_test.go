package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateImageFromOBS(t *testing.T) {
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

	err = others.WaitForJob(client1, *fromOBS, 1000)
	th.AssertNoErr(t, err)

	job, err := others.ShowJob(client1, *fromOBS)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = images.DeleteImage(client2, images.DeleteImageOpts{ImageId: job.Entities.ImageId})
		th.AssertNoErr(t, err)
	})
}
