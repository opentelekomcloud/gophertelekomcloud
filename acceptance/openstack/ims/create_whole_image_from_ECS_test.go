package ims

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/cbr/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	tag "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateWholeImageFromCBR(t *testing.T) {
	t.Skip("long run test only for manual purpose")
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	client1, client2 := getClient(t)

	vault, _, _, _ := v3.CreateCBR(t, client)
	list, err := backups.List(client, backups.ListOpts{VaultID: vault.ID})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = backups.Delete(client, list[0].ID)
		th.AssertNoErr(t, err)
	})

	fromCBR, err := images.CreateWholeImageFromCBRorCSBS(client1, images.CreateWholeImageFromCBRorCSBSOpts{
		Name:           tools.RandomString("ims-test-", 3),
		BackupId:       list[0].ID,
		WholeImageType: "CBR",
	})
	th.AssertNoErr(t, err)

	jobEntities(t, client1, client2, fromCBR)
}

func TestCreateWholeImageFromECS(t *testing.T) {
	client1, client2 := getClient(t)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	t.Cleanup(func() { openstack.DeleteCloudServer(t, computeClient, ecs.ID) })

	fromECS, err := images.CreateWholeImageFromECS(client1, images.CreateWholeImageFromECSOpts{
		Name:       tools.RandomString("ims-test-", 3),
		InstanceId: ecs.ID,
	})
	th.AssertNoErr(t, err)

	image := jobEntities(t, client1, client2, fromECS)

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
