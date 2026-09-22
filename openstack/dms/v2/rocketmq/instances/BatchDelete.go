package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteOpts struct {
	// Operation to be performed on instances. Value: delete.
	Action string `json:"action" required:"true"`
	// List of instance IDs.
	Instances []string `json:"instances,omitempty"`
	// When set to reliability, all RocketMQ instances that fail
	// to be created will be deleted.
	AllFailure string `json:"all_failure,omitempty"`
	// Whether to forcibly delete.
	//    true: forcibly deleted instances are not moved to the recycle bin.
	//    false: instances are moved to the recycle bin after the recycle bin
	//    function is enabled.
	ForceDelete *bool `json:"force_delete,omitempty"`
}

// BatchDeleteResponse is returned by BatchDelete.
type BatchDeleteResponse struct {
	// Result of instance modification.
	Results []DeleteResult `json:"results"`
}

type DeleteResult struct {
	// Operation result. Options: success, failed.
	Result string `json:"result"`
	// Instance ID.
	Instance string `json:"instance"`
}

// BatchDelete deletes RocketMQ instances in batches. Data in the instances will
// be deleted without any backup. Exercise caution when performing this operation.
// Send POST to /v2/{project_id}/instances/action
func BatchDelete(client *golangsdk.ServiceClient, opts BatchDeleteOpts) (*BatchDeleteResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	if err != nil {
		return nil, err
	}

	var res BatchDeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
