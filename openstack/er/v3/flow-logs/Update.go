package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Enterprise router ID
	RouterID string `json:"-"`
	// Flow log name
	Name string `json:"name,omitempty"`
	// Supplementary information about flow log
	Description string `json:"description,omitempty"`
}

func Update(client *golangsdk.ServiceClient, flowLogID string, opts UpdateOpts) (*FlowLogResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("enterprise-router", opts.RouterID, "flow-logs", flowLogID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FlowLogResponse
	return &res, extract.IntoStructPtr(raw.Body, &res, "flow_log")
}
