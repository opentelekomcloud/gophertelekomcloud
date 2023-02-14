package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/backup"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateWholeImageFromECS(t *testing.T) {
	clientCSBS, err := clients.NewCsbsV1Client()
	th.AssertNoErr(t, err)

	client1, client2 := getClient(t)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	t.Cleanup(func() { openstack.DeleteCloudServer(t, computeClient, ecs.ID) })

	checkpoint, err := backup.Create(clientCSBS, ecs.ID, backup.CreateOpts{
		BackupName:   tools.RandomString("backup-", 3),
		Description:  "bla-bla",
		ResourceType: "OS::Nova::Server",
	})
	t.Cleanup(func() {
		err = backup.Delete(clientCSBS, checkpoint.Id)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	fromECS, err := images.CreateWholeImageFromECS(client1, images.CreateWholeImageFromECSOpts{
		Name:       tools.RandomString("ims-test-", 3),
		InstanceId: ecs.ID,
	})
	th.AssertNoErr(t, err)

	image := jobEntities(t, client1, client2, fromECS)

	fromCSBS, err := images.CreateWholeImageFromCBRorCSBS(client1, images.CreateWholeImageFromCBRorCSBSOpts{
		Name:     tools.RandomString("ims-test-", 3),
		BackupId: checkpoint.Id,
	})

	jobEntities(t, client1, client2, fromCSBS)

	err = tags.AddImageTag(client2, tags.AddImageTagOpts{
		ImageId: image.ImageId,
		Tag: tag.ResourceTag{
			Key:   "test",
			Value: "testValue",
		},
	})
	th.AssertNoErr(t, err)

	imagesTags, err := tags.ListImagesTags(client2)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, imagesTags)

	err = tags.DeleteImageTag(client2, tags.DeleteImageTagOpts{
		ImageId: image.ImageId,
		Key:     "test",
	})
	th.AssertNoErr(t, err)

	err = tags.BatchAddOrDeleteTags(client2, tags.BatchAddOrDeleteTagsOpts{
		ImageId: image.ImageId,
		Action:  "create",
		Tags: []tag.ResourceTag{
			{
				Key:   "test1",
				Value: "testValue1",
			},
			{
				Key:   "test2",
				Value: "testValue2",
			},
		},
	})
	th.AssertNoErr(t, err)

	imageTags, err := tags.ListImageTags(client2, image.ImageId)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, imageTags)

	byTags, err := tags.ListImageByTags(client2, tags.ListImageByTagsOpts{
		Action: "count",
		Tags: []tag.ListedTag{{
			Key:    "test1",
			Values: []string{"testValue1"},
		}},
		Limit: "1",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, byTags)
}