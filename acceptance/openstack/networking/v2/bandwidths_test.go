package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBandwidthList(t *testing.T) {
	t.Skipf("disabled: working only in eu-nl")
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	bandwidthPages, err := bandwidths.List(client).AllPages()
	th.AssertNoErr(t, err)

	_, err = bandwidths.ExtractBandwidths(bandwidthPages)
	th.AssertNoErr(t, err)
}

func TestBandwidthLifecycle(t *testing.T) {
	t.Skipf("disabled: working only in eu-nl")
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
	defer func() {
		t.Logf("Attempting to delete BandwidthV2: %s", bandwidth.ID)
		err := bandwidths.Delete(client, bandwidth.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted BandwidthV2: %s", bandwidth.ID)
	}()
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
}
