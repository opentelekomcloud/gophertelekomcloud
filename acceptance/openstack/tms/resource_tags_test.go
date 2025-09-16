package tms

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/blockstorage/v2/volumes"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/tms/v1/quota"
	rt "github.com/opentelekomcloud/gophertelekomcloud/openstack/tms/v1/resource-tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTMSRecourceTagsLifecycle(t *testing.T) {
	// if os.Getenv("RUN_TMS_TAGS") == "" {
	// 	t.Skip("not for CI")
	// }
	client, err := clients.NewTmsV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewTmsV2Client()
	th.AssertNoErr(t, err)

	clientEvs, err := clients.NewBlockStorageV2Client()
	th.AssertNoErr(t, err)

	createVolOpts := volumes.CreateOpts{
		Size: 40,
		Name: tools.RandomString("tf-evs-disk-", 4),
	}
	t.Log("Attempting to create EVS volume")
	vol, err := volumes.Create(clientEvs, createVolOpts).Extract()
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(clientEvs, vol.ID, "available", 120)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete EVS Volume: %s", vol.ID)
		err = volumes.Delete(clientEvs, vol.ID, volumes.DeleteOpts{}).ExtractErr()
		th.AssertNoErr(t, err)
	})

	resourceTags := []rt.ResourceTag{
		{
			Key:   "test-1",
			Value: pointerto.String("test-1"),
		},
	}

	createOpts := rt.BatchOpts{
		Tags: resourceTags,
		Resources: []rt.Resource{
			{
				ResourceType: "disk",
				ResourceId:   vol.ID,
			},
		},
		ProjectId: clientEvs.ProjectID,
	}

	t.Logf("Attempting to create TMS Resource Tag: %s", resourceTags[0].Key)
	_, err = rt.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete TMS Resource Tag: %s", resourceTags[0].Key)
		_, err := rt.Delete(client, createOpts)
		th.AssertNoErr(t, err)
		t.Logf("Deleted TMS Resource Tag: %s", resourceTags[0].Key)
	})

	t.Logf("Attempting to list TMS Resource Tag for volume: %s", vol.ID)
	listTags, err := rt.List(clientV2, vol.ID, rt.ListOpts{
		ResourceType: "disk",
		ProjectId:    clientEvs.ProjectID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listTags[0].Key, "test-1")

	t.Log("Attempting to list Resources with TMS Resource Tags")
	res, err := rt.ListResources(client, rt.ListResourceOpts{
		ResourceTypes: []string{"disk", "ecs"},
		Tags: []rt.ListResourceTag{
			{
				Key:    "test-1",
				Values: []string{"test-1"},
			},
		},
		ProjectId: clientEvs.ProjectID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, res[0].ResourceId, vol.ID)

	t.Log("Attempting to list all TMS Resource Tag keys")
	keys, err := rt.ListKeys(client, rt.ListKeysOpts{
		RegionId: clientEvs.RegionID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(keys) > 0)

	values, err := rt.ListValues(client, rt.ListValueOpts{
		RegionId: clientEvs.RegionID,
		Key:      listTags[0].Key,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(values) > 0)
}

func TestTMSQuotaList(t *testing.T) {
	if os.Getenv("RUN_TMS_TAGS") == "" {
		t.Skip("unstable test")
	}
	client, err := clients.NewTmsV1Client()
	th.AssertNoErr(t, err)

	q, err := quota.List(client)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(q))
}
