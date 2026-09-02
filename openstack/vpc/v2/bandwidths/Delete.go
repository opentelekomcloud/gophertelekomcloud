package bandwidths

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete removes a shared bandwidth. Only a shared bandwidth can be deleted.
func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL(client.ProjectID, "bandwidths", id), nil)
	return err
}
