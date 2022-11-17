package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200},
		MoreHeaders: RequestOpts.MoreHeaders}
	_, r.Err = c.Get(c.ServiceURL("job-exes", id), &r.Body, reqOpt)
	return
}
