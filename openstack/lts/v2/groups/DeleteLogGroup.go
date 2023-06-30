package groups

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteLogGroup a log group by id
func DeleteLogGroup(client *golangsdk.ServiceClient, groupId string) (err error) {
	// DELETE /v2/{project_id}/groups/{log_group_id}
	_, err = client.Delete(client.ServiceURL("groups", groupId), &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	return
}
