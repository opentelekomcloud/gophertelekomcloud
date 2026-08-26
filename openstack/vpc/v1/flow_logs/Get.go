package flow_logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, flowLogID string) (*FlowLog, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "fl", "flow_logs", flowLogID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res flowLogResponse
	err = extract.Into(raw.Body, &res)
	return &res.FlowLog, err
}
