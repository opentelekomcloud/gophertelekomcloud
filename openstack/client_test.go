package openstack

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

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

func removeIfExist(t *testing.T, path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		t.Logf("File %s doesn't exist, skipping", path)
		return
	}
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, os.Remove(path))
}

// restoreBackup replaces files with the backups
func restoreBackup(t *testing.T, files ...string) {
	for _, original := range files {
		backup := fmt.Sprintf("%s%s", original, backupSuffix)
		removeIfExist(t, original)
		copyFile(t, backup, original)
		removeIfExist(t, backup)
	}
}

func checkLazyness(t *testing.T, env Env, expected bool) {
	authUrl0 := "http://url:0"
	_ = os.Setenv("OS_AUTH_URL", authUrl0)
	cloud0, err := env.Cloud()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, authUrl0, cloud0.AuthInfo.AuthURL)

	authUrl1 := "http://url:1"
	_ = os.Setenv("OS_AUTH_URL", authUrl1)
	cloud1, err := env.Cloud()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, expected, authUrl0 == cloud1.AuthInfo.AuthURL)
	th.AssertEquals(t, !expected, authUrl1 == cloud1.AuthInfo.AuthURL)
}

func TestLazyEnv(t *testing.T) {
	t.Run("lazy", func(sub *testing.T) {
		env := NewEnv("OS_", true)
		checkLazyness(sub, env, true)
	})
	t.Run("not lazy", func(sub *testing.T) {
		env := NewEnv("OS_", false)
		checkLazyness(sub, env, false)
	})
	t.Run("default", func(sub *testing.T) {
		env := NewEnv("OS_")
		checkLazyness(sub, env, true)
		sub.Log("Lazy by default")
	})
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
			cloud, err := NewEnv("OS_").Cloud()
			th.AssertNoErr(subT, err)
			th.AssertEquals(subT, "http://localhost/", cloud.AuthInfo.AuthURL)
			th.AssertEquals(subT, "some-useless-passw0rd", cloud.AuthInfo.Password)
			th.AssertEquals(subT, "some-name", cloud.AuthInfo.Username)
		})
	}
}
