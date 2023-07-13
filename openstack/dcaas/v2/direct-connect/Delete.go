package direct_connect

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular DirectConnect based on its
// unique ID.

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("dcaas", "direct-connects", id), nil)
	return
}
