package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v1/publicips"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createPublicIP(t *testing.T) (*publicips.PublicIP, *golangsdk.ServiceClient) {
	t.Helper()

	client, err := clients.NewVPCV1Client()
	th.AssertNoErr(t, err)

	createOpts := publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type:  "5_bgp",
			Alias: tools.RandomString("publicip-acc-", 3),
		},
		Bandwidth: publicips.BandWidth{
			Name:      tools.RandomString("publicip-acc-band-", 3),
			Size:      10,
			ShareType: "PER",
		},
	}

	publicIP, err := publicips.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created EIP: %s", publicIP.ID)

	t.Cleanup(func() {
		th.AssertNoErr(t, publicips.Delete(client, publicIP.ID))
		t.Logf("Deleted EIP: %s", publicIP.ID)
	})

	return publicIP, client
}

func TestPublicIPList(t *testing.T) {
	publicIP, client := createPublicIP(t)

	list, err := publicips.List(client, publicips.ListOpts{Limit: 100})
	th.AssertNoErr(t, err)

	found := false
	for _, item := range list {
		if item.ID == publicIP.ID {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)
}

func TestPublicIPLifecycle(t *testing.T) {
	publicIP, client := createPublicIP(t)

	th.AssertEquals(t, "5_bgp", publicIP.Type)
	th.AssertNotEquals(t, "", publicIP.PublicIpAddress)

	tools.PrintResource(t, publicIP)

	alias := tools.RandomString("publicip-acc-update-", 3)
	updated, err := publicips.Update(client, publicIP.ID, publicips.UpdateOpts{Alias: alias})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, alias, updated.Alias)

	got, err := publicips.Get(client, publicIP.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, publicIP.ID, got.ID)
	th.AssertEquals(t, alias, got.Alias)
	th.AssertEquals(t, 10, got.BandwidthSize)
	th.AssertEquals(t, "PER", got.BandwidthShareType)
}
