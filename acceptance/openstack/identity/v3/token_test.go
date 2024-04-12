package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/tokens"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGetToken(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	authOptions := tokens.AuthOptions{
		UserID:     cc.ProviderClient.UserID,
		Password:   cc.AuthInfo.Password,
		DomainName: cc.AuthInfo.DomainName,
		Passcode:   os.Getenv("OS_PASSCODE"),
	}

	if cc.AuthInfo.Password == "" {
		t.Skip("password auth is required for this test")
	}

	result := tokens.Create(client, &authOptions)
	token, err := result.Extract()
	th.AssertNoErr(t, err)

	_, err = tokens.Get(client, token.ID).ExtractServiceCatalog()
	th.AssertNoErr(t, err)

	user, err := tokens.Get(client, token.ID).ExtractUser()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, authOptions.UserID, user.ID)

	roles, err := tokens.Get(client, token.ID).ExtractRoles()
	th.AssertNoErr(t, err)
	if len(roles) == 0 {
		t.Fatalf("user has no roles")
	}

	_, err = tokens.Get(client, token.ID).ExtractProject()
	th.AssertNoErr(t, err)
}
