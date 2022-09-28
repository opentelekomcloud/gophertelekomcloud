package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Delete(client *golangsdk.ServiceClient, instanceId string) (string, error) {
	// DELETE https://{Endpoint}/v3/{project_id}/instances/{instance_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId), nil)
	if err != nil {
		return "", err
	}

	var res struct {
		JobId string `json:"job_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.JobId, err
}
