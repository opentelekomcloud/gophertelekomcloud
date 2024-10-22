package virtual_interface

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeletePeer(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("dcaas", "vif-peers", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}
