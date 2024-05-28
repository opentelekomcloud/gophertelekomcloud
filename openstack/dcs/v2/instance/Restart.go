package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ChangeInstanceStatusOpts struct {
	Instances []string `json:"instances,omitempty"`
	Action    string   `json:"action,omitempty"`
}

func Restart(client *golangsdk.ServiceClient, opts ChangeInstanceStatusOpts) ([]BatchOpsResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", "status"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []BatchOpsResult
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type BatchOpsResult struct {
	Result   string `json:"result"`
	Instance string `json:"instance"`
}
