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
	client, err := clients.NewObjectStorageV1Client()
	th.AssertNoErr(t, err)
	t.Logf("Located a storage service at endpoint: [%s]", client.Endpoint)
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

func TestReauth(t *testing.T) {
	setEnvToken(t)

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
