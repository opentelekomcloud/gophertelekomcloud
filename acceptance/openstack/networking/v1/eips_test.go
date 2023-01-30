package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEipList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	elasticIP := CreateEip(t, client, 100)
	listOpts := eips.ListOpts{
		BandwidthID: elasticIP.BandwidthID,
	}
	t.Cleanup(func() {
		DeleteEip(t, client, elasticIP.ID)

		elasticIPs, err := eips.List(client, listOpts)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 0, len(elasticIPs))
	})

	elasticIPs, err := eips.List(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(elasticIPs))
	th.AssertEquals(t, elasticIP.ID, elasticIPs[0].ID)
	th.AssertEquals(t, elasticIP.PublicAddress, elasticIPs[0].PublicAddress)
}

func TestEipLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eip := CreateEip(t, client, 100)
	t.Cleanup(func() { DeleteEip(t, client, eip.ID) })

	tools.PrintResource(t, eip)
}

func TestEipTagsLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eip := CreateEip(t, client, 100)
	t.Cleanup(func() { DeleteEip(t, client, eip.ID) })

	networkV2Client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	eipTags := []tags.ResourceTag{
		{
			Key:   "muh",
			Value: "lala",
		},
		{
			Key:   "kuh",
			Value: "lala",
		},
		{
			Key:   "luh",
			Value: "lala",
		},
	}
	createEipTags(t, networkV2Client, eip.ID, eipTags)
	t.Cleanup(func() { deleteEipTags(t, networkV2Client, eip.ID, eipTags) })

	newTags, err := tags.Get(networkV2Client, "publicips", eip.ID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Assert of length between `createOptsTags` and `tags.Get()`")
	th.AssertEquals(t, len(eipTags), len(newTags))
}
