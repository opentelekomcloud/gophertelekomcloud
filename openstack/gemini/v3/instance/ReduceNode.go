package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ReduceNodeOpts struct {
	InstanceID string   `json:"-"`
	Num        *int     `json:"num,omitempty"`
	NodeList   []string `json:"node_list,omitempty"`
}

func ReduceNode(client *golangsdk.ServiceClient, opts ReduceNodeOpts) (*ReduceNodeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceID, "reduce-node"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res ReduceNodeResponse
	return &res, extract.Into(raw.Body, &res)
}

type ReduceNodeResponse struct {
	JobId string `json:"job_id"`
}
