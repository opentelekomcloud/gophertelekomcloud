package shares

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will delete an existing SFS Turbo file system with the given UUID.
func Delete(client *golangsdk.ServiceClient, shareId string) error {
	// DELETE /v1/{project_id}/sfs-turbo/shares/{share_id}
	_, err := client.Delete(client.ServiceURL("sfs-turbo", "shares", shareId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return err
}
