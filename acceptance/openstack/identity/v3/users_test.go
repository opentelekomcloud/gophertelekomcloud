package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUsersList(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	listOpts := users.ListOpts{}

	allPages, err := users.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list users: %v", err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		t.Fatalf("Unable to extract users: %v", err)
	}

	for _, user := range allUsers {
		if len(user.Name) < 5 {
			t.Fatalf("Invalid user name")
		}
	}
}

func TestUserLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	createOpts := users.CreateOpts{
		Name:    tools.RandomString("user-name-", 4),
		Enabled: pointerto.Bool(true),
		Email:   "test-email@mail.com",
	}

	user, err := users.Create(client, createOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	t.Cleanup(func() {
		err = users.Delete(client, user.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	th.AssertEquals(t, createOpts.Name, user.Name)
	th.AssertEquals(t, *createOpts.Enabled, user.Enabled)
	th.AssertEquals(t, createOpts.Email, user.Email)

	userGet, err := users.Get(client, user.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve user: %v", err)
	}

	th.AssertEquals(t, userGet.Name, user.Name)
	th.AssertEquals(t, userGet.Enabled, user.Enabled)
	th.AssertEquals(t, userGet.Email, user.Email)
	th.AssertEquals(t, userGet.DomainID, user.DomainID)
	th.AssertEquals(t, userGet.DefaultProjectID, user.DefaultProjectID)

	updateOpts := users.UpdateOpts{
		Enabled:  pointerto.Bool(false),
		Name:     tools.RandomString("new-user-name-", 4),
		Password: tools.RandomString("Hello-world-", 4),
		Email:    "new-test-email@mail.com",
	}

	userUpdate, err := users.Update(client, user.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update user info: %v", err)
	}

	th.AssertEquals(t, userUpdate.Name, updateOpts.Name)
	th.AssertEquals(t, userUpdate.Enabled, *updateOpts.Enabled)
	th.AssertEquals(t, userUpdate.Email, updateOpts.Email)
	th.AssertEquals(t, userUpdate.DomainID, userGet.DomainID)
	th.AssertEquals(t, userUpdate.DefaultProjectID, userGet.DefaultProjectID)
}
