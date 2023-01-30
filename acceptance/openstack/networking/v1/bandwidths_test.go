package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBandwidthsList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	// Create eip/bandwidth
	eip := CreateEip(t, client, 1)
	defer DeleteEip(t, client, eip.ID)

	// Query a bandwidth
	existingBandwidths, err := bandwidths.List(client, bandwidths.ListOpts{
		ShareType: "PER",
	}).Extract()
	th.AssertNoErr(t, err)

	for _, b := range existingBandwidths {
		if b.ID == eip.BandwidthID {
			t.Logf("Bandwith with ID %s is in bandwidth list", b.ID)
			return
		}
	}
	t.Errorf("Failed to find created bandwidth")
}

func TestBandwidthsLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	bandwidthSize := 100
	// Create eip/bandwidth
	eip := CreateEip(t, client, bandwidthSize)
	defer DeleteEip(t, client, eip.ID)

	tools.PrintResource(t, eip)

	// Update a bandwidth
	newBandWidthSize := bandwidthSize / 2
	updateOpts := bandwidths.UpdateOpts{
		Size: newBandWidthSize,
	}
	updatedBand, err := bandwidths.Update(client, eip.BandwidthID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	// Query a bandwidth
	newBandWidth, err := bandwidths.Get(client, updatedBand.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newBandWidth)
}
