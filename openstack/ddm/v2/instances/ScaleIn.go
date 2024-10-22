package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ScaleInOpts struct {
	// Number of nodes to be removed. The maximum value is the instance nodes minus 1. (Mandatory)
	NodeNumber int `json:"node_number" required:"true"`
	// Group ID, which specifies the node group that is scaled out. This parameter must be specified if there is more than one node group (Optional)
	GroupId string `json:"group_id,omitempty"`
}

// ScaleIn is used to remove nodes from a specified DDM instance.
func ScaleIn(client *golangsdk.ServiceClient, instanceId string, opts ScaleInOpts) (*ScaleInResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/instances/{instance_id}/action/reduce
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action", "reduce"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ScaleInResponse
	return &res, extract.Into(raw.Body, &res)
}

type ScaleInResponse struct {
	// DDM instance ID
	InstanceId string `json:"instanceId"`
	// DDM instance name
	InstanceName string `json:"instanceName"`
	// Task ID
	JobId string `json:"jobId"`
}
