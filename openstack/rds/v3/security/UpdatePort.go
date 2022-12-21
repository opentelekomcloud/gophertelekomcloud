package security

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePortOpts struct {
	InstanceId string `json:"-"`
	// Specifies port information for all DB engines.
	// The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used).
	// The PostgreSQL database port ranges from 2100 to 9500.
	// The Microsoft SQL Server database port is 1433 or ranges from 2100 to 9500 (excluding 5355 and 5985).
	// The default values is as follows:
	// The default value of MySQL is 3306.
	// The default value of PostgreSQL is 5432.
	// The default value of Microsoft SQL Server is 1433.
	Port int32 `json:"port"`
}

func UpdatePort(c *golangsdk.ServiceClient, opts UpdatePortOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/port
	raw, err := c.Put(c.ServiceURL("instances", opts.InstanceId, "port"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}

	var res WorkflowId
	err = extract.Into(raw.Body, &res)
	return &res.WorkflowId, err
}

type WorkflowId struct {
	WorkflowId string `json:"workflowId"`
}
