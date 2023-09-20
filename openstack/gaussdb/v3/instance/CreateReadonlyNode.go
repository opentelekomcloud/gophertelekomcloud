package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateNodeOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string
	// Read replica failover priority ranging from 1 to 16.
	// The total number of primary node and read replicas is less than or equal to 16.
	Priorities []int `json:"priorities"`
}

func CreateReplica(client *golangsdk.ServiceClient, opts CreateNodeOpts) (*CreateNodeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/nodes/enlarge
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "nodes", "enlarge"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateNodeResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateNodeResponse struct {
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
