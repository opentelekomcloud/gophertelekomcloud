package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type UploadImageOptsBuilder interface {
	ToVolumeUploadImageMap() (map[string]interface{}, error)
}

type UploadImageOpts struct {
	// Container format, may be bare, ofv, ova, etc.
	ContainerFormat string `json:"container_format,omitempty"`

	// Disk format, may be raw, qcow2, vhd, vdi, vmdk, etc.
	DiskFormat string `json:"disk_format,omitempty"`

	// The name of image that will be stored in glance.
	ImageName string `json:"image_name,omitempty"`

	// Force image creation, usable if volume attached to instance.
	Force bool `json:"force,omitempty"`
}

func (opts UploadImageOpts) ToVolumeUploadImageMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-volume_upload_image")
}

func UploadImage(client *golangsdk.ServiceClient, id string, opts UploadImageOptsBuilder) (r UploadImageResult) {
	b, err := opts.ToVolumeUploadImageMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
