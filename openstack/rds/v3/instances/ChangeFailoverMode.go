package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeFailoverModeOpts struct {
	InstanceId string `json:"-"`
	// Specifies the synchronize model.
	// For MySQL, the value can be any of the following:
	// async: asynchronous
	// semisync: semi-synchronous
	// For PostgreSQL, the value can be any of the following:
	// async: asynchronous
	// sync: synchronous
	Mode string `json:"mode" required:"true"`
}

func ChangeFailoverMode(client *golangsdk.ServiceClient, opts ChangeFailoverModeOpts) (*ChangeFailoverModeResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/failover/mode
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "failover", "mode"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ChangeFailoverModeResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ChangeFailoverModeResponse struct {
	// Indicates the DB instance ID.
	InstanceId string `json:"instanceId,omitempty"`
	// Indicates the replication mode.
	ReplicationMode string `json:"replicationMode,omitempty"`
	// Indicates the workflow ID.
	WorkflowId string `json:"workflowId,omitempty"`
}
