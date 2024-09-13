package schemas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateSchemaOpts struct {

	// Schema Information. Value is of the type CreateSchemaDetail
	Databases []CreateDatabaseDetail `json:"databases" required:"true"`
}

type CreateDatabaseDetail struct {
	// Schema name. It can include 2 to 48 characters. It must start with a letter.
	// It should contain only lowercase letters, digits, and underscores (_).
	// It cannot contain keywords: information_schema, mysql, performance_schema, or sys.
	Name string `json:"name" required:"true"`
	// Sharding mode of the schema. The value can be:
	// cluster: indicates that the schema is in sharded mode.
	// single: indicates that the schema is in unsharded mode.
	// Enumerated values: cluster, single
	ShardMode string `json:"shard_mode" required:"true"`
	// Number of shards in the same working mode
	// If ShardUnit is not empty, the value is the product of shard_unit multiplied by the associated RDS instances.
	// If ShardUnit is left blank, the value must be greater than the number of associated RDS instances
	// and less than or equal to the product of the associated RDS instances multiplied by 64.
	ShardNumber int `json:"shard_number" required:"true"`
	// Number of shards per RDS instance This parameter is optional.
	// The value is 1 if the schema is unsharded.
	// The value ranges from 1 (min) to 64 (max) if the schema is sharded.
	ShardUnit int `json:"shard_unit,omitempty"`
	// RDS instances associated with the schema
	UsedRds []DatabaseInstabcesParam `json:"used_rds" required:"true"`
}

type DatabaseInstabcesParam struct {
	// ID of the RDS instance associated with the schema
	Id string `json:"id" required:"true"`
	// Username for logging in to the associated RDS instance
	AdminUser string `json:"adminUser" required:"true"`
	// Password for logging in to the associated RDS instance
	AdminPassword string `json:"adminPassword" required:"true"`
}

// This function is used to create a schema.
// Before creating a schema, ensure that you have associated RDS instances with your DDM instance and that the RDS instances are not associated with other DDM instances.
func CreateSchema(client *golangsdk.ServiceClient, instanceId string, opts CreateSchemaOpts) (*CreateSchemaResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/instances/{instance_id}/databases
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "databases"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateSchemaResponse
	return &res, extract.Into(raw.Body, &res)
}

type CreateSchemaResponse struct {
	// Schema information.
	Databases []CreateDatabaseDetailResponses `json:"databases"`
	// ID of the job for creating a schema.
	JobId string `json:"job_id"`
}

type CreateDatabaseDetailResponses struct {
	// Schema name
	Name string `json:"name"`
}
