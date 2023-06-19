package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/schedulerhints"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSchedulerHints(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volumeName := tools.RandomString("ACPTTEST", 16)
	createOpts := volumes.CreateOpts{
		Size:             1,
		Name:             volumeName,
		AvailabilityZone: clients.EnvOS.GetEnv("AVAILABILITY_ZONE"),
	}

	volume1, err := volumes.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, volume1.ID, "available", 60)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err := volumes.Delete(client, volumes.DeleteOpts{VolumeId: volume1.ID})
		th.AssertNoErr(t, err)
	})

	volumeName = tools.RandomString("ACPTTEST", 16)
	base := volumes.CreateOpts{
		Size: 1,
		Name: volumeName,
	}

	schedulerHints := schedulerhints.SchedulerHints{
		SameHost: []string{
			volume1.ID,
		},
	}

	createOptsWithHints := schedulerhints.CreateOptsExt{
		CreateOptsBuilder: base,
		SchedulerHints:    schedulerHints,
	}

	volume2, err := volumes.Create(client, createOptsWithHints)
	th.AssertNoErr(t, err)

	err = volumes.WaitForStatus(client, volume2.ID, "available", 60)
	th.AssertNoErr(t, err)

	err = volumes.Delete(client, volumes.DeleteOpts{VolumeId: volume2.ID})
	th.AssertNoErr(t, err)
}
