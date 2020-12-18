package openstack

import (
	"os"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAuthenticatedClient(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	if cc.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	if cc.ProjectID == "" {
		t.Errorf("Project ID is not set for the client")
	}
	if cc.UserID == "" {
		t.Errorf("User ID is not set for the client")
	}
	if cc.DomainID == "" {
		t.Errorf("Domain ID is not set for the client")
	}

	// Find the storage service in the service catalog.
	storage, err := openstack.NewObjectStorageV1(cc.ProviderClient, golangsdk.EndpointOpts{
		Region: cc.RegionName,
	})
	th.AssertNoErr(t, err)
	t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)
}

func TestAuthTokenNoRegion(t *testing.T) {
	osEnv := openstack.NewEnv("OS_")
	preClient, err := osEnv.AuthenticatedClient()
	th.AssertNoErr(t, err)

	envPrefix := tools.RandomString("", 5)
	th.AssertNoErr(t, os.Setenv(envPrefix+"_TOKEN", preClient.IdentityEndpoint))
	th.AssertNoErr(t, os.Setenv(envPrefix+"_AUTH_URL", preClient.IdentityEndpoint))

	env := openstack.NewEnv(envPrefix)
	client, err := env.AuthenticatedClient()
	th.AssertNoErr(t, err)
	_, err = openstack.NewComputeV2(client, golangsdk.EndpointOpts{})
	th.AssertNoErr(t, err)
}

func setEnvToken(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)
	if cc.AuthInfo.Token == "" {
		t.Fatalf("No token is set in client")
	}
	th.AssertNoErr(t, os.Setenv("OS_TOKEN", cc.AuthInfo.Token))
}

func unsetEnvToken(t *testing.T) {
	th.AssertNoErr(t, os.Unsetenv("OS_TOKEN"))
}

func TestReauth(t *testing.T) {
	setEnvToken(t)
	defer unsetEnvToken(t)

	ao, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	// Allow reauth
	ao.AllowReauth = true

	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	th.AssertNoErr(t, err)

	err = openstack.Authenticate(provider, ao)
	th.AssertNoErr(t, err)

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
	th.AssertNoErr(t, err)

	t.Logf("Sleeping for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to reauthenticate")

	th.AssertNoErr(t, provider.ReauthFunc())

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
	th.AssertNoErr(t, err)
}
