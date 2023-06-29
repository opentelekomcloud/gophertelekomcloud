package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchLimitSpeedOpts struct {
	SpeedLimits []LimitSpeedReq `json:"speed_limits"`
}

type LimitSpeedReq struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Request body of flow control information.
	SpeedLimit []SpeedLimitInfo `json:"speed_limit"`
}

type SpeedLimitInfo struct {
	// Start time (UTC) of flow control. The start time is an integer in hh:mm format and the minutes part is ignored.
	// hh indicates the hour, for example, 01:00.
	Begin string `json:"begin" required:"true"`
	// End time (UTC) in the format of hh:mm, for example, 15:59. The value must end with 59.
	End string `json:"end" required:"true"`
	// Speed. The value ranges from 1 to 9,999, in MB/s.
	Speed string `json:"speed" required:"true"`
	// Whether the UTC time is used.
	IsUtc *bool `json:"is_utc,omitempty"`
}

func BatchSetSpeed(client *golangsdk.ServiceClient, opts BatchLimitSpeedOpts) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/jobs/batch-limit-speed
	raw, err := client.Put(client.ServiceURL("jobs", "batch-limit-speed"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchUpdateDatabaseObjectOpts struct {
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

type BatchJobsResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int         `json:"count,omitempty"`
}
