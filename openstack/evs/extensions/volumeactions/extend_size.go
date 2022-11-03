package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ExtendSizeOpts struct {
	// Specifies the size of the disk after capacity expansion, in GB.
	// The new disk size ranges from the original disk size to the maximum size (32768 for a data disk and 1024 for a system disk).
	NewSize int `json:"new_size" required:"true"`
}

func ExtendSize(client *golangsdk.ServiceClient, id string, opts ExtendSizeOpts) (err error) {
	b, err := build.RequestBody(opts, "os-extend")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
