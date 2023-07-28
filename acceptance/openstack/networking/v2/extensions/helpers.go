package extensions

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createEip(t *testing.T) *eips.PublicIp {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create eip/bandwidth")
	eipCreateOpts := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
		},
		Bandwidth: eips.BandwidthOpts{
			ShareType: "PER",
			Name:      tools.RandomString("acc-band-", 3),
			Size:      100,
		},
	}

	eip, err := eips.Apply(client, eipCreateOpts).Extract()
	th.AssertNoErr(t, err)

	// wait to be DOWN
	t.Logf("Waiting for eip %s to be active", eip.ID)
	err = waitForEipToActive(client, eip.ID, 600)
	th.AssertNoErr(t, err)

	newEip, err := eips.Get(client, eip.ID).Extract()
	th.AssertNoErr(t, err)

	t.Logf("Created eip/bandwidth: %s", newEip.ID)

	return newEip
}

func deleteEip(t *testing.T, eipID string) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to delete eip/bandwidth: %s", eipID)

	err = eips.Delete(client, eipID).Err
	th.AssertNoErr(t, err)

	// wait to be deleted
	t.Logf("Waitting for eip %s to be deleted", eipID)

	err = waitForEipToDelete(client, eipID, 600)
	th.AssertNoErr(t, err)

	t.Logf("Deleted eip/bandwidth: %s", eipID)
}

func waitForEipToActive(client *golangsdk.ServiceClient, eipID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		eip, err := eips.Get(client, eipID).Extract()
		if err != nil {
			return false, err
		}
		if eip.Status == "DOWN" {
			return true, nil
		}

		return false, nil
	})
}

func waitForEipToDelete(client *golangsdk.ServiceClient, eipID string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := eips.Get(client, eipID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
