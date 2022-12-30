package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListFlavorOpts struct {
	// Index offset.
	// If offset is set to N, the resource query starts from the N+1 piece of data. The default value is 0, indicating that the query starts from the first piece of data.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// Maximum number of specifications that can be queried
	// Value range: 1-100
	// If this parameter is not transferred, the first 100 pieces of specification information can be queried by default.
	Limit int `q:"limit"`
	// Specifies the database type. The value is DDS-Community.
	EngineName string `q:"engine_name"`
	// DB version number.
	EngineVersion string `q:"engine_version"`
}

func List(client *golangsdk.ServiceClient, opts ListFlavorOpts) (*ListResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3.1/{project_id}/flavors
	raw, err := client.Get(client.ServiceURL("flavors")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Flavors    []FlavorResponse `json:"flavors"`
	TotalCount int              `json:"total_count"`
}

type FlavorResponse struct {
	// Indicates the engine name.
	EngineName string `json:"engine_name"`
	// Indicates the node type. DDS contains the following types of nodes:
	//
	// mongos
	// shard
	// config
	// replica
	// single
	Type string `json:"type"`
	// Indicates the number of vCPUs.
	Vcpus string `json:"vcpus"`
	// Indicates the memory size in gigabyte (GB).
	Ram string `json:"ram"`
	// Indicates the resource specification code.
	//
	// Example: dds.mongodb.s2.xlarge.4.shard
	//
	// NOTE:
	// dds.mongodb: indicates the DDS service.
	// s2.xlarge.4: indicates the performance specification, which is high memory.
	// shard: indicates the node type.
	// When querying the specifications, check whether the specifications are of the same series. The specification series includes general-purpose (s6), enhanced (c3), and enhanced II (c6).
	SpecCode string `json:"spec_code"`
	// Indicates the status of specifications in an AZ. Its value can be any of the following:
	//
	// normal: indicates that the specifications are on sale.
	// unsupported: indicates that the DB instance specifications are not supported.
	// sellout: indicates the specifications are sold out.
	// NOTE:
	// ReplicaSet flavors supports cross AZ creation in case "eu-de-01,eu-de-02,eu-de-03": "normal".
	AZStatus map[string]string `json:"az_status"`
	// Database versions
	//
	// For example, DDS mongos node, {"3.4", "4.0"}
	EngineVersions []string `json:"engine_versions"`
}
