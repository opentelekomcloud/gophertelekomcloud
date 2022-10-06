package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type QueryCompareResultOpts struct {
	// Task ID
	JobId string `json:"job_id"`
	// ID of the object-level comparison task that requests the query result.
	ObjectLevelCompareId string `json:"object_level_compare_id,omitempty"`
	// ID of the row comparison task that requests the query result.
	LineCompareId string `json:"line_compare_id,omitempty"`
	// ID of the value comparison task that requests the query result.
	ContentCompareId string `json:"content_compare_id,omitempty"`
	// Current page number for pagination query, which is valid for the query result of comparison tasks.
	CurrentPage int32 `json:"current_page"`
	// Number of records on each page. This parameter is valid only for query results of comparison tasks.
	PerPage int32 `json:"per_page"`
}

func ListCompareResult(client *golangsdk.ServiceClient, opts QueryCompareResultOpts) (*ListCompareResultResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/query-compare-result
	raw, err := client.Post(client.ServiceURL("jobs", "query-compare-result"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListCompareResultResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListCompareResultResponse struct {
	// Task ID
	JobId string `json:"job_id,omitempty"`
	// Object-level comparison result.
	ObjectLevelCompareResults ObjectCompareResult `json:"object_level_compare_results,omitempty"`
	// Row comparison result.
	LineCompareResults LineCompareResult `json:"line_compare_results,omitempty"`
	// Value comparison result.
	ContentCompareResults ContentCompareResult `json:"content_compare_results,omitempty"`
	// List of comparison tasks
	CompareTaskListResults CompareTaskListResult `json:"compare_task_list_results,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type ObjectCompareResult struct {
	// ID of an object-level comparison task.
	CompareTaskId string `json:"compare_task_id"`
	// Overview of object comparison results.
	ObjectCompareOverview []ObjectCompareResultOverview `json:"object_compare_overview,omitempty"`
	// Object comparison result details. The key value is the object type in the object comparison result overview.
	ObjectCompareDetails map[string][]ObjectCompareResultDetails `json:"object_compare_details,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type ObjectCompareResultOverview struct {
	// Object type.
	// Values:
	// DB-database
	// TABLE-table
	// VIEW-view
	// EVENT-event
	// ROUTINE - stored procedure and function
	// INDEX: index
	// TRIGGER: trigger
	// SYNONYM - synonym
	// FUNCTION-function
	// PROCEDURE: stored procedure
	// TYPE: user-defined type
	// RULE-rule
	// DEFAULT_TYPE: default value
	// PLAN_GUIDE-execution plan
	// CONSTRAINT-constraint
	// FILE_GROUP-file group
	// PARTITION_FUNCTION-partition function
	// PARTITION_SCHEME-partition scheme
	// TABLE_COLLATION-table sorting rule
	ObjectType string `json:"object_type"`
	// Comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	ObjectCompareResult string `json:"object_compare_result"`
	// Number of objects of this type in the destination database.
	TargetCount int32 `json:"target_count"`
	// Number of objects of this type in the source database.
	SourceCount int32 `json:"source_count"`
	// Number of differences between the source and destination databases.
	DiffCount int32 `json:"diff_count"`
}

type ObjectCompareResultDetails struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Destination database name.
	TargetDbName string `json:"target_db_name"`
	// Value in the source database.
	SourceDbValue string `json:"source_db_value,omitempty"`
	// Value in the destination database.
	TargetDbValue string `json:"target_db_value,omitempty"`
	// Error message.
	ErrorMessage string `json:"error_message,omitempty"`
}

type LineCompareResult struct {
	// ID of a row comparison task.
	CompareTaskId string `json:"compare_task_id,omitempty"`
	// Row comparison result overview.
	LineCompareOverview []LineCompareResultOverview `json:"line_compare_overview,omitempty"`
	// Row comparison result overview.
	LineCompareOverviewCount int32 `json:"line_compare_overview_count,omitempty"`
	// Row comparison result details.
	LineCompareDetails []LineCompareResultDetails `json:"line_compare_details,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type LineCompareResultOverview struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Destination database name.
	TargetDbName string `json:"target_db_name"`
	// Comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	LineCompareResult string `json:"line_compare_result"`
}

type LineCompareResultDetails struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Row comparison details of the tables in the database.
	LineCompareDetail []LineCompareDetail `json:"LineCompareDetail"`
	// Total number of row comparison results in the database.
	LineCompareDetailCount int32 `json:"line_compare_detail_count"`
}

type LineCompareDetail struct {
	// Table name of the source database.
	SourceTableName string `json:"source_table_name"`
	// Table name of the destination database.
	TargetTableName string `json:"target_table_name"`
	// Number of table rows in the source database.
	SourceRowNum int32 `json:"source_row_num"`
	// Number of table rows in the destination database.
	TargetRowNum int32 `json:"target_row_num"`
	// Difference between the tables in the source and destination databases.
	DiffRowNum int32 `json:"diff_row_num"`
	// Comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	LineCompareResult string `json:"line_compare_result"`
	// Additional information.
	Message string `json:"message,omitempty"`
}

type ContentCompareResult struct {
	// ID of a value comparison task.
	CompareTaskId string `json:"compare_task_id"`
	// Content comparison result overview.
	ContentCompareOverview []ContentCompareResultOverview `json:"content_compare_overview,omitempty"`
	// Total number of value comparison results.
	ContentCompareOverviewCount int32 `json:"content_compare_overview_count,omitempty"`
	// Value comparison result details.
	ContentCompareDetails []ContentCompareResultDetails `json:"content_compare_details,omitempty"`
	// The value comparison results are different.
	ContentCompareDiffs []ContentCompareResultDiffs `json:"content_compare_diffs,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}

type ContentCompareResultOverview struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Destination database name.
	TargetDbName string `json:"target_db_name"`
	// Comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	ContentCompareResult string `json:"content_compare_result"`
}

type ContentCompareResultDetails struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Value comparison details of the tables in the database.
	ContentCompareDetail []ContentCompareDetail `json:"content_compare_detail"`
	// Total number of value comparison results.
	ContentCompareDetailCount int32 `json:"content_compare_detail_count"`
	// Value comparison details of tables in the database (tables that cannot be compared).
	ContentUncompareDetail []ContentCompareDetail `json:"content_uncompare_detail,omitempty"`
	// Total number of value comparison results (tables that cannot be compared).
	ContentUncompareDetailCount int32 `json:"content_uncompare_detail_count"`
}

type ContentCompareDetail struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Destination database name.
	TargetDbName string `json:"target_db_name"`
	// Source database name.
	SourceTableName string `json:"source_table_name"`
	// Name of a table in the destination database.
	TargetTableName string `json:"target_table_name"`
	// Number of rows in the table of the source database.
	SourceRowNum int32 `json:"source_row_num"`
	// Number of rows in the table of the destination database.
	TargetRowNum int32 `json:"target_row_num"`
	// Difference between the tables in the source and destination databases.
	DiffRowNum int32 `json:"diff_row_num"`
	// Row comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	LineCompareResult string `json:"line_compare_result,omitempty"`
	// Value comparison result.
	// Values:
	// CONSISTENT: consistent
	// INCONSISTENT: inconsistent
	// COMPARING: The comparison is in progress
	// WAITING_FOR_COMPARISON: waiting for comparison
	// FAILED_TO_COMPARE: comparison failure
	// TARGET_DB_NOT_EXIT-Destination database does not exist
	// CAN_NOT_COMPARE-Cannot be compared
	ContentCompareResult string `json:"content_compare_result"`
	// Provides additional information.
	Message string `json:"message,omitempty"`
}

