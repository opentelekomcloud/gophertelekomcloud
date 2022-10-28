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

func TestMFA(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV3AdminClient()
	th.AssertNoErr(t, err)

	user, err := users.Create(client, users.CreateOpts{
		Name:    tools.RandomString("user-name-", 4),
		Enabled: pointerto.Bool(true),
		Email:   "test-email@mail.com",
	}).Extract()
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}

	t.Cleanup(func() {
		err = users.Delete(client, user.ID).ExtractErr()
		th.AssertNoErr(t, err)
	})

	mfa, err := users.CreateMfaDevice(client, users.CreateMfaDeviceOpts{
		Name:   "test-mfa",
		UserId: user.ID,
	})
	t.Cleanup(func() {
		err = users.DeleteMfaDevice(client, users.DeleteMfaDeviceOpts{
			UserId:       user.ID,
			SerialNumber: mfa.SerialNumber,
		})
		th.AssertNoErr(t, err)
	})

	th.AssertNoErr(t, err)
	t.Logf("MFA device created: %v", mfa)

	err = users.CreateBindingDevice(client, users.BindMfaDevice{
		UserId:                   user.ID,
		SerialNumber:             mfa.SerialNumber,
		AuthenticationCodeFirst:  "977931",
		AuthenticationCodeSecond: "527347",
	})
	th.AssertNoErr(t, err)

	err = users.DeleteBindingDevice(client, users.UnbindMfaDevice{
		UserId:             user.ID,
		AuthenticationCode: "977931",
		SerialNumber:       mfa.SerialNumber,
	})

	get, err := users.ShowUserMfaDevice(client, user.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, get.SerialNumber, mfa.SerialNumber)

	list, err := users.ListUserMfaDevices(client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)
}
