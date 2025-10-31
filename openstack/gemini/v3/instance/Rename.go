package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RenameInstanceOpts struct {
	InstanceID string `json:"-"`
	Name       string `json:"name" required:"true"`
}

func RenameInstance(client *golangsdk.ServiceClient, opts RenameInstanceOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", opts.InstanceID, "name"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
