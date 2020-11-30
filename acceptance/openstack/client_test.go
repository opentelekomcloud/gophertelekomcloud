package openstack

import (
	"os"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/utils"
)

func TestAuthenticatedClient(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	if client.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	if client.ProjectID == "" {
		t.Errorf("Project ID is not set for the client")
	}
	if client.UserID == "" {
		t.Errorf("User ID is not set for the client")
	}
	if client.DomainID == "" {
		t.Errorf("Domain ID is not set for the client")
	}

	t.Logf("Client successfully acquired a token: %v", client.TokenID)

	// Find the storage service in the service catalog.
	storage, err := openstack.NewObjectStorageV1(client, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao)})
	if err != nil {
		t.Errorf("Unable to locate a storage service: %v", err)
	} else {
		t.Logf("Located a storage service at endpoint: [%s]", storage.Endpoint)
	}
}

func TestAuthTokenNoRegion(t *testing.T) {
	osEnv := openstack.NewEnv("OS_")
	preClient, err := osEnv.AuthenticatedClient()
	if err != nil {
		t.Fatalf("Failed to auth client: %s", err)
	}

	envPrefix := tools.RandomString("", 5)
	if err := os.Setenv(envPrefix+"_TOKEN", preClient.Token()); err != nil {
		t.Errorf("Failed to set token: %s", err)
	}
	if err := os.Setenv(envPrefix+"_AUTH_URL", preClient.IdentityEndpoint); err != nil {
		t.Errorf("Failed to set auth url: %s", err)
	}

	env := openstack.NewEnv(envPrefix)
	client, err := env.AuthenticatedClient()
	if err != nil {
		t.Errorf("Failed to auth client: %s", err)
	}
	_, err = openstack.NewComputeV2(client, golangsdk.EndpointOpts{})
	if err != nil {
		t.Errorf("Failed to get compute client: %s", err)
	}
}

func TestReauth(t *testing.T) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to obtain environment auth options: %v", err)
	}

	// Allow reauth
	ao.AllowReauth = true

	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		t.Fatalf("Unable to create provider: %v", err)
	}

	err = openstack.Authenticate(provider, ao)
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
	if err != nil {
		t.Fatalf("Unable to create compute client: %v", err)
	}

	t.Logf("Sleeping for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to reauthenticate")

	err = provider.ReauthFunc()
	if err != nil {
		t.Fatalf("Unable to reauthenticate: %v", err)
	}

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: utils.GetRegion(ao),
	})
	if err != nil {
		t.Fatalf("Unable to create compute client: %v", err)
	}
}
