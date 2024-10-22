package schemas

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QuerySchemasOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of schemas to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 schemas are queried by default.
	Limit int `q:"limit"`
}

// QuerySchemas is used to query schemas of a DDM instance.
func QuerySchemas(client *golangsdk.ServiceClient, instanceId string, opts QuerySchemasOpts) ([]GetDatabaseInfo, error) {
	// GET /v1/{project_id}/instances/{instance_id}/databases?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "databases").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SchemasPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSchemas(pages)
}

func ExtractSchemas(r pagination.NewPage) ([]GetDatabaseInfo, error) {
	var querySchemasResponse QuerySchemasResponse
	err := extract.Into(bytes.NewReader((r.(SchemasPage)).Body), &querySchemasResponse)
	return querySchemasResponse.Databases, err
}

type SchemasPage struct {
	pagination.NewSinglePageBase
}

type QuerySchemasResponse struct {
	// Schema information.
	Databases []GetDatabaseInfo `json:"databases"`
	// Total records.
	Total int `json:"total"`
}

// GetDatabaseInfo represents the schema details of a database.
type GetDatabaseInfo struct {
	// Schema name
	Name string `json:"name"`
	// Sharding mode of the schema
	// cluster: indicates that the schema is in sharded mode.
	// single: indicates that the schema is in unsharded mode.
	ShardMode string `json:"shard_mode"`
	// Number of shards in the same working mode
	ShardNumber int `json:"shard_number"`
	// Schema status
	Status string `json:"status"`
	// Time when the schema is created
	Created int `json:"created"`
	// RDS instances associated with the schema
	UsedRds []GetDatabasesUsedRds `json:"used_rds"`
	// Number of shards per RDS instance
	ShardUnit int `json:"shard_unit"`
}
