package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMFA(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	list, err := users.ListUserMfaDevices(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)

	get, err := users.ShowUserMfaDevice(client, client.UserID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, get.UserId, client.UserID)
}
