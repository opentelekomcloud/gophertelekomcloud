package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ucs/v1/fleets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUCSFleetLifecycle(t *testing.T) {
	client, err := clients.NewUCSV1Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("ucs-fleet-", 5)
	uid, err := fleets.Create(client, fleets.CreateOpts{
		Metadata: fleets.CreateMetadata{Name: name},
		Spec: &fleets.CreateSpec{
			Description: "created by acceptance test",
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, fleets.Delete(client, uid))
	})

	got, err := fleets.Get(client, uid)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, got.Metadata.Name)

	err = fleets.Update(client, uid, fleets.UpdateOpts{Description: "updated by acceptance test"})
	th.AssertNoErr(t, err)

	list, err := fleets.List(client, fleets.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, list.Total >= 1)
}
