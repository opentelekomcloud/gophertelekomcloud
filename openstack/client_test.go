package openstack

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var (
	backupSuffix = ".backup"
	tmpl         = []byte(`
clouds:
  useless_cloud:
    auth:
      auth_url: "http://localhost/"
      password: "some-useless-passw0rd"
      username: "some-name"
`)
)

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
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, os.Remove(originalName))
		copyFile(t, backupFile, originalName)
		if err != nil && os.IsNotExist(err) {
			t.Logf("File %s doesn't exist, skipping", backupFile)
			continue
		}
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, os.Remove(originalName))
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

			th.AssertNoErr(subT, ioutil.WriteFile(fileName, tmpl, 0644))
			cloud, err := clients.EnvOS.Cloud()
			th.AssertNoErr(subT, err)
			th.AssertEquals(subT, "http://localhost/", cloud.AuthInfo.AuthURL)
			th.AssertEquals(subT, "some-useless-passw0rd", cloud.AuthInfo.Password)
			th.AssertEquals(subT, "some-name", cloud.AuthInfo.Username)
		})
	}
}
