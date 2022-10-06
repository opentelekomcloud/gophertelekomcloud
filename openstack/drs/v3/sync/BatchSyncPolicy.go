package sync

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/drs/v3/public"
)

type BatchSyncPolicyOpts struct {
	Jobs []SyncPolicyReq `json:"jobs"`
}

type SyncPolicyReq struct {
	// Conflict policy.
	// Values:
	// ignore: Ignore the conflict.
	// overwrite: Overwrite the existing data with the conflicting data.
	// stop: Report an error.
	ConflictPolicy string `json:"conflict_policy"`
	// Whether to synchronize DDL during incremental synchronization.
	DdlTrans bool `json:"ddl_trans"`
	// DDL filtering policy.
	// Values:
	// drop_database
	FilterDdlPolicy string `json:"filter_ddl_policy"`
	// Whether to synchronize indexes during incremental synchronization.
	IndexTrans bool `json:"index_trans"`
	// Task ID.
	JobId string `json:"job_id"`
}

func BatchSyncPolicy(client *golangsdk.ServiceClient, opts BatchSyncPolicyOpts) (*public.IdJobResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-sync-policy
	raw, err := client.Post(client.ServiceURL("jobs", "batch-sync-policy"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res public.IdJobResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
