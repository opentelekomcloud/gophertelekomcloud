package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/bandwidths"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBandwidthLifecycle(t *testing.T) {
	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	tools.AcquireQuota(t, "eip", 1)

	publicIP, err := publicips.Create(client, publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type:  "5_bgp",
			Alias: tools.RandomString("bandwidth-acc-", 3),
		},
		Bandwidth: publicips.BandWidth{
			Name:      tools.RandomString("bandwidth-acc-band-", 3),
			Size:      10,
			ShareType: "PER",
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, publicips.Delete(client, publicIP.ID))
	})

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
