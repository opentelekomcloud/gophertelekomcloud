package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

// DeleteConsumerGroup is used to delete a consumer group from a Kafka instance.
// Send DELETE /v2/{project_id}/instances/{instance_id}/groups/{group}
func DeleteConsumerGroup(client *golangsdk.ServiceClient, instanceId, groupId string) error {
	_, err := client.Delete(client.ServiceURL(instances.ResourcePath, instanceId, groupsPath, groupId), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
