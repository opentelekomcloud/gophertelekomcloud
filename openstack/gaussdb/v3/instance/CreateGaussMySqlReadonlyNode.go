package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type MysqlCreateReadonlyNodeOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string
	// Read replica failover priority ranging from 1 to 16.
	// The total number of primary node and read replicas is less than or equal to 16.
	Priorities []int32 `json:"priorities"`
	// Whether the order will be automatically paid after yearly/monthly instances are created.
	// This parameter does not affect the payment method of automatic renewal.
	// 	true: The order will be automatically paid from your account. The default value is true.
	// 	false: The order will be manually paid.
	IsAutoPay string `json:"is_auto_pay,omitempty"`
}

func CreateGaussMySqlReadonlyNode(client *golangsdk.ServiceClient, opts MysqlCreateReadonlyNodeOpts) (*CreateGaussMySqlReadonlyNodeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/nodes/enlarge
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateGaussMySqlReadonlyNodeResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateGaussMySqlReadonlyNodeResponse struct {
	// Instance ID
	InstanceId string `json:"instance_id"`
	// Node name list
	NodeNames []string `json:"node_names"`
	// Instance creation task ID
	// This parameter is returned only for the creation of pay-per-use instances.
	JobId string `json:"job_id"`
	// Order ID. This parameter is returned only for the creation of yearly/monthly instances.
	OrderId string `json:"order_id"`
}
