package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	cbrtags "github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTags(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	firstTag := tags.ResourceTag{
		Key:   "TestKey",
		Value: "TestValue",
	}

	secondTag := tags.ResourceTag{
		Key:   "TestKey2",
		Value: "TestValue2",
	}

	combineTag := []tags.ResourceTag{
		firstTag,
		secondTag,
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
		Tags:        combineTag,
	}
	vault, err := vaults.Create(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID))
	})

	projectTags, err := cbrtags.ShowVaultProjectTag(client).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(projectTags), 2)

	instances, err := cbrtags.ShowVaultResourceInstances(client, cbrtags.ResourceInstancesRequest{
		Tags: []tags.ListedTag{{
			Key:    firstTag.Key,
			Values: []string{firstTag.Value},
		}},
		Action: cbrtags.Filter,
	})
	th.AssertNoErr(t, err)

	resourceID := instances.Resources[0].ResourceID
	th.AssertEquals(t, resourceID, vault.ID)

	vaultTags, err := cbrtags.ShowVaultTag(client, resourceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vaultTags[0].Key == firstTag.Key || vaultTags[0].Key == secondTag.Key, true)

	th.AssertNoErr(t, cbrtags.DeleteVaultTag(client, resourceID, combineTag).Err)
	vaultTags, _ = cbrtags.ShowVaultTag(client, resourceID).Extract()
	th.AssertEquals(t, len(vaultTags), 0)

	th.AssertNoErr(t, cbrtags.CreateVaultTags(client, resourceID, []tags.ResourceTag{firstTag}).Err)
	vaultTags, _ = cbrtags.ShowVaultTag(client, resourceID).Extract()
	th.AssertEquals(t, len(vaultTags), 1)
}
