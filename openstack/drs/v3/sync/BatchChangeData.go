package sync

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDataTransformationOpts struct {
	Jobs []CheckDataTransformationReq `json:"jobs"`
}

type CheckDataTransformationReq struct {
	// Task ID.
	JobId string `json:"job_id,omitempty"`
	// Provides object information.
	ObjectInfo []DatabaseObjectVo `json:"object_info,omitempty"`
	// Processing information.
	TransformationInfo TransformationInfo `json:"transformation_info"`
	// Configuration information. If there are multiple associated tables, generate multiple configuration rules.
	// The data that meets the configuration conditions is temporarily stored in the cache and used in the data filtering scenario.
	// The database name and table name can contain digits, letters, and underscores (_).
	// Ensure that the column names, primary keys, and indexes are the same as those in the source database.
	ConfigTransformation ConfigTransformationVo `json:"config_transformation,omitempty"`
}

type DatabaseObjectVo struct {
	// Database name and database table name. For example, the format is lxl_test1-*-*-test_1,
	// where lxl_test1 is the database name and test_1 is the table name.
	Id string `json:"id,omitempty"`
	// Whether to select advanced configuration. The value is true.
	Select string `json:"select,omitempty"`
	// Level of a processing object. The default value is 0.
	LeavesNum string `json:"leavesNum"`
}

type TransformationInfo struct {
	// The processing rule value is contentConditionalFilter.
	// The configuration rule value is configConditionalFilter.
	// Values:
	// contentConditionalFilter
	// configConditionalFilter
	TransformationType string `json:"transformation_type"`
	// Filter criteria. The processing rule value is a SQL statement,
	// and the configuration rule value is config. The value contains a maximum of 256 characters.
	Value string `json:"value"`
}

type ConfigTransformationVo struct {
	// Database-name.Table-name, for example, lxl_test1.test_1, where lxl_test1 is the database name and test_1 is the table name.
	DbTableName string `json:"db_table_name"`
	// Database name. The value contains a maximum of 256 characters.
	DbName string `json:"db_name"`
	// Table name The value contains a maximum of 256 characters.
	TableName string `json:"table_name"`
	// Column name The value contains a maximum of 256 characters.
	ColNames string `json:"col_names"`
	// Primary key or unique index The value contains a maximum of 256 characters.
	PrimKeyOrIndex string `json:"prim_key_or_index"`
	// Index that requires optimization. The value contains a maximum of 256 characters.
	Indexs string `json:"indexs"`
	// Filtering criteria. The value contains a maximum of 256 characters.
	Values string `json:"values"`
}

func BatchChangeData(client *golangsdk.ServiceClient, opts BatchDataTransformationOpts) (*BatchChangeDataResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-transformation
	raw, err := client.Post(client.ServiceURL("jobs", "batch-transformation"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchChangeDataResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchChangeDataResponse struct {
	Results []DataTransformationResp `json:"results,omitempty"`
	Count   int32                    `json:"count,omitempty"`
}

type DataTransformationResp struct {
	// Task ID.
	Id string `json:"id,omitempty"`
	// Status Values:
	// success
	// failed
	Status string `json:"status,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
}
