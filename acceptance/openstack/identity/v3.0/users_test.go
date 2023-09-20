package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/users"
	oldusers "github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUserLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	oldClient, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	createOpts := users.CreateOpts{
		Name:     tools.RandomString("user-name-", 4),
		Enabled:  pointerto.Bool(true),
		DomainID: client.DomainID,
	}

	user, err := users.CreateUser(client, createOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	t.Cleanup(func() {
		err = oldusers.Delete(oldClient, user.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	th.AssertEquals(t, createOpts.Name, user.Name)
	th.AssertEquals(t, *createOpts.Enabled, user.Enabled)

	userGet, err := users.GetUser(client, user.ID)
	if err != nil {
		t.Fatalf("Unable to retrieve user: %v", err)
	}

	th.AssertEquals(t, userGet.Name, user.Name)
	th.AssertEquals(t, userGet.Enabled, user.Enabled)
	th.AssertEquals(t, userGet.Email, user.Email)
	th.AssertEquals(t, userGet.DomainID, user.DomainID)

	updateOpts := users.UpdateOpts{
		Enabled:  pointerto.Bool(false),
		Name:     tools.RandomString("new-user-name-", 4),
		Password: tools.RandomString("Hello-world-", 5),
	}

	userUpdate, err := users.ModifyUser(client, user.ID, updateOpts)
	if err != nil {
		t.Fatalf("Unable to update user info: %v", err)
	}

	th.AssertEquals(t, userUpdate.Name, updateOpts.Name)
	th.AssertEquals(t, userUpdate.Enabled, *updateOpts.Enabled)
	th.AssertEquals(t, userUpdate.Email, updateOpts.Email)
	th.AssertEquals(t, userUpdate.DomainID, userGet.DomainID)
}
