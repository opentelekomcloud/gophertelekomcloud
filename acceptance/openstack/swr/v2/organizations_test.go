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
	err = organizations.Create(client, opts).ExtractErr()
	th.AssertNoErr(t, err)
	defer func() {
		th.AssertNoErr(t, organizations.Delete(client, name).ExtractErr())
	}()

	org, err := organizations.Get(client, name).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, name, org.Name)

	pages, err := organizations.List(client, nil).AllPages()
	th.AssertNoErr(t, err)
	orgs, err := organizations.ExtractOrganizations(pages)
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
	defer dep.deleteOrganization(orgName)
	//

	auth := organizations.Auth{
		UserID:   userID,
		Username: username,
		Auth:     3,
	}

	err = organizations.CreatePermissions(client, orgName, organizations.CreatePermissionsOpts(auth)).ExtractErr()
	th.AssertNoErr(t, err)

	defer func() {
		err := organizations.DeletePermissions(client, orgName, auth.UserID).ExtractErr()
		th.AssertNoErr(t, err)
		perms, err := organizations.GetPermissions(client, orgName).Extract()
		th.AssertNoErr(t, err)
		assertAuthNotInPermissions(t, auth, perms)
	}()

	perms, err := organizations.GetPermissions(client, orgName).Extract()
	th.AssertNoErr(t, err)
	assertAuthInPermissions(t, auth, perms)

	newAuth := organizations.Auth{
		UserID:   auth.UserID,
		Username: auth.Username,
		Auth:     1,
	}
	err = organizations.UpdatePermissions(client, orgName, organizations.UpdatePermissionsOpts(newAuth)).ExtractErr()
	th.AssertNoErr(t, err)

	updatedPerms, err := organizations.GetPermissions(client, orgName).Extract()
	th.AssertNoErr(t, err)
	assertAuthInPermissions(t, newAuth, updatedPerms)
}

func assertAuthInPermissions(t *testing.T, expected organizations.Auth, actual *organizations.OrganizationPermissions) {
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

func assertAuthNotInPermissions(t *testing.T, expected organizations.Auth, actual *organizations.OrganizationPermissions) {
	if actual == nil {
		t.Fatal("actual organization permissions are nil")
	}

	for _, a := range actual.OthersAuth {
		if a.UserID == expected.UserID {
			t.Fatalf("expected permission to be deleted, but it exist")
		}
	}
}
