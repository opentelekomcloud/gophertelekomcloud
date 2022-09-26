package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func DeleteGaussMySqlReadonlyNode(client *golangsdk.ServiceClient, instanceId string, nodeId string) (string, error) {
	// DELETE https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/nodes/{node_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId, "nodes", nodeId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		JobId string `json:"job_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
