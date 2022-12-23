package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResizeOpts struct {
	InstanceId string `json:"-"`
	// Specifies the resource specification code. Use rds.mysql.m1.xlarge as an example. rds indicates RDS, mysql indicates the DB engine, and m1.xlarge indicates the performance specification (large-memory). The parameter containing rr indicates the read replica specifications. The parameter not containing rr indicates the single or primary/standby DB instance specifications.
	SpecCode string `json:"spec_code" required:"true"`
}

func Resize(client *golangsdk.ServiceClient, opts ResizeOpts) (*string, error) {
	b, err := build.RequestBody(opts, "resize_flavor")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
