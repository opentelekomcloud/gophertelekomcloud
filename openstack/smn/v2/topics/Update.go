package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOpts is a struct that contains all the parameters.
type UpdateOpts struct {
	Id string `json:"-"`
	// Topic display name
	DisplayName string `json:"display_name,omitempty"`
}

// Update a topic display name with given parameters.
func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*UpdateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("topics", opts.Id), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateResp struct {
	RequestId string `json:"request_id"`
}
