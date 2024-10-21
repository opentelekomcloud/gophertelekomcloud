package hosted_connect

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("dcaas", "hosted-connects", id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}
