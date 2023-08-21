package dc_endpoint_group

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete is used to delete a Direct Connect endpoint group
func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("dcaas", "dc-endpoint-groups", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}
