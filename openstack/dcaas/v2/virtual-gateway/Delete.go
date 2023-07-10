package virtual_gateway

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will permanently delete a particular VirtualGateway based on its
// unique ID.

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("virtual-gateways", id), nil)
	return
}
