package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	Description string `json:"description" required:"true"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("clustergroups", id, "description"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
