package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ScaleOutOpts struct {
	// Flavor ID of the VM for deploying the DDM instance that is to be scaled out (Mandatory)
	FlavorId string `json:"flavor_id" required:"true"`
	// Number of nodes to be added (Mandatory)
	NodeNumber int `json:"node_number" required:"true"`
	// Group ID, which specifies the node group that is scaled out. This parameter must be specified if there is more than one node group (Optional)
	GroupId string `json:"group_id,omitempty"`
	// The function has not been supported, and this field is reserved (Optional)
	IsAutoPay bool `json:"is_auto_pay,omitempty"`
}

func ScaleOut(client *golangsdk.ServiceClient, opts ScaleOutOpts) (*ScaleOutResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/instances/{instance_id}/action/enlarge
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ScaleOutResponse
	return &res, extract.Into(raw.Body, &res)
}

type ScaleOutResponse struct {
	// DDM instance ID
	InstanceId string `json:"instanceId"`
	// DDM instance name
	InstanceName string `json:"instanceName"`
	// Task ID
	JobId string `json:"jobId"`
}
