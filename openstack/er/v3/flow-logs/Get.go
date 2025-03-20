package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, routerID, flowLogID string) (*FlowLogResponse, error) {
	raw, err := client.Get(client.ServiceURL("enterprise-router", routerID, "flow-logs", flowLogID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FlowLogResponse
	err = extract.IntoStructPtr(raw.Body, &res, "flow_log")
	return &res, err
}
