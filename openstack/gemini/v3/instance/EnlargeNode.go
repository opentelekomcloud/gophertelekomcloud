package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnlargeNodeOpts struct {
	InstanceId string `json:"-"`
	Num        int    `json:"num" required:"true"`
	SubnetId   string `json:"subnet_id,omitempty"`
}

func EnlargeNode(client *golangsdk.ServiceClient, opts EnlargeNodeOpts) (*EnlargeNodeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "enlarge-node"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res EnlargeNodeResponse
	return &res, extract.Into(raw.Body, &res)
}

type EnlargeNodeResponse struct {
	JobId string `json:"job_id"`
}
