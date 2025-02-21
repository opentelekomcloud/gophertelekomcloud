package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to obtain the status of a CFW task.
func Get(client *golangsdk.ServiceClient, jobId string) (*GetCreateFirewallJobResponseData, error) {
	// GET /v3/{project_id}/jobs/{job_id}
	raw, err := client.Get(client.ServiceURL("jobs", jobId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res JobResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err
}

type JobResponse struct {
	Data GetCreateFirewallJobResponseData `json:"data"`
}

type GetCreateFirewallJobResponseData struct {
	// ID of the task for creating a pay-per-use firewall.
	Id string `json:"id"`
	// Task execution status, which indicates whether a firewall is successfully created.
	// Enumerated Values: Running, Success, Failed
	Status string `json:"status"`
	// Creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	BeginTime string `json:"begin_time"`
	// End time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndTime string `json:"end_time"`
}
