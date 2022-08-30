package testing

import (
	"testing"
	"time"

	volumeactions2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/volumeactions"

	"github.com/opentelekomcloud/gophertelekomcloud"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestAttach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockAttachResponse(t)

	options := volumeactions2.AttachOpts{
		MountPoint:   "/mnt",
		Mode:         "rw",
		InstanceUUID: "50902f4f-a974-46a0-85e9-7efc5e22dfdd",
	}
	err := volumeactions2.Attach(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options)
	th.AssertNoErr(t, err)
}

func TestBeginDetaching(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockBeginDetachingResponse(t)

	err := volumeactions2.BeginDetaching(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c")
	th.AssertNoErr(t, err)
}

func TestDetach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDetachResponse(t)

	err := volumeactions2.Detach(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", volumeactions2.DetachOpts{})
	th.AssertNoErr(t, err)
}

func TestUploadImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	MockUploadImageResponse(t)
	options := volumeactions2.UploadImageOpts{
		ContainerFormat: "bare",
		DiskFormat:      "raw",
		ImageName:       "test",
		Force:           true,
	}

	actual, err := volumeactions2.UploadImage(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options)
	th.AssertNoErr(t, err)

	expected := volumeactions2.VolumeImage{
		VolumeID:        "cd281d77-8217-4830-be95-9528227c105c",
		ContainerFormat: "bare",
		DiskFormat:      "raw",
		Description:     "",
		ImageID:         "ecb92d98-de08-45db-8235-bbafe317269c",
		ImageName:       "test",
		Size:            5,
		Status:          "uploading",
		UpdatedAt:       time.Date(2017, 7, 17, 9, 29, 22, 0, time.UTC),
		VolumeType: volumeactions2.ImageVolumeType{
			ID:          "b7133444-62f6-4433-8da3-70ac332229b7",
			Name:        "basic.ru-2a",
			Description: "",
			IsPublic:    true,
			ExtraSpecs:  map[string]interface{}{"volume_backend_name": "basic.ru-2a"},
			QosSpecsID:  "",
			Deleted:     false,
			DeletedAt:   time.Time{},
			CreatedAt:   time.Date(2016, 5, 4, 8, 54, 14, 0, time.UTC),
			UpdatedAt:   time.Date(2016, 5, 4, 9, 15, 33, 0, time.UTC),
		},
	}
	th.AssertDeepEquals(t, &expected, actual)
}

func TestReserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockReserveResponse(t)

	err := volumeactions2.Reserve(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c")
	th.AssertNoErr(t, err)
}

func TestUnreserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUnreserveResponse(t)

	err := volumeactions2.Unreserve(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c")
	th.AssertNoErr(t, err)
}

func TestInitializeConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockInitializeConnectionResponse(t)

	options := volumeactions2.InitializeConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: golangsdk.Disabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	_, err := volumeactions2.InitializeConnection(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options)
	th.AssertNoErr(t, err)
}

func TestTerminateConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockTerminateConnectionResponse(t)

	options := volumeactions2.TerminateConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: golangsdk.Enabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	err := volumeactions2.TerminateConnection(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options)
	th.AssertNoErr(t, err)
}

func TestExtendSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockExtendSizeResponse(t)

	options := volumeactions2.ExtendSizeOpts{
		NewSize: 3,
	}

	err := volumeactions2.ExtendSize(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options)
	th.AssertNoErr(t, err)
}

func TestForceDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockForceDeleteResponse(t)

	res := volumeactions2.ForceDelete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, res)
}
