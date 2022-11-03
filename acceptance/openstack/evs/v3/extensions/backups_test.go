package extensions

import (
	"testing"

	backups2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/extensions/backups"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	blockstorage "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/evs/v3"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupsCRUD(t *testing.T) {
	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorage.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	defer blockstorage.DeleteVolume(t, blockClient, volume)

	backup, err := CreateBackup(t, blockClient, volume.ID)
	th.AssertNoErr(t, err)
	defer DeleteBackup(t, blockClient, backup.ID)

	allPages, err := backups2.List(blockClient, nil).AllPages()
	th.AssertNoErr(t, err)

	allBackups, err := backups2.ExtractBackups(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allBackups {
		if backup.Name == v.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
