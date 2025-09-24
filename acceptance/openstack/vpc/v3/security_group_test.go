package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpc/v3/security/group"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVPCSecurityGroupV3Lifecycle(t *testing.T) {
	client, err := clients.NewVPCV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("sec-grp-acc-", 3)
	createOpts := group.CreateOpts{
		SecurityGroup: group.SecurityGroupOptions{
			Name:        name,
			Description: "created",
		},
	}

	secGroup, err := group.Create(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, secGroup.Description, "created")

	t.Cleanup(func() {
		th.AssertNoErr(t, group.Delete(client, secGroup.ID))
	})

	groupList, err := group.List(client, group.ListQueryParams{
		Id: []string{secGroup.ID},
	})
	th.AssertNoErr(t, err)
	th.AssertNotEquals(t, 0, len(groupList.SecurityGroups))
	th.AssertEquals(t, "created", groupList.SecurityGroups[0].Description)

	_, err = group.Update(client, secGroup.ID, group.UpdateOpts{
		SecurityGroup: group.SecurityGroupUpdateOptions{
			Description: "updated",
		},
	})
	th.AssertNoErr(t, err)

	updatedGroup, err := group.Get(client, secGroup.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated", updatedGroup.Description)
}
