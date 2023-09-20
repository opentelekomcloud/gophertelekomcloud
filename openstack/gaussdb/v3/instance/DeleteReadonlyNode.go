package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteReplica(client *golangsdk.ServiceClient, instanceId string, nodeId string) (*string, error) {
	// DELETE https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/nodes/{node_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId, "nodes", nodeId), &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return extraJob(err, raw)
}
