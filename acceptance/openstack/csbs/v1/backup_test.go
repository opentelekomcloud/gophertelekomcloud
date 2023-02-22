package v1

import (
	"testing"

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
		checkpoint := CreateCSBS(t, client, ecs.ID)

		csbsBackupList, err := backup.List(client, backup.ListOpts{
			CheckpointId: checkpoint.Id,
		})
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
