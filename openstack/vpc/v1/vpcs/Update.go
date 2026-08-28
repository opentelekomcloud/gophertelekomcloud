package vpcs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name             string  `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	CIDR             string  `json:"cidr,omitempty"`
	Routes           []Route `json:"routes,omitempty"`
	EnableSharedSnat *bool   `json:"enable_shared_snat,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Vpc, error) {
	b, err := build.RequestBody(opts, "vpc")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(client.ProjectID, "vpcs", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		VPC Vpc `json:"vpc"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VPC, err
}
