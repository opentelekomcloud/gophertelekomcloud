package openstack

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	fileName = "./clouds.yaml"
	tmpl     = `
clouds:
  useless_cloud:
    auth:
      auth_url: "http://localhost/"
      password: "some-useless-passw0rd"
      username: "some-name"
`
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

func TestCloudYamlPaths(t *testing.T) {
	_ = os.Setenv("OS_CLOUD", "useless_cloud")
	home, _ := os.UserHomeDir()
	cwd, _ := os.Getwd()

	currentConfigDir, _ := filepath.Abs(filepath.Join(cwd, "clouds.yaml"))
	userConfigDir, _ := filepath.Abs(filepath.Join(home, ".config/openstack/clouds.yaml"))
	unixConfigDir, _ := filepath.Abs("/etc/openstack/clouds.yaml")

	if _, err := os.Stat(currentConfigDir); os.IsNotExist(err) {
		err := writeYamlFile(tmpl, currentConfigDir)
		if err != nil {
			th.AssertNoErr(t, err)
		}
		defer func() { _ = os.Remove(currentConfigDir) }()
		cloud, err := clients.EnvOS.Cloud()
		if err != nil {
			th.AssertNoErr(t, err)
		}
		th.AssertEquals(t, "http://localhost/", cloud.AuthInfo.AuthURL)
		th.AssertEquals(t, "some-useless-passw0rd", cloud.AuthInfo.Password)
		th.AssertEquals(t, "some-name", cloud.AuthInfo.Username)
	}

	if _, err := os.Stat(userConfigDir); os.IsNotExist(err) {
		err := writeYamlFile(tmpl, userConfigDir)
		if err != nil {
			th.AssertNoErr(t, err)
		}
		defer func() { _ = os.Remove(userConfigDir) }()
		cloud, err := clients.EnvOS.Cloud()
		if err != nil {
			th.AssertNoErr(t, err)
		}
		th.AssertEquals(t, "http://localhost/", cloud.AuthInfo.AuthURL)
		th.AssertEquals(t, "some-useless-passw0rd", cloud.AuthInfo.Password)
		th.AssertEquals(t, "some-name", cloud.AuthInfo.Username)
	}

	if runtime.GOOS != "windows" {
		if _, err := os.Stat(unixConfigDir); os.IsNotExist(err) {
			err := writeYamlFile(tmpl, unixConfigDir)
			if err != nil {
				th.AssertNoErr(t, err)
			}
			defer func() { _ = os.Remove(unixConfigDir) }()
			cloud, err := clients.EnvOS.Cloud()
			if err != nil {
				th.AssertNoErr(t, err)
			}
			th.AssertEquals(t, "http://localhost/", cloud.AuthInfo.AuthURL)
			th.AssertEquals(t, "some-useless-passw0rd", cloud.AuthInfo.Password)
			th.AssertEquals(t, "some-name", cloud.AuthInfo.Username)
		}
	}
}

func TestAuthTokenNoRegion(t *testing.T) {
	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	envPrefix := tools.RandomString("", 5)
	th.AssertNoErr(t, os.Setenv(envPrefix+"_TOKEN", cc.TokenID))
	th.AssertNoErr(t, os.Setenv(envPrefix+"_AUTH_URL", cc.IdentityEndpoint))

	env := openstack.NewEnv(envPrefix)
	client, err := env.AuthenticatedClient()
	th.AssertNoErr(t, err)
	_, err = openstack.NewComputeV2(client, golangsdk.EndpointOpts{})
	th.AssertNoErr(t, err)
}

func TestReAuth(t *testing.T) {
	cloud, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	opts, err := openstack.AuthOptionsFromInfo(&cloud.AuthInfo, cloud.AuthType)
	th.AssertNoErr(t, err)

	ao := opts.(golangsdk.AuthOptions)
	ao.AllowReauth = true

	scl, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(scl, golangsdk.EndpointOpts{
		Region: cloud.RegionName,
	})
	th.AssertNoErr(t, err)

	t.Logf("Sleeping for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Attempting to reauthenticate")

	th.AssertNoErr(t, scl.ReauthFunc())

	t.Logf("Creating a compute client")
	_, err = openstack.NewComputeV2(scl, golangsdk.EndpointOpts{
		Region: cloud.RegionName,
	})
	th.AssertNoErr(t, err)
}
