package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type EnableOpts struct {
	// DMS instance id
	InstanceId string `json:"-" required:"true"`
	// Bandwidth for deploying Smart Connect, that is, the maximum amount
	// of data transferred per unit time. Use the bandwidth of the current instance.
	Specification string `json:"specification,omitempty"`
	// Number of connectors. Min.: 2.
	// The default value is 2 if it is not specified.
	NodeCount string `json:"node_cnt,omitempty"`
	// Specification code of the connector.
	// This parameter is mandatory only for old instance flavors.
	SpecCode string `json:"spec_code,omitempty"`
}

// This API is used to enable Smart Connect so you can create a connector.
// Send POST /v2/{project_id}/instances/{instance_id}/connector
func Enable(client *golangsdk.ServiceClient, opts EnableOpts) (*EnableSmartResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "connector"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res EnableSmartResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type EnableSmartResp struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Instance dump ID.
	ConnectorId string `json:"connector_id"`
}
