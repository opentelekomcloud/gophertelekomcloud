package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// New name of the dedicated WAF engine
	Name string `json:"instancename" required:"true"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/premium-waf/instance
	raw, err := client.Put(client.ServiceURL("instance", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Instance
	return &res, extract.Into(raw.Body, &res)
}
