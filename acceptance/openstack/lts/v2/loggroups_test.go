package v2

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/loggroups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsGroupsLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-de-02"
	}

	name := tools.RandomString("test-group-", 3)
	createOpts := loggroups.CreateOpts{
		LogGroupName: name,
		TTL:          7,
	}

	created, err := loggroups.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		err = loggroups.Delete(client, created.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	got, err := loggroups.Get(client, created.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, got.Name)

	log.Printf("Creating LTS Group, ID: %s", got.ID)
	th.AssertEquals(t, created.ID, got.ID)
}
