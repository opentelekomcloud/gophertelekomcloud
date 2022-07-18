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

	tagsResp := tags.ShowVaultProjectTag(client)
	tagsRes, err := tagsResp.Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(tagsRes) > 0, true)

	tag := tagsRes[0]
	sysTag := tags.SysTag{
		Key:   tag.Key,
		Value: tag.Values[0],
	}

	req := tags.ResourceInstancesRequest{
		Tags:   []tags.Tag{tag},
		Action: tags.Filter,
	}

	insResp := tags.ShowVaultResourceInstances(client, req)
	insRes, err := insResp.Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, insRes.TotalCount > 0, true)

	resId := insRes.Resources[0].ResourceID

	vaultResp := tags.ShowVaultTag(client, resId)
	vaultRes, err := vaultResp.Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(vaultRes.Tags) > 0, true)

	delResp := tags.DeleteVaultTag(client, resId, tag.Key)
	th.AssertNoErr(t, delResp.ExtractErr())

	createResp := tags.CreateVaultTags(client, resId, sysTag)
	th.AssertNoErr(t, createResp.ExtractErr())

	batchDel := tags.BatchCreateAndDeleteVaultTags(client, resId, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []tags.SysTag{sysTag},
		Action: tags.Delete,
	})
	th.AssertNoErr(t, batchDel.ExtractErr())

	batchCreate := tags.BatchCreateAndDeleteVaultTags(client, resId, tags.BulkCreateAndDeleteVaultTagsRequest{
		Tags:   []tags.SysTag{sysTag},
		Action: tags.Create,
	})
	th.AssertNoErr(t, batchCreate.ExtractErr())
}