type ContentCompareResultDiffs struct {
	// Source database name.
	SourceDbName string `json:"source_db_name"`
	// Table name of the source database.
	SourceTableName string `json:"source_table_name"`
	// The value comparison results are different.
	ContentCompareDiff []ContentCompareDiff `json:"ContentCompareDiff"`
	// Total number of differences in the value comparison result.
	ContentCompareDiffCount int32 `json:"content_compare_diff_count"`
}

type ContentCompareDiff struct {
	// Query the SQL statements of the destination database.
	TargetSelectSql string `json:"target_select_sql,omitempty"`
	// Query the SQL statements of the source database.
	SourceSelectSql string `json:"source_select_sql,omitempty"`
	// Key value list of the source database.
	SourceKeyValue []string `json:"source_key_value"`
	// Key value list of the destination database.
	TargetKeyValue []string `json:"target_key_value"`
}

type CompareTaskListResult struct {
	// List of comparison tasks.
	CompareTaskList []CompareTaskList `json:"CompareTaskList,omitempty"`
	// Total number of comparison tasks.
	CompareTaskListCount int32 `json:"compare_task_list_count,omitempty"`

	ErrorMsg  string `json:"error_msg,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}

type CompareTaskList struct {
	// ID of a comparison task.
	CompareTaskId string `json:"compare_task_id"`
	// Type of comparison task.
	CompareType string `json:"compare_type"`
	// Status of a comparison task.
	// Values:
	// RUNNING: The instance is running.
	// WAITING_FOR_RUNNING: waiting to be started
	// SUCCESSFUL: complete
	// FAILED: The migration task failed.
	// CANCELLED: canceled
	// TIMEOUT_INTERRUPT: timeout interrupt
	// FULL_DOING: Full verification is in progress
	// INCRE_DOING: incremental verification in progress
	CompareTaskStatus string `json:"compare_task_status"`
	// Comparison start time
	CreateTime string `json:"create_time"`
	// Comparison end time
	EndTime string `json:"end_time,omitempty"`
}
