package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func StartFailover(client *golangsdk.ServiceClient, instanceId string) (*StartFailoverResponse, error) {
	// PUT /v3/{project_id}/instances/{instance_id}/failover
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "failover"), struct{}{}, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res StartFailoverResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type StartFailoverResponse struct {
	// Indicates the DB instance ID.
	InstanceId string `json:"instanceId,omitempty"`
	// Indicates former slave node that becomes a master node.
	NodeId string `json:"nodeId,omitempty"`
	// Indicates the workflow ID.
	WorkflowId string `json:"workflowId,omitempty"`
}
