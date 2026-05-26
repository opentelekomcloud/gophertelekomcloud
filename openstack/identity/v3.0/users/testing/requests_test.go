package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreateUserWithExternalIdentity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateUserWithExternalIdentitySuccessfully(t)

	iTrue := true
	createOpts := users.CreateOpts{
		Name:      "jsmith",
		DomainID:  "1789d1",
		Enabled:   &iTrue,
		XuserType: "TenantIdp",
		XuserID:   "external-user-123",
	}

	actual, err := users.CreateUser(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CreatedUserWithExternalIdentity, *actual)
}

func TestModifyUserAdminWithExternalIdentity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleModifyUserAdminWithExternalIdentitySuccessfully(t)

	updateOpts := users.UpdateAdminOpts{
		Id:        "9fe1d3",
		XuserType: "TenantIdp",
		XuserId:   "external-user-456",
	}

	actual, err := users.ModifyUserAdmin(client.ServiceClient(), updateOpts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, UpdatedUserWithExternalIdentity, *actual)
}

func TestModifyUserAdminClearExternalIdentity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleModifyUserAdminClearExternalIdentitySuccessfully(t)

	updateOpts := users.UpdateAdminOpts{
		Id:        "9fe1d3",
		XuserType: "",
		XuserId:   "",
	}

	actual, err := users.ModifyUserAdmin(client.ServiceClient(), updateOpts)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, UpdatedUserWithClearedExternalIdentity, *actual)
}
