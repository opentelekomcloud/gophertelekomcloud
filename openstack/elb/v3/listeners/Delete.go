package listeners

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular Listeners based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/elb/listeners/{listener_id}
	_, err = client.Delete(client.ServiceURL("listeners", id), nil)
	return
}
