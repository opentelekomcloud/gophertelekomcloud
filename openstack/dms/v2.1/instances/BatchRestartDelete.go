package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const actionEdnpoint = "action"

type BatchRestartDeleteOpts struct {
	// Operation to be performed on instances. The value can be restart or delete.
	Action string `json:"action" required:"true"`
	// Indicates List of instance IDs.
	Instances string `json:"instances,omitempty"`
	// Value kafka indicates all Kafka instances that fail to be created are to be deleted.
	AllFailure string `json:"all_failure,omitempty"`
}

// BatchRestartDelete This API is used to restart or delete instances in batches.
// When an instance is being restarted, message retrieval and creation requests of the client will be rejected.
// Deleting an instance will delete the data in the instance without any backup. Exercise caution when performing this operation.
// Send POST to /v2/{project_id}/instances/action
func BatchRestartDelete(client *golangsdk.ServiceClient, opts CreateOpts) (*BatchRestartDeleteResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(ResourcePath, actionEdnpoint), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	if err != nil {
		return nil, err
	}

	var res BatchRestartDeleteResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchRestartDeleteResp struct {
	// Result of instance modification.
	Results []ResultResp `json:"results"`
}
type ResultResp struct {
	// Operation result.
	//    success: The operation succeeded.
	//    failed: The operation failed.
	Result string `json:"result"`
	// Instance ID.
	Instance string `json:"instance"`
}
