package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type AttachOptsBuilder interface {
	ToVolumeAttachMap() (map[string]interface{}, error)
}

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

func (opts AttachOpts) ToVolumeAttachMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-attach")
}

func Attach(client *golangsdk.ServiceClient, id string, opts AttachOptsBuilder) (r AttachResult) {
	b, err := opts.ToVolumeAttachMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
