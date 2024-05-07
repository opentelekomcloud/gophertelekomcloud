package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const startEndpoint = "start"

// Start is used to start a job.
// Send request PUT /v1.1/{project_id}/clusters/{cluster_id}/cdm/job/{job_name}/start
func Start(client *golangsdk.ServiceClient, clusterId, jobName string) (*JobResp, error) {
	raw, err := client.Put(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName, startEndpoint), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp JobResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err
}
