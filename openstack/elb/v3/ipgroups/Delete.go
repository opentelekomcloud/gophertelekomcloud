package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete will permanently delete a particular IpGroup based on its
// unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("ipgroups", id), nil)
	return
}
