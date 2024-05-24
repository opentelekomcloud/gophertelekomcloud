package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const stopEndpoint = "stop"

// Stop is used to stop a job.
// Send request PUT /v1.1/{project_id}/clusters/{cluster_id}/cdm/job/{job_name}/stop
func Stop(client *golangsdk.ServiceClient, clusterId, jobName string) (*UpdateResp, error) {
	raw, err := client.Put(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName, stopEndpoint), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp UpdateResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err
}
