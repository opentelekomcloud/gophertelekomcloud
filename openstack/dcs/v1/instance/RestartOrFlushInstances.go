package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeInstanceStatusOpts struct {
	// List of DCS instance IDs.
	Instances []string `json:"instances,omitempty"`
	// Action performed on DCS instances. Options: restart, and flush.
	// NOTE
	// Only DCS Redis 4.0 and 5.0 instances can be flushed.
	Action string `json:"action,omitempty"`
}

func RestartOrFlushInstances(client *golangsdk.ServiceClient, opts ChangeInstanceStatusOpts) ([]BatchOpsResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", "status"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return nil, err
	}

	var res []BatchOpsResult
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type BatchOpsResult struct {
	// Instance modification result. Options: success or failed
	Result string `json:"result,omitempty"`
	// DCS instance ID.
	Instance string `json:"instance,omitempty"`
}
