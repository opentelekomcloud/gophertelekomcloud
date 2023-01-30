package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	// DELETE https://{Endpoint}/v3/{project_id}/instances/{instance_id}
	raw, err := client.Delete(client.ServiceURL("instances", instanceId), nil)
	return extraJob(err, raw)
}
