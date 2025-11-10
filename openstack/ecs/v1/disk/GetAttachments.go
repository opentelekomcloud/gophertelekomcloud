package disk

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetAttachments(client *golangsdk.ServiceClient, serverID string) (*Attachments, error) {
	// GET /v1/{project_id}/cloudservers/{server_id}/block_device
	raw, err := client.Get(client.ServiceURL("cloudservers", serverID, "block_device"),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res Attachments
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Attachments struct {
	// Specifies the disks attached to an ECS.
	VolumeAttachments []Volume `json:"volumeAttachments"`
	// Specifies the number of disks that can be attached to an ECS.
	AttachableQuantity *AttachableQuantity `json:"attachableQuantity"`
}

type Volume struct {
	// Specifies the ECS ID in UUID format.
	ServerID string `json:"serverId"`
	// Specifies the EVS disk ID in UUID format.
	VolumeID string `json:"volumeId"`
	// Specifies the mount ID, which is the same as the EVS disk ID.
	// The value is in UUID format.
	ID string `json:"id"`
	// Specifies the EVS disk size in GB.
	Size int `json:"size"`
	// Specifies the drive letter of the EVS disk, displayed as the
	// device name on the console, for example /dev/vda or /dev/vdb.
	Device string `json:"device"`
	// Specifies the PCI address.
	PCIAddress string `json:"pciAddress"`
	// Specifies the EVS disk boot sequence.
	// 0 indicates the system disk; a non-zero value indicates a data disk.
	BootIndex int `json:"bootIndex"`
	// Specifies the disk bus type.
	// Options: virtio and scsi.
	Bus string `json:"bus"`
}

// AttachableQuantity describes how many additional disks can be attached.
type AttachableQuantity struct {
	// Specifies the number of SCSI disks that can be attached to an ECS.
	FreeSCSI int `json:"free_scsi"`
	// Specifies the number of virtio_blk disks that can be attached to an ECS.
	FreeBLK int `json:"free_blk"`
	// Specifies the total number of disks that can be attached to an ECS.
	FreeDisk int `json:"free_disk"`
}
