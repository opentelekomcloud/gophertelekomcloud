package virtual_interface

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular VirtualInterface based on its
// unique ID.

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("virtual-virtual-interfaces", id), nil)
	return
}
