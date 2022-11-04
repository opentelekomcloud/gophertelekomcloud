package volumeactions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// BootableOpts contains options for setting bootable status to a volume.
type BootableOpts struct {
	// Enables or disables the bootable attribute. You can boot an instance from a bootable volume.
	Bootable bool `json:"bootable"`
}

// SetBootable will set bootable status on a volume based on the values in BootableOpts
func SetBootable(client *golangsdk.ServiceClient, id string, opts BootableOpts) (err error) {
	b, err := build.RequestBody(opts, "os-set_bootable")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
