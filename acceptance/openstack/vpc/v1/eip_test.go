package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEIPListing(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	eipName := tools.RandomString("eip-test-", 5)
	eipCreateOpts := eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type: "5_bgp",
			Name: eipName,
		},
		Bandwidth: eips.BandwidthOpts{
			ShareType: "PER",
			Name:      tools.RandomString("acc-band-", 3),
			Size:      100,
		},
	}
	eip, err := eips.Apply(client, eipCreateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = eips.Delete(client, eip.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	cases := map[string]eips.ListOpts{
		"ID": {
			ID: eip.ID,
		},
		"PublicIP": {
			PublicAddress: eip.PublicAddress,
		},
	}
	for name, opts := range cases {
		t.Run(name, func(t *testing.T) {
			opts := opts
			t.Parallel()

			list, err := eips.List(client, opts)
			th.AssertNoErr(t, err)
			th.AssertEquals(t, 1, len(list))
			th.AssertEquals(t, eip.ID, list[0].ID)
		})
	}
}
