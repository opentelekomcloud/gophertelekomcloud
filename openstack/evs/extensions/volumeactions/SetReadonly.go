package volumeactions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ReadonlyOpts struct {
	Readonly bool `json:"readonly"`
}

func SetReadonly(client *golangsdk.ServiceClient, id string, opts ReadonlyOpts) (err error) {
	b, err := build.RequestBody(opts, "os-update_readonly_flag")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
