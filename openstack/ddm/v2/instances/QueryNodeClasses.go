package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type QueryNodeClassesOpts struct {
	// Engine ID, which can be obtained by calling the API for querying DDM engine information.
	EngineId string `q:"engine_id" required:"true"`
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of node classes to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 node classes are queried by default.
	Limit int `q:"limit"`
}

// QueryNodeClasses is used to query DDM node classes available in an AZ.
func QueryNodeClasses(client *golangsdk.ServiceClient, opts QueryNodeClassesOpts) (*QueryNodeClassesResponse, error) {
	// GET /v2/{project_id}/flavors?engine_id={engine_id}?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("flavors").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res QueryNodeClassesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err

}

type QueryNodeClassesResponse struct {
	// Compute flavor information
	ComputeFlavorGroups []ComputeFlavorGroupsInfo `json:"computeFlavorGroups"`
}

type ComputeFlavorGroupsInfo struct {
	// Compute resource architecture type. The value can be x86 or ARM.
	GroupType string `json:"groupType"`
	// Compute flavors
	ComputeFlavors []ComputeFlavors `json:"computeFlavors"`
	// Which page the server starts returning items.
	Offset int `json:"offset"`
	// Number of records displayed on each page
	Limit int `json:"limit"`
	// Total number of compute flavors
	Total int `json:"total"`
}

// ComputeFlavors represents the details of a compute resource flavor.
type ComputeFlavors struct {
	// Flavor ID
	ID string `json:"id"`
	// Resource type code
	TypeCode string `json:"typeCode"`
	// VM flavor types recorded in DDM
	Code string `json:"code"`
	// VM flavor types recorded by the IaaS layer
	IaaSCode string `json:"iaasCode"`
	// Number of CPUs
	CPU string `json:"cpu"`
	// Memory size, in GB
	Mem string `json:"mem"`
	// Maximum number of connections
	MaxConnections string `json:"maxConnections"`
	// Compute resource type
	ServerType string `json:"serverType"`
	// Compute resource architecture type. The value can be x86 or ARM.
	Architecture string `json:"architecture"`
	// Status of the AZ where node classes are available.
	// The key is the AZ ID and the value is the AZ status.
	// The value can be:
	// - normal: indicates that the node classes in the AZ are available.
	// - unsupported: indicates that the node classes are not supported by the AZ.
	// - sellout: indicates that the node classes in the AZ are sold out.
	AZStatus map[string]string `json:"azStatus"`
	// Region status
	RegionStatus string `json:"regionStatus"`
	// Compute resource architecture type. The value can be x86 or ARM.
	GroupType string `json:"groupType"`
	// Engine type
	DBType string `json:"dbType"`
	// Extension field for storing AZ information
	ExtendFields map[string]string `json:"extendFields"`
}
