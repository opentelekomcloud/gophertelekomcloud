package extensions

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	compute "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/compute/v2"
	blockstorageV3 "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/evs/v3"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestVolumeActionsUploadImageDestroy(t *testing.T) {
	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorageV3.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { blockstorageV3.DeleteVolume(t, blockClient, volume) })

	volumeImage, err := CreateUploadImage(t, blockClient, volume)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, volumeImage)

	err = DeleteUploadedImage(t, computeClient, volumeImage.ImageID)
	th.AssertNoErr(t, err)
}

func TestVolumeActionsAttachCreateDestroy(t *testing.T) {
	// TODO
	t.Skip("images.ListDetail discarded and not working.")

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	computeClient, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := compute.CreateServer(t, computeClient)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { compute.DeleteServer(t, computeClient, server) })

	volume, err := blockstorageV3.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { blockstorageV3.DeleteVolume(t, blockClient, volume) })

	err = CreateVolumeAttach(t, blockClient, volume, server)
	th.AssertNoErr(t, err)

	newVolume, err := volumes.Get(blockClient, volume.ID)
	th.AssertNoErr(t, err)

	DeleteVolumeAttach(t, blockClient, newVolume)
}

func TestVolumeActionsReserveUnreserve(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorageV3.CreateVolume(t, client)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { blockstorageV3.DeleteVolume(t, client, volume) })

	err = CreateVolumeReserve(t, client, volume)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { DeleteVolumeReserve(t, client, volume) })
}

func TestVolumeActionsExtendSize(t *testing.T) {
	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorageV3.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { blockstorageV3.DeleteVolume(t, blockClient, volume) })

	tools.PrintResource(t, volume)

	err = ExtendVolumeSize(t, blockClient, volume)
	th.AssertNoErr(t, err)

	newVolume, err := volumes.Get(blockClient, volume.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newVolume)
}

func TestVolumeActionsSetBootable(t *testing.T) {
	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	volume, err := blockstorageV3.CreateVolume(t, blockClient)
	th.AssertNoErr(t, err)
	t.Cleanup(func() { blockstorageV3.DeleteVolume(t, blockClient, volume) })

	err = SetBootable(t, blockClient, volume)
	th.AssertNoErr(t, err)
}
