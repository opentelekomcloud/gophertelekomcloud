package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/bandwidths"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/floatingips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBandwidthLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create BandwidthV2")
	bandwidthName := tools.RandomString("band-create", 3)
	createOpts := bandwidths.CreateOpts{
		Name: bandwidthName,
		Size: 20,
	}

	bandwidth, err := bandwidths.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete BandwidthV2: %s", bandwidth.ID)
		err := bandwidths.Delete(client, bandwidth.ID).Err
		th.AssertNoErr(t, err)
		t.Logf("Deleted BandwidthV2: %s", bandwidth.ID)
	})
	th.AssertEquals(t, createOpts.Name, bandwidth.Name)
	th.AssertEquals(t, createOpts.Size, bandwidth.Size)
	t.Logf("Created BandwidthV2: %s", bandwidth.ID)

	t.Logf("Attempting to update BandwidthV2: %s", bandwidth.ID)
	bandwidthName = tools.RandomString("band-update", 3)
	updateOpts := bandwidths.UpdateOpts{
		Name: bandwidthName,
		Size: 50,
	}
	_, err = bandwidths.Update(client, bandwidth.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newBandwidth, err := bandwidths.Get(client, bandwidth.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newBandwidth.Name)
	th.AssertEquals(t, updateOpts.Size, newBandwidth.Size)
	t.Logf("Updated BandwidthV2: %s", newBandwidth.ID)

	bandwidthPages, err := bandwidths.List(client).AllPages()
	th.AssertNoErr(t, err)

	bandwidthList, err := bandwidths.ExtractBandwidths(bandwidthPages)
	th.AssertNoErr(t, err)
	flag := -1
	for i, v := range bandwidthList {
		if v.ID == newBandwidth.ID {
			flag = i
			break
		}
	}
	th.AssertDeepEquals(t, newBandwidth, &bandwidthList[flag])
}

func TestBandwidthAssociate(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	bandwidthName := tools.RandomString("band-create", 3)
	createOpts := bandwidths.CreateOpts{
		Name: bandwidthName,
		Size: 20,
	}

	bandwidth, err := bandwidths.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	id := bandwidth.ID
	t.Cleanup(func() {
		t.Logf("Attempting to delete BandwidthV2: %s", id)
		err := bandwidths.Delete(client, id).Err
		th.CheckNoErr(t, err)
		t.Logf("Deleted BandwidthV2: %s", id)
	})

	eip, err := floatingips.Create(client, floatingips.CreateOpts{
		FloatingNetworkID: "0a2228f2-7f8a-45f1-8e09-9039e1d09975", // yep, it's static
	}).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err := floatingips.Delete(client, eip.ID).Err
		th.CheckNoErr(t, err)
	})

	err = tools.WaitFor(func() (bool, error) {
		fip, err := floatingips.Get(client, eip.ID).Extract()
		if err != nil {
			return false, err
		}
		return fip.Status == "ACTIVE" || fip.Status == "DOWN", nil
	})
	th.AssertNoErr(t, err)

	err = tools.WaitFor(func() (bool, error) {
		bdw, err := bandwidths.Get(client, id).Extract()
		if err != nil {
			return false, err
		}
		return bdw.Status == "NORMAL", nil
	})
	th.AssertNoErr(t, err)

	assOpts := bandwidths.InsertOpts{PublicIpInfo: []bandwidths.PublicIpInfoInsertOpts{
		{PublicIpID: eip.ID},
	}}
	_, err = bandwidths.Insert(client, id, assOpts).Extract()
	th.AssertNoErr(t, err)

	bnw, err := bandwidths.Get(client, id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(bnw.PublicIpInfo))

	remOpts := bandwidths.RemoveOpts{
		ChargeMode: "bandwidth",
		Size:       1,
		PublicIpInfo: []bandwidths.PublicIpInfoID{
			{PublicIpID: eip.ID},
		},
	}
	err = bandwidths.Remove(client, id, remOpts).Err
	th.AssertNoErr(t, err)

	bnwAfter, err := bandwidths.Get(client, id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(bnwAfter.PublicIpInfo))
}
