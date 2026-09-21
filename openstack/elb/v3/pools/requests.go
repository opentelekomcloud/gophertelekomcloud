package pools

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular pool based on its unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}
