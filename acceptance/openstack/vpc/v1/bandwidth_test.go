package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBandwidthLifecycle(t *testing.T) {
	publicIP, client := createPublicIP(t)

	got, err := bandwidths.Get(client, publicIP.BandwidthId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, publicIP.BandwidthId, got.ID)
	th.AssertEquals(t, 10, got.Size)
	th.AssertEquals(t, "PER", got.ShareType)

	list, err := bandwidths.List(client, bandwidths.ListOpts{Limit: 100})
	th.AssertNoErr(t, err)
	found := false
	for _, item := range list {
		if item.ID == publicIP.BandwidthId {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	updated, err := bandwidths.Update(client, publicIP.BandwidthId, bandwidths.UpdateOpts{
		Size: 15,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 15, updated.Size)
}
