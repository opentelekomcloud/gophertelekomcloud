package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := c.Delete(c.ServiceURL("backups", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}
