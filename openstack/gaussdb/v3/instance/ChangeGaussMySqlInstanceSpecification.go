package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type MysqlChangeSpecificationOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string
	// Specification change information.
	ResizeFlavor MysqlResizeFlavor `json:"resize_flavor"`
	// Whether the order will be automatically paid after yearly/monthly instances are changed.
	// true: The order will be automatically paid from your account. The default value is true.
	// false: The order will be manually paid.
	IsAutoPay string `json:"is_auto_pay,omitempty"`
}

type MysqlResizeFlavor struct {
	// Specification code
	SpecCode string `json:"spec_code"`
}

func ChangeGaussMySqlInstanceSpecification(client *golangsdk.ServiceClient, opts MysqlChangeSpecificationOpts) (*ChangeGaussMySqlInstanceSpecificationResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ChangeGaussMySqlInstanceSpecificationResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ChangeGaussMySqlInstanceSpecificationResponse struct {
	// Job ID for changing instance specifications.
	// This parameter is returned only when you change the specifications of a pay-per-use instance.
	JobId string `json:"job_id"`
	// Order ID for changing instance specifications.
	// This parameter is returned only when you change the specification of a yearly/monthly instance.
	OrderId string `json:"order_id"`
}
