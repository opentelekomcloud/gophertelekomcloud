package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type ExtendSizeOptsBuilder interface {
	ToVolumeExtendSizeMap() (map[string]interface{}, error)
}

type ExtendSizeOpts struct {
	// NewSize is the new size of the volume, in GB.
	NewSize int `json:"new_size" required:"true"`
}

func (opts ExtendSizeOpts) ToVolumeExtendSizeMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-extend")
}

func ExtendSize(client *golangsdk.ServiceClient, id string, opts ExtendSizeOptsBuilder) (r ExtendSizeResult) {
	b, err := opts.ToVolumeExtendSizeMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
