package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/backups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupWorkflow(t *testing.T) {
	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	// Create RDSv3 instance
	rds := createRDS(t, client, cc.RegionName)
	t.Cleanup(func() {
		deleteRDS(t, client, rds.Id)
	})

	b, err := backups.Create(client, backups.CreateOpts{
		InstanceID: rds.Id,
		Name:       tools.RandomString("rds-backup-", 5),
	}).Extract()
	th.AssertNoErr(t, err)
	t.Log("Backup creation started")

	t.Cleanup(func() {
		th.AssertNoErr(t, backups.Delete(client, b.ID).ExtractErr())
		t.Log("Backup deleted")
	})

	err = backups.WaitForBackup(client, rds.Id, b.ID, backups.StatusCompleted)
	th.AssertNoErr(t, err)
	t.Log("Backup creation complete")
}
