package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AdminState  *bool   `json:"admin_state,omitempty"`
}

func Update(client *golangsdk.ServiceClient, flowLogID string, opts UpdateOpts) (*FlowLog, error) {
	b, err := build.RequestBody(opts, "flow_log")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(client.ProjectID, "fl", "flow_logs", flowLogID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res flowLogResponse
	err = extract.Into(raw.Body, &res)
	return &res.FlowLog, err
}
