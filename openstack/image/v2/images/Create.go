package images

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateOpts represents options used to create an image.
type CreateOpts struct {
	// Name is the name of the new image.
	Name string `json:"name" required:"true"`
	// Id is the the image ID.
	ID string `json:"id,omitempty"`
	// Visibility defines who can see/use the image.
	Visibility *ImageVisibility `json:"visibility,omitempty"`
	// Tags is a set of image tags.
	Tags []string `json:"tags,omitempty"`
	// ContainerFormat is the format of the
	// container. Valid values are ami, ari, aki, bare, and ovf.
	ContainerFormat string `json:"container_format,omitempty"`
	// DiskFormat is the format of the disk. If set,
	// valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi,
	// and iso.
	DiskFormat string `json:"disk_format,omitempty"`
	// MinDisk is the amount of disk space in
	// GB that is required to boot the image.
	MinDisk int `json:"min_disk,omitempty"`
	// MinRAM is the amount of RAM in MB that
	// is required to boot the image.
	MinRAM int `json:"min_ram,omitempty"`
	// protected is whether the image is not deletable.
	Protected *bool `json:"protected,omitempty"`
	// properties is a set of properties, if any, that
	// are associated with the image.
	Properties map[string]string `json:"-"`
}

// ToImageCreateMap assembles a request body based on the contents of
// a CreateOpts.
func (opts CreateOpts) ToImageCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.Properties != nil {
		for k, v := range opts.Properties {
			b[k] = v
		}
	}
	return b, nil
}

// Create implements create image request.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToImageCreateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Post(client.ServiceURL("images"), b, &r.Body, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}

// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToImageCreateMap() (map[string]interface{}, error)
}
