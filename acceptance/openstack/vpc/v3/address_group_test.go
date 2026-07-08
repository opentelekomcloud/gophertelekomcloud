package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/addressgroup"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCAddressGroupV3Lifecycle(t *testing.T) {
	client, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("address-grp-acc-", 3)

	addressGroup, err := addressgroup.Create(client, addressgroup.CreateOpts{
		AddressGroup: addressgroup.AddressGroupOptions{
			Name:        name,
			Description: "create-address-group",
			IpVersion:   4,
			IpSet:       []string{"192.168.10.10"},
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, addressgroup.Delete(client, addressGroup.ID))
	})

	updatedDescription := "updated-description"
	updatedName := name + "-updated"

	addressGroup, err = addressgroup.Update(client, addressGroup.ID, addressgroup.UpdateOpts{
		AddressGroup: addressgroup.UpdateAddressGroupOptions{
			Name:        updatedName,
			Description: updatedDescription,
			IpSet:       []string{"192.168.10.11"},
		},
	})
	th.AssertNoErr(t, err)

	updatedGroup, err := addressgroup.Get(client, addressGroup.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedName, updatedGroup.Name)
	th.AssertEquals(t, updatedDescription, updatedGroup.Description)

	addressGroupList, err := addressgroup.List(client, addressgroup.ListQueryParams{})
	th.AssertNoErr(t, err)
	th.AssertNotEquals(t, 0, len(addressGroupList.AddressGroups))
}
