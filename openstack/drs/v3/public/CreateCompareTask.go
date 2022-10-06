package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateCompareTaskOpts struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Object-level comparison type. If the value is empty, the object-level comparison is not created.
	// If both object_level_compare_type and data_level_compare_info are left empty,
	// only the created comparison task list is queried. Value: l objects
	ObjectLevelCompareType string `json:"object_level_compare_type,omitempty"`
	// Data-level comparison information. If no data-level comparison is created, this parameter does not need to be transferred.
	// If both object_level_compare_type and data_level_compare_info are left empty, only the created comparison task list is queried.
	DataLevelCompareInfo CreateDataLevelCompareReq `json:"data_level_compare_info,omitempty"`
}

type CreateDataLevelCompareReq struct {
	// Only one unfinished data-level comparison task can exist.
	// This field determines how to process unfinished data-level comparison tasks.
	// cancel: Cancel the task and create a new one.
	// keep: Do not create a new task if the previous one is not complete.
	// Values: cancel keep
	ConflictPolicy string `json:"conflict_policy"`
	// Data-level comparison type. lines: indicates row comparison. contents: value comparison.
	// Values: lines contents
	CompareType string `json:"compare_type"`
	// Data-level comparison mode. If the value is empty, the object information needs to be transferred in
	// compare_object_infos or compare_object_infos_with_token. quick_comparison: indicates quick comparison,
	// which can be used only by whitelisted users. Default value: quick_comparison
	// Value: quick_comparison
	CompareMode string `json:"compare_mode,omitempty"`
	// Start time of a comparison task. If the value is empty, the task is started immediately.
	StartTime string `json:"start_time,omitempty"`
	// Data-level comparison object. In non-quick comparison mode, either compare_object_infos or
	// compare_object_infos_with_token must be specified based on the migration scenario.
	CompareObjectInfos []CompareObjectInfo `json:"compare_object_infos,omitempty"`
	// Object for data-level comparison (Cassandra DR only, with token information). In non-quick comparison mode, either
	// compare_object_infos or compare_object_infos_with_token must be specified based on the migration scenario.
	CompareObjectInfosWithToken []CompareObjectInfoWithToken `json:"compare_object_infos_with_token,omitempty"`
}

type CompareObjectInfo struct {
	// Database name.
	DbName string `json:"db_name"`
	// List of table names in the database.
	TableName []string `json:"table_name,omitempty"`
}

type CompareObjectInfoWithToken struct {
	// Database name.
	DbName string `json:"db_name"`
	// List of tables (with tokens) in the database.
	TableNameWithToken []CompareTableInfoWithToken `json:"table_name_with_token,omitempty"`
}

type CompareTableInfoWithToken struct {
	// Table name
	TableName string `json:"table_name"`
	// Min token of a table.
	MinToken string `json:"min_token,omitempty"`
	// Max token of a table.
	MaxToken string `json:"max_token,omitempty"`
}

func CreateCompareTask(client *golangsdk.ServiceClient, opts CreateCompareTaskOpts) (*CreateCompareTaskResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/create-compare-task
	raw, err := client.Post(client.ServiceURL("jobs", "create-compare-task"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateCompareTaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateCompareTaskResponse struct {
	// Task ID
	JobId                          string                  `json:"job_id,omitempty"`
	ObjectLevelCompareCreateResult CreateCompareTaskResult `json:"object_level_compare_create_result,omitempty"`
	DataLevelCompareCreateResult   CreateCompareTaskResult `json:"data_level_compare_create_result,omitempty"`
	ErrorCode                      string                  `json:"error_code,omitempty"`
	ErrorMsg                       string                  `json:"error_msg,omitempty"`
}

type CreateCompareTaskResult struct {
	// After the comparison task is created, the ID of the comparison task is returned for querying the result of the comparison task.
	CompareTaskId string `json:"compare_task_id,omitempty"`
	ErrorCode     string `json:"error_code,omitempty"`
	ErrorMsg      string `json:"error_msg,omitempty"`
}
