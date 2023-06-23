package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/backup"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/resource"
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
	t.Cleanup(func() { openstack.DeleteCloudServer(t, computeClient, ecs.ID) })

	t.Logf("Check if resource is protectable")
	query, err := resource.GetResBackupCapabilities(client, []resource.ResourceBackupCapOpts{
		{
			ResourceId:   ecs.ID,
			ResourceType: "OS::Nova::Server",
		},
	})
	th.AssertNoErr(t, err)

	if query[0].Result {
		t.Logf("Resource is protectable")
		backupName := tools.RandomString("backup-", 3)

		t.Logf("Attempting to create CSBS backup")
		checkpoint, err := backup.Create(client, ecs.ID, backup.CreateOpts{
			BackupName:   backupName,
			Description:  "bla-bla",
			ResourceType: "OS::Nova::Server",
		})
		th.AssertNoErr(t, err)

		t.Cleanup(func() {
			t.Logf("Attempting to delete CSBS backup: %s", checkpoint.Id)
			err = backup.Delete(client, checkpoint.Id)
			th.AssertNoErr(t, err)

			err = waitForBackupDeleted(client, 600, checkpoint.Id)
			th.AssertNoErr(t, err)
			t.Logf("Deleted CSBS backup: %s", checkpoint.Id)
		})

		csbsBackupList, err := backup.List(client, backup.ListOpts{
			CheckpointId: checkpoint.Id,
		})
		th.AssertNoErr(t, err)

		err = waitForBackupCreated(client, 600, csbsBackupList[0].Id)
		th.AssertNoErr(t, err)
		t.Logf("Created CSBS backup: %s", checkpoint.Id)

		capabilities, err := resource.GetResRestorationCapabilities(client, []resource.GetRestorationOpts{
			{
				CheckpointItemId: checkpoint.Id,
				Target: resource.RestorableTarget{
					ResourceId:   ecs.ID,
					ResourceType: "OS::Nova::Server",
					Volumes: []resource.RestoreVolumeMapping{
						{
							BackupId: csbsBackupList[0].Id,
							VolumeId: ecs.VolumeAttached[0].ID,
						},
					},
				},
			},
		})
		th.AssertNoErr(t, err)
		tools.PrintResource(t, capabilities)
	} else {
		t.Logf("Resource isn't protectable")
	}
}

func waitForBackupCreated(client *golangsdk.ServiceClient, secs int, backupID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		csbsBackup, err := backup.Get(client, backupID)
		if err != nil {
			return false, err
		}

		if csbsBackup.Id == "error" {
			return false, err
		}

		if csbsBackup.Status == "available" {
			return true, nil
		}

		return false, nil
	})
}

func waitForBackupDeleted(client *golangsdk.ServiceClient, secs int, backupID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := backup.Get(client, backupID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}
		return false, nil
	})
}
