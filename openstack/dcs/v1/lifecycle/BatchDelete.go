package lifecycle

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteOpts struct {
	// An indicator of whether all DCS instances failed to be created will be deleted. Options:
	// Options:
	// true: all instances that fail to be created are deleted. In this case, the instances parameter in the request can be empty.
	// false or other values: The DCS instances specified by the instances parameter in the API request will be deleted.
	AllFailure *bool `q:"allFailure"`
	Body       BatchDeleteBody
}

type BatchDeleteBody struct {
	// IDs of DCS instances to be deleted.
	// This parameter is set only when the allFailure parameter in the URI is set to false or another value.
	// A maximum of 50 instances can be deleted at a time.
	Instances []string `json:"instances,omitempty"`
}

func BatchDelete(client *golangsdk.ServiceClient, opts BatchDeleteOpts) ([]BatchOpsResult, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts.Body, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.DeleteWithBody(client.ServiceURL("instances")+q.String(), b, nil)
	if err != nil {
		return nil, err
	}

	var res []BatchOpsResult
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type BatchOpsResult struct {
	// Instance deletion result. Options: success and failed
	Result string `json:"result,omitempty"`
	// DCS instance ID.
	Instance string `json:"instance,omitempty"`
}
