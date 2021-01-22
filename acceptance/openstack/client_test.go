package openstack

import (
	"fmt"
	"io/ioutil"
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
	backupSuffix = ".backup"
	tmpl         = `
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

// copyFile copies file if it exists
func copyFile(t *testing.T, src, dest string) {
	fileStat, err := os.Stat(src)
	if err != nil && os.IsNotExist(err) {
		t.Logf("File %s doesn't exist, skipping", src)
		return
	}

	data, err := ioutil.ReadFile(src)
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, ioutil.WriteFile(dest, data, fileStat.Mode()))
}

// backupFile creates copy of the file and return path to the copy
func backupFiles(t *testing.T, originals ...string) {
	for _, file := range originals {
		backupFile := fmt.Sprintf("%s%s", file, backupSuffix)
		copyFile(t, file, backupFile)
	}
}

// restoreBackup replaces files with the backups
func restoreBackup(t *testing.T, files ...string) {
	for _, originalName := range files {
		backupFile := fmt.Sprintf("%s%s", originalName, backupSuffix)
		_, err := os.Stat(originalName)
		if err != nil && os.IsNotExist(err) {
			t.Logf("File %s doesn't exist, skipping", originalName)
			continue
		}
		th.AssertNoErr(t, os.Remove(originalName))
		copyFile(t, backupFile, originalName)
		th.AssertNoErr(t, os.Remove(backupFile))
	}
}

func TestCloudYamlPaths(t *testing.T) {
	_ = os.Setenv("OS_CLOUD", "useless_cloud")
	home, _ := os.UserHomeDir()
	cwd, _ := os.Getwd()

	fileName := "clouds.yaml"
	currentConfigDir := filepath.Join(cwd, fileName)
	userConfigDir := filepath.Join(home, ".config/openstack", fileName)
	unixConfigDir := filepath.Join("/etc/openstack", fileName)
	files := []string{currentConfigDir, userConfigDir, unixConfigDir}
	backupFiles(t, currentConfigDir, userConfigDir, unixConfigDir)
	defer restoreBackup(t, files...)

	for _, fileName := range files {
		t.Run(fmt.Sprintf("Config at %s", fileName), func(subT *testing.T) {
			if runtime.GOOS == "windows" && fileName == unixConfigDir {
				subT.Skip("Skipping on windows")
			}

			dir := filepath.Dir(fileName)
			if err := os.MkdirAll(dir, 0755); err != nil { // make sure that dir exists
				if os.IsPermission(err) {
					subT.Skip(err.Error())
				}
				th.AssertNoErr(t, err)
			}

			th.AssertNoErr(subT, writeYamlFile(tmpl, fileName))
			cloud, err := clients.EnvOS.Cloud()
			th.AssertNoErr(subT, err)
			th.AssertEquals(subT, "http://localhost/", cloud.AuthInfo.AuthURL)
			th.AssertEquals(subT, "some-useless-passw0rd", cloud.AuthInfo.Password)
			th.AssertEquals(subT, "some-name", cloud.AuthInfo.Username)
		})
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
