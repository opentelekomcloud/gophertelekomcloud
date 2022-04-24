package dns

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/nameservers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestZonesList(t *testing.T) {
	client, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)

	listOpts := zones.ListOpts{}
	pages, err := zones.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allZones, err := zones.ExtractZones(pages)
	th.AssertNoErr(t, err)

	for _, zone := range allZones {
		tools.PrintResource(t, zone)
	}
}

func TestZonesCRUD(t *testing.T) {
	client, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create DNS public zone")
	zoneName := tools.RandomString("public-zone", 3)
	createOpts := zones.CreateOpts{
		Description: "interesting public zone",
		Name:        zoneName + ".com",
		TTL:         300,
	}
	zone, err := zones.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created DNS public zone: %s", zone.ID)

	n, err := nameservers.List(client, zone.ID).Extract()
	th.AssertNoErr(t, err)

	for _, nameserver := range n {
		tools.PrintResource(t, nameserver)
	}

	defer func() {
		t.Logf("Attempting to delete DNS public zone: %s", zone.ID)
		_, err := zones.Delete(client, zone.ID).Extract()
		th.AssertNoErr(t, err)
		t.Logf("Deleted DNS public zone: %s", zone.ID)
	}()
	th.AssertEquals(t, createOpts.TTL, zone.TTL)

	t.Logf("Attempting to update DNS public zone: %s", zone.ID)
	updateOpts := zones.UpdateOpts{
		Email: "bla-bla@mail.com",
		TTL:   400,
	}
	update, err := zones.Update(client, zone.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated DNS public zone")
	th.AssertEquals(t, updateOpts.Email, update.Email)
	th.AssertEquals(t, updateOpts.TTL, update.TTL)
}
