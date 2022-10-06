package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchSpecialTestConnectionOpts struct {
	Jobs []BatchJobActionReq `json:"jobs"`
}

type BatchJobActionReq struct {
	// Specific operation to be performed.
	Action string `json:"action"`
	// Task ID. (In cluster mode, the value is the ID of the parent task.).
	JobId string `json:"job_id"`
	// The parameter that corresponds to the operation.
	// Please refer to Docs API for details.
	Property string `json:"property"`
}

func BatchValidateClustersConnections(client *golangsdk.ServiceClient, opts BatchSpecialTestConnectionOpts) (*BatchValidateClustersConnectionsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/cluster/batch-connection
	raw, err := client.Post(client.ServiceURL("jobs", "cluster", "batch-connection"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchValidateClustersConnectionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchValidateClustersConnectionsResponse struct {
	Results []CheckJobResp `json:"results,omitempty"`
	Count   int32          `json:"count,omitempty"`
}
