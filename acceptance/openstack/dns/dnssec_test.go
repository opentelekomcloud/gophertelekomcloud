package dns

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestZonesDNSSEC(t *testing.T) {
	client, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)

	zoneName := tools.RandomString("dnssec-zone", 3)
	zone, err := zones.Create(client, zones.CreateOpts{
		Name: zoneName + ".com.",
		TTL:  300,
	}).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created DNS public zone: %s", zone.ID)

	defer func() {
		_, err := zones.Delete(client, zone.ID).Extract()
		th.AssertNoErr(t, err)
		t.Logf("Deleted DNS public zone: %s", zone.ID)
	}()

	initial, err := zones.GetDNSSEC(client, zone.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "DISABLE", initial.Status)

	enabled, err := zones.EnableDNSSEC(client, zone.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, enabled)
	th.AssertEquals(t, "ENABLE", enabled.Status)
	th.AssertEquals(t, 257, enabled.Flag)
	if enabled.DSRecord == "" {
		t.Fatal("expected DS record after enabling DNSSEC")
	}

	current, err := zones.GetDNSSEC(client, zone.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, enabled.KeyTag, current.KeyTag)

	disabled, err := zones.DisableDNSSEC(client, zone.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "DISABLE", disabled.Status)
}
