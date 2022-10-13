package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type ExtendSizeOpts struct {
	// NewSize is the new size of the volume, in GB.
	NewSize int `json:"new_size" required:"true"`
}

func ExtendSize(client *golangsdk.ServiceClient, id string, opts ExtendSizeOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "os-extend")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
