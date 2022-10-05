package sync

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

// POST /v3/{project_id}/jobs/batch-sync-policy

// IdJobResp
