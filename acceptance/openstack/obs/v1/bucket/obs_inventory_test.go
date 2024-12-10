package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestOBSInventories(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	var (
		configId   = "test-id"
		bucketName = strings.ToLower(tools.RandomString("obs-sdk-test-", 5))
	)

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	inventoryOpts := obs.SetBucketInventoryInput{
		Bucket:            bucketName,
		InventoryConfigId: configId,
		BucketInventoryConfiguration: obs.BucketInventoryConfiguration{
			Id:        configId,
			IsEnabled: true,
			Schedule: obs.InventorySchedule{
				Frequency: "Daily",
			},
			Destination: obs.InventoryDestination{
				Format: "CSV",
				Bucket: bucketName,
				Prefix: "test",
			},
			Filter: obs.InventoryFilter{
				Prefix: "test",
			},
			IncludedObjectVersions: "All",
			OptionalFields: []obs.InventoryOptionalFields{
				{
					Field: "Size",
				},
			},
		},
	}

	_, err = client.SetBucketInventory(&inventoryOpts)
	th.AssertNoErr(t, err)

	getResp, err := client.GetBucketInventory(obs.GetBucketInventoryInput{
		BucketName:        bucketName,
		InventoryConfigId: configId,
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getResp)

	_, err = client.DeleteBucketInventory(&obs.DeleteBucketInventoryInput{
		Bucket:            bucketName,
		InventoryConfigId: configId,
	})
	th.AssertNoErr(t, err)
}
