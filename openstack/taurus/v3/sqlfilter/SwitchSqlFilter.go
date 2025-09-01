package sqlfilter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type SqlFilterSwitchOpts struct {
	InstanceId   string `json:"-"`
	SwitchStatus string `json:"switch_status" required:"true"`
}

func SwitchSqlFilter(client *golangsdk.ServiceClient, opts SqlFilterSwitchOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "sql-filter", "switch"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobId, extract.Into(raw.Body, &res)
}

type jobResponse struct {
	JobId string `json:"job_id"`
}
