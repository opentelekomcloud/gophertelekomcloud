package extensions

import (
	"fmt"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/images"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/volumeactions"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
	v3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
)

// CreateUploadImage will upload volume it as volume-baked image. A name of new image or err will be
// returned
func CreateUploadImage(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) (*volumeactions.VolumeImage, error) {
	if testing.Short() {
		t.Skip("Skipping test that requires volume-backed image uploading in short mode.")
	}

	imageName := tools.RandomString("ACPTTEST", 16)
	uploadImageOpts := volumeactions.UploadImageOpts{
		ImageName: imageName,
		Force:     true,
	}

	volumeImage, err := volumeactions.UploadImage(client, volume.ID, uploadImageOpts)
	if err != nil {
		return volumeImage, err
	}

	t.Logf("Uploading volume %s as volume-backed image %s", volume.ID, imageName)

	if err := volumes.WaitForStatus(client, volume.ID, "available", 60); err != nil {
		return volumeImage, err
	}

	t.Logf("Uploaded volume %s as volume-backed image %s", volume.ID, imageName)

	return volumeImage, nil

}

// DeleteUploadedImage deletes uploaded image. An error will be returned
// if the deletion request failed.
func DeleteUploadedImage(t *testing.T, client *golangsdk.ServiceClient, imageID string) error {
	if testing.Short() {
		t.Skip("Skipping test that requires volume-backed image removing in short mode.")
	}

	t.Logf("Removing image %s", imageID)

	err := images.Delete(client, imageID).ExtractErr()
	if err != nil {
		return err
	}

	return nil
}

// CreateVolumeAttach will attach a volume to an instance. An error will be
// returned if the attachment failed.
func CreateVolumeAttach(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume, server *servers.Server) error {
	if testing.Short() {
		t.Skip("Skipping test that requires volume attachment in short mode.")
	}

	attachOpts := volumeactions.AttachOpts{
		MountPoint:   "/mnt",
		Mode:         "rw",
		InstanceUUID: server.ID,
	}

	t.Logf("Attempting to attach volume %s to server %s", volume.ID, server.ID)

	if err := volumeactions.Attach(client, volume.ID, attachOpts); err != nil {
		return err
	}

	if err := volumes.WaitForStatus(client, volume.ID, "in-use", 60); err != nil {
		return err
	}

	t.Logf("Attached volume %s to server %s", volume.ID, server.ID)

	return nil
}

// CreateVolumeReserve creates a volume reservation. An error will be returned
// if the reservation failed.
func CreateVolumeReserve(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) error {
	if testing.Short() {
		t.Skip("Skipping test that requires volume reservation in short mode.")
	}

	t.Logf("Attempting to reserve volume %s", volume.ID)

	if err := volumeactions.Reserve(client, volume.ID); err != nil {
		return err
	}

	t.Logf("Reserved volume %s", volume.ID)

	return nil
}

// DeleteVolumeAttach will detach a volume from an instance. A fatal error will
// occur if the snapshot failed to be deleted. This works best when used as a
// deferred function.
func DeleteVolumeAttach(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) {
	t.Logf("Attepting to detach volume volume: %s", volume.ID)

	detachOpts := volumeactions.DetachOpts{
		AttachmentID: volume.Attachments[0].AttachmentID,
	}

	if err := volumeactions.Detach(client, volume.ID, detachOpts); err != nil {
		t.Fatalf("Unable to detach volume %s: %v", volume.ID, err)
	}

	if err := volumes.WaitForStatus(client, volume.ID, "available", 60); err != nil {
		t.Fatalf("Volume %s failed to become unavailable in 60 seconds: %v", volume.ID, err)
	}

	t.Logf("Detached volume: %s", volume.ID)
}

// DeleteVolumeReserve deletes a volume reservation. A fatal error will occur
// if the deletion request failed. This works best when used as a deferred
// function.
func DeleteVolumeReserve(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) {
	if testing.Short() {
		t.Skip("Skipping test that requires volume reservation in short mode.")
	}

	t.Logf("Attempting to unreserve volume %s", volume.ID)

	if err := volumeactions.Unreserve(client, volume.ID); err != nil {
		t.Fatalf("Unable to unreserve volume %s: %v", volume.ID, err)
	}

	t.Logf("Unreserved volume %s", volume.ID)
}

// ExtendVolumeSize will extend the size of a volume.
func ExtendVolumeSize(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) error {
	t.Logf("Attempting to extend the size of volume %s", volume.ID)

	extendOpts := volumeactions.ExtendSizeOpts{
		NewSize: 2,
	}

	err := volumeactions.ExtendSize(client, volume.ID, extendOpts)
	if err != nil {
		return err
	}

	if err := volumes.WaitForStatus(client, volume.ID, "available", 60); err != nil {
		return err
	}

	return nil
}

// SetBootable will set a bootable status to a volume.
func SetBootable(t *testing.T, client *golangsdk.ServiceClient, volume *volumes.Volume) error {
	t.Logf("Attempting to apply bootable status to volume %s", volume.ID)

	bootableOpts := volumeactions.BootableOpts{
		Bootable: true,
	}

	err := volumeactions.SetBootable(client, volume.ID, bootableOpts)
	if err != nil {
		return err
	}

	vol, err := v3.Get(client, volume.ID)
	if err != nil {
		return err
	}

	if strings.ToLower(vol.Bootable) != "true" {
		return fmt.Errorf("Volume bootable status is %q, expected 'true'", vol.Bootable)
	}

	bootableOpts = volumeactions.BootableOpts{
		Bootable: false,
	}

	err = volumeactions.SetBootable(client, volume.ID, bootableOpts)
	if err != nil {
		return err
	}

	vol, err = v3.Get(client, volume.ID)
	if err != nil {
		return err
	}

	if strings.ToLower(vol.Bootable) == "true" {
		return fmt.Errorf("Volume bootable status is %q, expected 'false'", vol.Bootable)
	}

	return nil
}
