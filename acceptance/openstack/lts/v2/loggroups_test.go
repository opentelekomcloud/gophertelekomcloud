package v2

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsGroupsLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}

	created, err := groups.CreateLogGroup(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = groups.Delete(client, created)
		th.AssertNoErr(t, err)
	})

	got, err := groups.ListLogGroups(client, created)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, got.LogGroupName)

	log.Printf("Creating LTS Group, ID: %s", got.LogGroupId)
	th.AssertEquals(t, created, got.LogGroupId)
}
