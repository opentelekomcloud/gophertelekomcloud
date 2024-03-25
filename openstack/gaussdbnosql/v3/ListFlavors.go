package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListFlavorsOpts struct {
	// Region where the instance is deployed.
	// Valid value:
	// The value cannot be empty. For details about how to obtain this parameter value, see Regions and Endpoints.
	Region string `q:"region"`
	// engine_name	No	Database type
	// If the value is cassandra, the GaussDB(for Cassandra) DB instance specifications are queried,
	// If this parameter is not transferred, the default value is cassandra.
	EngineName string `q:"engine_name,omitempty"`
}

func ListFlavors(client *golangsdk.ServiceClient, opts ListFlavorsOpts) (*ListFlavorsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/flavors
	raw, err := client.Get(client.ServiceURL("flavors")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListFlavorsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListFlavorsResponse struct {
	// Total number of records
	TotalCount int32 `json:"total_count"`
	// Instance specifications.
	Flavors []ListFlavorsResult `json:"flavors"`
}

type ListFlavorsResult struct {
	// Engine name
	EngineName string `json:"engine_name"`
	// Engine version
	EngineVersion string `json:"engine_version"`
	// Number of vCPUs
	Vcpus string `json:"vcpus"`
	// Memory size in megabytes (MB)
	Ram string `json:"ram"`
	// Resource specification code
	// Example: geminidb.cassandra.8xlarge.4
	// NOTE
	// geminidb.cassandra indicates GaussDB(for Cassandra).
	// 8xlarge.4 indicates the single-node specifications.
	SpecCode string `json:"spec_code"`
	// ID of the AZ that supports this specification
	// NOTE
	// This field is discarded and cannot be used.
	AvailabilityZone []string `json:"availability_zone"`
	// Status of specifications in an AZ. The value can be:
	// normal: indicates that the specifications are on sale.
	// unsupported: not supported.
	// sellout: indicates that the instance specifications are sold out.
	AzStatus string `json:"az_status"`
}
