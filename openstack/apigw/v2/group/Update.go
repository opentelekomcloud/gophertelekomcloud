package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	GatewayID   string `json:"-"`
	GroupID     string `json:"-"`
	Name        string `json:"name" required:"true"`
	Description string `json:"remark,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*GroupResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res GroupResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
