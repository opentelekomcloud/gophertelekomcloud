package loggroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a log group by id
func Delete(client *golangsdk.ServiceClient, groupId string) (err error) {
	// DELETE /v2.0/{project_id}/log-groups/{group_id}
	_, err = client.Delete(client.ServiceURL("log-groups", groupId), nil)
	return
}
