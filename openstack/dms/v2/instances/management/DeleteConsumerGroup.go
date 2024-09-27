package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DeleteConsumerGroup is used to delete a consumer group from a Kafka instance.
// Send DELETE /v2/{project_id}/instances/{instance_id}/groups/{group}
func DeleteConsumerGroup(client *golangsdk.ServiceClient, instanceId, groupId string) error {
	_, err := client.Delete(client.ServiceURL("instances", instanceId, "groups", groupId), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
