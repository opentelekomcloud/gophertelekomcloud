package gpfs

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gpfs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGPFSFileSystemLifecycle(t *testing.T) {
	client, err := clients.NewGPFSClient()
	th.AssertNoErr(t, err)

	fsName := strings.ToLower(tools.RandomString("gpfs-sdk-test", 5))

	_, err = client.CreateFS(&gpfs.CreateFSInput{
		FSName:     fsName,
		Redundancy: "3az",
		BucketType: "SFS",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		_, err = client.DeleteFS(fsName)
		th.AssertNoErr(t, err)
	})

	fileSystems, err := client.ListFS(&gpfs.ListFSInput{
		BucketType: "SFS",
	})
	th.AssertNoErr(t, err)
	th.AssertNotEquals(t, 0, len(fileSystems.Buckets))
}
