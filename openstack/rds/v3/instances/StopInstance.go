package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func StopInstance(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action/shutdown
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action", "shutdown"), struct{}{}, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraJob(err, raw)
}
