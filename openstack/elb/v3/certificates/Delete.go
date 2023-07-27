package certificates

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular Certificate based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/elb/certificates/{certificate_id}
	_, err = client.Delete(client.ServiceURL("certificates", id), nil)
	return
}
