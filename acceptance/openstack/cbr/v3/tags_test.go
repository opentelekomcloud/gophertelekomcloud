package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTags(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	monoTag := vaults.Tag{
		Key:   "TestKey",
		Value: "TestValue",
	}

	opts := vaults.CreateOpts{
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			ObjectType:      "disk",
			ProtectType:     "backup",
			Size:            100,
		},
		Description: "gophertelemocloud testing vault",
		Name:        tools.RandomString("cbr-test-", 5),
		Resources:   []vaults.ResourceCreate{},
		Tags: []vaults.Tag{
			monoTag,
		},
	}
	vault, err := vaults.Create(client, opts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID).ExtractErr())
	}()

	projectTags, err := tags.ShowVaultProjectTag(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(projectTags) > 0, true)

	instances, err := tags.ShowVaultResourceInstances(client, tags.ResourceInstancesRequest{
		Tags: []tags.Tag{{
			Key:    monoTag.Key,
			Values: []string{monoTag.Value},
		}},
		Action: tags.Filter,
	}).Extract()
	th.AssertNoErr(t, err)

	resourceID := instances.Resources[0].ResourceID
	th.AssertEquals(t, resourceID, vault.ID)

	vaultTags, err := tags.ShowVaultTag(client, resourceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vaultTags.Tags[0].Key, monoTag.Key)

	th.AssertNoErr(t, tags.DeleteVaultTag(client, resourceID, monoTag.Key).ExtractErr())
	vaultTags, _ = tags.ShowVaultTag(client, resourceID).Extract()
	for i := range vaultTags.Tags {
		if vaultTags.Tags[i].Key == monoTag.Key {
			panic("Tag should be deleted")
		}
	}

	_, err = tags.CreateVaultTags(client, resourceID, monoTag).Extract()
	th.AssertNoErr(t, err)

	vaultTags, _ = tags.ShowVaultTag(client, resourceID).Extract()
	isExist := false
	for i := range vaultTags.Tags {
		if vaultTags.Tags[i].Key == monoTag.Key {
			th.AssertEquals(t, vaultTags.Tags[i].Value, monoTag.Value)
			isExist = true
		}
	}
	th.AssertEquals(t, isExist, true)

	th.AssertNoErr(t, tags.BatchCreateAndDeleteVaultTags(client, resourceID, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []vaults.Tag{monoTag},
		Action: tags.Delete,
	}).ExtractErr())

	th.AssertNoErr(t, tags.BatchCreateAndDeleteVaultTags(client, resourceID, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []vaults.Tag{monoTag},
		Action: tags.Create,
	}).ExtractErr())
}
