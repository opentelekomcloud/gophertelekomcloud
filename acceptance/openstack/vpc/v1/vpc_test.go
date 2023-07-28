package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCListing(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	createOpts := vpcs.CreateOpts{
		Name: tools.RandomString("vpc-acc-", 3),
	}
	vpc, err := vpcs.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = vpcs.Delete(client, vpc.ID).Err
		th.AssertNoErr(t, err)
	})

	cases := map[string]vpcs.ListOpts{
		"ID": {
			ID: vpc.ID,
		},
		"Name": {
			Name: vpc.Name,
		},
	}
	for name, opts := range cases {
		t.Run(name, func(t *testing.T) {
			opts := opts
			t.Parallel()

			list, err := vpcs.List(client, opts)
			th.AssertNoErr(t, err)
			th.AssertEquals(t, 1, len(list))
			th.AssertEquals(t, vpc.ID, list[0].ID)
		})
	}
}

func TestVPCLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("vpc-acc-", 3)
	createOpts := vpcs.CreateOpts{
		Name:        name,
		Description: "some interesting description",
	}

	vpc, err := vpcs.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpc.EnableSharedSnat, false)
	th.AssertEquals(t, vpc.Description, "some interesting description")
	th.AssertEquals(t, vpc.Name, name)

	t.Cleanup(func() {
		err = vpcs.Delete(client, vpc.ID).Err
		th.AssertNoErr(t, err)
	})

	snatEnable := true
	name = tools.RandomString("vpc-acc-update-", 3)
	description := "some interesting description (update)"
	updateOpts := vpcs.UpdateOpts{
		Name:             name,
		Description:      &description,
		EnableSharedSnat: &snatEnable,
	}

	vpc, err = vpcs.Update(client, vpc.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, vpc.EnableSharedSnat, snatEnable)
	th.AssertEquals(t, vpc.Description, description)
	th.AssertEquals(t, vpc.Name, name)
}
