package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/backup"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupList(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	th.AssertNoErr(t, err)

	backupList, err := backup.List(client, backup.ListOpts{})
	th.AssertNoErr(t, err)

	for _, element := range backupList {
		tools.PrintResource(t, element)
	}
}

func TestBackupLifeCycle(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	th.AssertNoErr(t, err)

	computeClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, computeClient, openstack.GetCloudServerCreateOpts(t))
	defer func() {
		openstack.DeleteCloudServer(t, computeClient, ecs.ID)
	}()

	backupName := tools.RandomString("backup-", 3)
	createOpts := backup.CreateOpts{
		BackupName:   backupName,
		Description:  "bla-bla",
		ResourceType: "OS::Nova::Server",
	}

	csbsBackup, err := backup.Create(client, ecs.ID, createOpts).ExtractBackup()
	th.AssertNoErr(t, err)
	defer func() {
		err := backup.Delete(client, csbsBackup.CheckpointId).ExtractErr()
		th.AssertNoErr(t, err)
	}()
}
