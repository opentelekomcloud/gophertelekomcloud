package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type AttachMode string

const (
	ReadOnly  AttachMode = "ro"
	ReadWrite AttachMode = "rw"
)

type AttachOpts struct {
	// The mountpoint of this volume.
	MountPoint string `json:"mountpoint,omitempty"`
	// The nova instance ID, can't set simultaneously with HostName.
	InstanceUUID string `json:"instance_uuid,omitempty"`
	// The hostname of baremetal host, can't set simultaneously with InstanceUUID.
	HostName string `json:"host_name,omitempty"`
	// Mount mode of this volume.
	Mode AttachMode `json:"mode,omitempty"`
}

func Attach(client *golangsdk.ServiceClient, id string, opts AttachOpts) (err error) {
	b, err := build.RequestBody(opts, "os-attach")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
