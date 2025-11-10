package disk

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// ID of the ECS to which the disk will be attached.
	ServerID string `json:"-"`
	// Specifies the ECS attachment information.
	VolumeAttachment *VolumeAttachment `json:"volumeAttachment" required:"true"`
	// Specifies whether to check the request without actually attaching the disk.
	// If set to true, only a pre-check is performed and no disk is attached.
	// If set to false or omitted, the disk is attached after the check passes.
	DryRun *bool `json:"dry_run,omitempty"`
}

type VolumeAttachment struct {
	// ID of the disk to be attached. The value is in UUID format.
	VolumeID string `json:"volumeId" required:"true"`
	// Disk device name, such as /dev/sda or /dev/vdb.
	Device string `json:"device,omitempty"`
	// Disk type, for example SSD.
	VolumeType string `json:"volume_type,omitempty"`
	// Number of disks to attach.
	Count int `json:"count,omitempty"`
	// Indicates whether to attach the disk in passthrough (SCSI) mode.
	// If set to "true", the disk device type is SCSI.
	// If set to "false", the disk device type is VBD.
	HwPassthrough string `json:"hw:passthrough,omitempty"`
}

func Attach(client *golangsdk.ServiceClient, opts CreateOpts) (*JobResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/cloudservers/{server_id}/attachvolume
	raw, err := client.Post(client.ServiceURL("cloudservers", opts.ServerID, "attachvolume"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res JobResponse

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type JobResponse struct {
	// ID of the task.
	JobID string `json:"job_id"`
}
