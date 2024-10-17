package schemas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This API is used to query details about a schema.
// schemaName is the name of the schema to be queried, which is case-insensitive
func QuerySchemaDetails(client *golangsdk.ServiceClient, instanceId string, schemaName string) (*QuerySchemaDetailsResponse, error) {

	// GET /v1/{project_id}/instances/{instance_id}/databases/{ddm_dbname}
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "databases", schemaName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuerySchemaDetailsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QuerySchemaDetailsResponse struct {
	Database GetDatabaseResponseBean `json:"database"`
}

// GetDatabaseResponseBean represents the response for database details
type GetDatabaseResponseBean struct {
	// Schema name
	Name string `json:"name"`
	// Time when the schema is created
	Created string `json:"created"`
	// Schema status
	Status string `json:"status"`
	// Time when the DDM instance is last updated
	Updated string `json:"updated"`
	// Sharding information of the schema
	Databases []GetDatabases `json:"databases"`
	// Sharding mode of the schema
	// cluster: indicates that the schema is in sharded mode.
	// single: indicates that the schema is in unsharded mode.
	ShardMode string `json:"shard_mode"`
	// Number of shards in the same working mode
	ShardNumber int `json:"shard_number"`
	// Number of shards per RDS instance
	ShardUnit int `json:"shard_unit"`
	// IP address and port number for connecting to the schema
	DataVips []string `json:"dataVips"`
	// Associated RDS instances
	UsedRds []GetDatabasesUsedRds `json:"used_rds"`
}

// GetDatabases represents the details of a specific database shard
type GetDatabases struct {
	// Number of shards
	DbSlot int `json:"dbslot"`
	// Shard name
	Name string `json:"name"`
	// Shard status
	Status string `json:"status"`
	// Time when the shard is created
	Created string `json:"created"`
	// Time when the shard is last updated
	Updated string `json:"updated"`
	// ID of the RDS instance where the shard is located
	Id string `json:"id"`
	// Name of the RDS instance where the shard is located
	IdName string `json:"idName"`
}

// GetDatabasesUsedRds represents the RDS instances associated with the database schema
type GetDatabasesUsedRds struct {
	// Node ID of the associated RDS instance
	ID string `json:"id"`
	// Name of the associated RDS instance
	Name string `json:"name"`
	// Status of the associated RDS instance
	Status string `json:"status"`
	// Response message. This parameter is not returned if no abnormality occurs.
	ErrorMsg string `json:"error_msg,omitempty"`
}
