package proxy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnlargeOpts struct {
	InstanceID string `json:"-"`
	NodeNum    int    `json:"node_num" required:"true"`
	ProxyId    string `json:"proxy_id,omitempty"`
}

func Enlarge(client *golangsdk.ServiceClient, opts EnlargeOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceID, "proxy", "enlarge"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}
