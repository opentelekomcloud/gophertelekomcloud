package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func StopInstance(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action/shutdown
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action", "shutdown"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
