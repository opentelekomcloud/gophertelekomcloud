package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v2/extensions/admin/roles"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListRoleResponse(t)

	count := 0

	err := roles.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := roles.ExtractRoles(page)
		if err != nil {
			t.Errorf("Failed to extract users: %v", err)
			return false, err
		}

		expected := []roles.Role{
			{
				ID:          "123",
				Name:        "compute:admin",
				Description: "Nova Administrator",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, count)
}

func TestAddUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockAddUserRoleResponse(t)

	err := roles.AddUser(client.ServiceClient(), "{tenant_id}", "{user_id}", "{role_id}").ExtractErr()

	th.AssertNoErr(t, err)
}

func TestDeleteUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteUserRoleResponse(t)

	err := roles.DeleteUser(client.ServiceClient(), "{tenant_id}", "{user_id}", "{role_id}").ExtractErr()

	th.AssertNoErr(t, err)
}
