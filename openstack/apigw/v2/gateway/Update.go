package gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	ID               string `json:"-"`
	Description      string `json:"description,omitempty"`
	MaintainBegin    string `json:"maintain_begin,omitempty"`
	MaintainEnd      string `json:"maintain_end,omitempty"`
	InstanceName     string `json:"instance_name,omitempty"`
	SecGroupID       string `json:"security_group_id,omitempty"`
	VpcepServiceName string `json:"vpcep_service_name,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Gateway, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw/instances", opts.ID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Gateway

	err = extract.Into(raw.Body, &res)
	return &res, err
}
