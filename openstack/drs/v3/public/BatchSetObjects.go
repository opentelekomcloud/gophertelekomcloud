package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchUpdateDatabaseObjectReq struct {
	Jobs []UpdateDatabaseObjectReq `json:"jobs"`
}

type UpdateDatabaseObjectReq struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Whether to select an object. If this parameter is not set, the default value is No.
	// Yes: Customize the objects to be migrated.
	// No: Migrate all VMs.
	Selected bool `json:"selected,omitempty"`
	// Whether to perform database-level synchronization.
	SyncDatabase bool `json:"sync_database,omitempty"`
	// Data object selection information. This parameter is mandatory when selected is set to true.
	Job []DatabaseInfo `json:"job,omitempty"`
}

type DatabaseInfo struct {
	// When object_type is set to database, this parameter indicates the database name.
	// If object_type is set to table or view, set the field value by referring to the example.
	Id string `json:"id,omitempty"`
	// Database name. This parameter is mandatory when object_type is set to table or view.
	ParentId string `json:"parent_id,omitempty"`
	// Type. Values:
	// ⦁	database
	// ⦁	table
	// ⦁	schema
	// ⦁	view
	ObjectType string `json:"object_type,omitempty"`
	// Database object name, database name, table name, and view name.
	ObjectName string `json:"object_name,omitempty"`
	// Alias, which is the new mapped name.
	ObjectAliasName string `json:"object_alias_name,omitempty"`
	// Whether to migrate the database objects. true indicates that the database objects will be migrated.
	// false indicates that the database objects will not be migrated.
	// partial indicates some tables in the database will be migrated.
	// If this parameter is not specified, the default value is false.
	Select string `json:"select,omitempty"`
}

func BatchSetObjects(client *golangsdk.ServiceClient, opts BatchUpdateDatabaseObjectReq) (*BatchSetObjectsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/jobs/batch-select-objects
	raw, err := client.Put(client.ServiceURL("jobs", "batch-select-objects"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchSetObjectsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchSetObjectsResponse struct {
	AllCounts int64                `json:"all_counts,omitempty"`
	Results   []DatabaseObjectResp `json:"results,omitempty"`
}

type DatabaseObjectResp struct {
	// Task ID.
	JobId string `json:"job_id,omitempty"`
	// The status that indicates that objects are selected.
	Status bool `json:"status,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
