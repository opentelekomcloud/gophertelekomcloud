package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/bandwidths"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
)

func TestBandwidthsUpdate(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create NetworkV1 client: %v", err)
	}
	bandwidthSize := 100

	// Create eip/bandwidth
	eip, err := createEipResource(t, client, bandwidthSize)
	if err != nil {
		t.Fatalf("Unable to create eip/banwidth pair: %s", err)
	}

	// Delete an eip
	defer deleteEipResource(t, client, eip.ID)

	tools.PrintResource(t, eip)

	// Update a bandwidth
	newBandWidthSize := bandwidthSize / 2
	updateOpts := bandwidths.UpdateOpts{
		Size: newBandWidthSize,
	}
	updatedBand, err := bandwidths.Update(client, eip.BandwidthID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update bandwidth: %s", err)
	}

	// Query a bandwidth
	newBandWidth, err := bandwidths.Get(client, updatedBand.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve bandwidth: %s", err)
	}

	tools.PrintResource(t, newBandWidth)
}

func createEipResource(t *testing.T, nwClient *golangsdk.ServiceClient, bandwidthSize int) (*eips.PublicIp, error) {
	bandName := tools.RandomString("testacc-", 8)

	t.Logf("Attempting to create eip/bandwidth: %s", bandName)
	eipCreateOpts := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			ShareType: "PER",
			Name:      bandName,
			Size:      bandwidthSize,
		},
	}

	eip, err := eips.Apply(nwClient, eipCreateOpts).Extract()
	if err != nil {
		return nil, err
	}

	// wait to be DOWN
	t.Logf("Waiting for eip %s to be active", eip.ID)
	if err := waitForEipToActive(nwClient, eip.ID, 600); err != nil {
		t.Fatalf("Error creating eip: %s", err)
	}
	newEip, err := eips.Get(nwClient, eip.ID).Extract()
	if err != nil {
		t.Fatalf("Error reading eip: %s", err)
	}

	t.Logf("Created eip/bandwidth: %s", bandName)

	return &newEip, nil
}

func deleteEipResource(t *testing.T, nwClient *golangsdk.ServiceClient, eipId string) {
	t.Logf("Attempting to delete eip/bandwidth: %s", eipId)

	err := eips.Delete(nwClient, eipId).ExtractErr()
	if err != nil {
		t.Fatalf("Error delete eip: %s", eipId)
	}

	// wait to be deleted
	t.Logf("Waitting for eip %s to be deleted", eipId)
	if err := waitForEipToDelete(nwClient, eipId, 600); err != nil {
		t.Fatalf("Error wait for deleting eip: %s", err)
	}

	t.Logf("Deleted eip/bandwidth: %s", eipId)
}

func waitForEipToActive(client *golangsdk.ServiceClient, eipId string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		eip, err := eips.Get(client, eipId).Extract()
		if err != nil {
			return false, err
		}
		if eip.Status == "DOWN" {
			return true, nil
		}

		return false, nil
	})
}

func waitForEipToDelete(client *golangsdk.ServiceClient, eipId string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := eips.Get(client, eipId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
		}

		return false, nil
	})
}
