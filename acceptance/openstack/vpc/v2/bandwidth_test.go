package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	vpcv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v2/bandwidths"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSharedBandwidthLifecycle(t *testing.T) {
	bandwidthClient, err := clients.NewVPCV2Client()
	th.AssertNoErr(t, err)
	vpcClient, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	tools.AcquireQuota(t, "eip", 1)

	shared, err := bandwidths.Create(bandwidthClient, bandwidths.CreateOpts{
		Name: tools.RandomString("shared-bandwidth-acc-", 4),
		Size: 5,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "WHOLE", shared.ShareType)
	bandwidthID := shared.ID
	t.Cleanup(func() {
		if bandwidthID != "" {
			th.AssertNoErr(t, bandwidths.Delete(bandwidthClient, bandwidthID))
		}
	})

	publicIP, err := vpcv1.Create(vpcClient, vpcv1.CreateOpts{
		Publicip: vpcv1.PublicIPRequest{
			Type:  "5_bgp",
			Alias: tools.RandomString("shared-bandwidth-acc-", 3),
		},
		Bandwidth: vpcv1.BandWidth{
			Name:      tools.RandomString("shared-bandwidth-acc-band-", 3),
			Size:      10,
			ShareType: "PER",
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, vpcv1.Delete(vpcClient, publicIP.ID))
	})

	added, err := bandwidths.AddEip(bandwidthClient, bandwidthID, bandwidths.AddEipOpts{
		PublicipInfo: []bandwidths.InsertPublicIPInfo{
			{PublicipId: publicIP.ID},
		},
	})
	th.AssertNoErr(t, err)
	found := false
	for _, info := range added.PublicipInfo {
		if info.PublicipId == publicIP.ID {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	th.AssertNoErr(t, bandwidths.RemoveEip(bandwidthClient, bandwidthID, bandwidths.RemoveEipOpts{
		PublicipInfo: []bandwidths.RemovePublicIPInfo{
			{PublicipId: publicIP.ID},
		},
		ChargeMode: "traffic",
		Size:       5,
	}))

	th.AssertNoErr(t, bandwidths.Delete(bandwidthClient, bandwidthID))
	bandwidthID = ""
}
