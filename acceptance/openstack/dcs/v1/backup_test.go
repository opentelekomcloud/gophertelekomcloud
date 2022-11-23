package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v1/backups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsBackupLifeCycle(t *testing.T) {
	client, err := clients.NewDcsV1Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)
	backupId, err := backups.BackupInstance(client, dcsInstance.InstanceID, backups.BackupInstanceOpts{Remark: "test"})
	t.Cleanup(func() {
		err := backups.DeleteBackupFile(client, dcsInstance.InstanceID, backupId)
		th.AssertNoErr(t, err)
	})

	th.AssertNoErr(t, err)
	t.Logf("Created DCSv1 backup: %s", backupId)

	backupList, err := backups.ListBackupRecords(client, dcsInstance.InstanceID, backups.ListBackupOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, backupList.TotalNum, 1)

	restoreId, err := backups.RestoreInstance(client, dcsInstance.InstanceID, backups.RestoreInstanceOpts{BackupId: backupId, Remark: "test"})
	th.AssertNoErr(t, err)
	t.Logf("Restored DCSv1 backup: %s", restoreId)

	restoreList, err := backups.ListRestoreRecords(client, dcsInstance.InstanceID, backups.ListBackupOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, restoreList.TotalNum, 1)
}
