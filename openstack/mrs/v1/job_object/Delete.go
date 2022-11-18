package job_object

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(c *golangsdk.ServiceClient, id string) (err error) {
	_, err = c.Delete(c.ServiceURL("job-executions", id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
		MoreHeaders: golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
		}.MoreHeaders,
	})
	return
}
