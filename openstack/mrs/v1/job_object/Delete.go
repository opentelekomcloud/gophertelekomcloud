package job_object

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204},
		MoreHeaders: RequestOpts.MoreHeaders}
	_, r.Err = c.Delete(c.ServiceURL("job-executions", id), reqOpt)
	return
}
