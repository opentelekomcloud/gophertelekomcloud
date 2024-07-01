package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, groupId string) (err error) {
	// DELETE /v5/{project_id}/host-management/groups
	_, err = client.Delete(client.ServiceURL("host-management", "groups", groupId), nil)
	return
}
