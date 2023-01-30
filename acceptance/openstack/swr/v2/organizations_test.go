package v2

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func randomRepoName(prefix string, n int) string {
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	_, _ = rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return prefix + string(bytes)
}

func TestOrganizationWorkflow(t *testing.T) {
	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	name := randomRepoName("test-org", 6)
	opts := organizations.CreateOpts{Namespace: name}
	err = organizations.Create(client, opts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, organizations.Delete(client, name))
	})

	org, err := organizations.Get(client, name)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, org.Name)

	orgs, err := organizations.List(client, organizations.ListOpts{})
	th.AssertNoErr(t, err)
	found := false
	for _, o := range orgs {
		if o.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("can't find organization '%s' in the list", name)
	}
}

func TestOrganizationPermissionsWorkflow(t *testing.T) {
	userID := os.Getenv("OS_USER_ID_2")
	username := os.Getenv("OS_USERNAME_2")

	if username == "" || userID == "" {
		t.Skip("OS_USER_ID_2 and OS_USERNAME_2 should be set to test permission granting")
	}

	client, err := clients.NewSwrV2Client()
	th.AssertNoErr(t, err)

	// setup org
	orgName := fmt.Sprintf("repo-test-%d", tools.RandomInt(0, 0xf))
	dep := dependencies{t: t, client: client}
	dep.createOrganization(orgName)
	t.Cleanup(func() { dep.deleteOrganization(orgName) })

	auth := organizations.Auth{
		UserID:   userID,
		Username: username,
		Auth:     3,
	}

	err = organizations.CreatePermissions(client, orgName, []organizations.Auth{auth})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err := organizations.DeletePermissions(client, orgName, auth.UserID)
		th.AssertNoErr(t, err)
		perms, err := organizations.GetPermissions(client, orgName)
		th.AssertNoErr(t, err)
		assertAuthNotInPermissions(t, auth, perms)
	})

	perms, err := organizations.GetPermissions(client, orgName)
	th.AssertNoErr(t, err)
	assertAuthInPermissions(t, auth, perms)

	newAuth := organizations.Auth{
		UserID:   auth.UserID,
		Username: auth.Username,
		Auth:     1,
	}
	err = organizations.UpdatePermissions(client, orgName, []organizations.Auth{newAuth})
	th.AssertNoErr(t, err)

	updatedPerms, err := organizations.GetPermissions(client, orgName)
	th.AssertNoErr(t, err)
	assertAuthInPermissions(t, newAuth, updatedPerms)
}

func assertAuthInPermissions(t *testing.T, expected organizations.Auth, actual *organizations.Permissions) {
	if actual == nil {
		t.Fatal("actual organization permissions are nil")
	}

	for _, a := range actual.OthersAuth {
		if a.UserID == expected.UserID {
			if a.Username != expected.Username {
				t.Fatalf(
					"user ID is the same, but username differ - this is unexpected (%s/%s instead of %[1]s/%s)",
					a.UserID, a.Username, expected.Username,
				)
			}
			if a.Auth != expected.Auth {
				t.Fatalf("auth was %d, but %d expected", a.Auth, expected.Auth)
			}
			return
		}
	}
	t.Fatal("expected permission is not found in the `others_auth` list")
}

func assertAuthNotInPermissions(t *testing.T, expected organizations.Auth, actual *organizations.Permissions) {
	if actual == nil {
		t.Fatal("actual organization permissions are nil")
	}

	for _, a := range actual.OthersAuth {
		if a.UserID == expected.UserID {
			t.Fatalf("expected permission to be deleted, but it exist")
		}
	}
}
