package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Enable(client *golangsdk.ServiceClient, routerID, flowLogID string) (*FlowLogResponse, error) {
	raw, err := client.Post(client.ServiceURL("enterprise-router", routerID, "flow-logs", flowLogID, "enable"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res FlowLogResponse
	return &res, extract.IntoStructPtr(raw.Body, &res, "flow_log")
}
