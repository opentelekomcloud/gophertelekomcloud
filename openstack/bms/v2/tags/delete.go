package tags

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular tag based on its unique ID.
func Delete(c *golangsdk.ServiceClient, serverId string) (err error) {
	_, err = c.Delete(c.ServiceURL("servers", serverId, "tags"), nil)
	return
}
