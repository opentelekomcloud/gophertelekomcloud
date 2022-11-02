package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	blockstorage "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/evs/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/volumetenants"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVolumeTenants(t *testing.T) {
	type volumeWithTenant struct {
		volumes.Volume
		volumetenants.VolumeTenantExt
	}

	var allVolumes []volumeWithTenant

	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	listOpts := volumes.ListOpts{
		Name: "I SHOULD NOT EXIST",
	}
	allPages, err := volumes.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(allVolumes))

	volume1, err := blockstorage.CreateVolume(t, client)
	th.AssertNoErr(t, err)
	defer blockstorage.DeleteVolume(t, client, volume1)

	allPages, err = volumes.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(allVolumes) > 0)
}
