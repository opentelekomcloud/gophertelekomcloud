package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete a job.
// Send request DELETE /v1.1/{project_id}/clusters/{cluster_id}/cdm/job/{job_name}
func Delete(client *golangsdk.ServiceClient, clusterId, jobName string) error {
	_, err := client.Delete(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName), nil)
	if err != nil {
		return err
	}

	return nil
}
