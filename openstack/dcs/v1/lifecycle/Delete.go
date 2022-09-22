package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete an instance by id
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("instances", id), nil)
	return
}
