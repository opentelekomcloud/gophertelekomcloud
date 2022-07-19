package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTags(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	projectTags, err := tags.ShowVaultProjectTag(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(projectTags) > 0, true)

	testTag := projectTags[0]

	instances, err := tags.ShowVaultResourceInstances(client, tags.ResourceInstancesRequest{
		Tags:   []tags.Tag{testTag},
		Action: tags.Filter,
	}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, instances.TotalCount > 0, true)

	resourceID := instances.Resources[0].ResourceID

	vaultTags, err := tags.ShowVaultTag(client, resourceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(vaultTags.Tags) > 0, true)

	monoTag := tags.MonoTag{
		Key:   "TestKey",
		Value: "TestValue",
	}

	th.AssertNoErr(t, tags.CreateVaultTags(client, resourceID, monoTag).ExtractErr())
	vaultTags, _ = tags.ShowVaultTag(client, resourceID).Extract()
	isExist := false
	for i := range vaultTags.Tags {
		if vaultTags.Tags[i].Key == monoTag.Key {
			th.AssertEquals(t, vaultTags.Tags[i].Value, monoTag.Value)
			isExist = true
		}
	}
	th.AssertEquals(t, isExist, true)

	th.AssertNoErr(t, tags.DeleteVaultTag(client, resourceID, testTag.Key).ExtractErr())
	vaultTags, _ = tags.ShowVaultTag(client, resourceID).Extract()
	for i := range vaultTags.Tags {
		if vaultTags.Tags[i].Key == monoTag.Key {
			panic("Tag should be deleted")
		}
	}

	th.AssertNoErr(t, tags.BatchCreateAndDeleteVaultTags(client, resourceID, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []tags.MonoTag{monoTag},
		Action: tags.Delete,
	}).ExtractErr())

	th.AssertNoErr(t, tags.BatchCreateAndDeleteVaultTags(client, resourceID, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []tags.MonoTag{monoTag},
		Action: tags.Create,
	}).ExtractErr())
}
