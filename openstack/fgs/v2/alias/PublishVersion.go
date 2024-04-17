package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
)

type PublishOpts struct {
	FuncUrn     string `json:"-"`
	Digest      string `json:"digest,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
}

func PublishVersion(client *golangsdk.ServiceClient, opts PublishOpts) (*function.FuncGraph, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", opts.FuncUrn, "versions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res function.FuncGraph
	return &res, extract.Into(raw.Body, &res)
}
