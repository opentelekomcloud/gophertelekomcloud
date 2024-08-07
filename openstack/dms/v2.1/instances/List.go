package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Message engine. Value: kafka.
	// Default: kafka
	Engine string `q:"engine,omitempty"`
	// Instance name.
	Name string `q:"name,omitempty"`
	// Instance ID.
	InstanceId string `q:"instance_id,omitempty"`
	// Instance status. For details, see Instance Status.
	Status string `q:"status,omitempty"`
	// Indicates whether to return the number of instances that fail to be created.
	// If the value is true, the number of instances that failed to be created is returned. If the value is false, the number is not returned.
	IncludeFailure string `q:"include_failure,omitempty"`
	// Whether to search for the instance that precisely matches a specified instance name.
	// The default value is false*, indicating that a fuzzy search is performed based on a specified instance name. If the value is true, the instance that precisely matches a specified instance name is queried.
	ExactMatchName string `q:"exact_match_name,omitempty"`
	// Enterprise project ID.
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
	// Offset, which is the position where the query starts. The value must be greater than or equal to 0.
	Offset string `q:"offset,omitempty"`
	// Maximum number of instances returned in the current query. The default value is 10. The value ranges from 1 to 50.
	Limit string `q:"limit,omitempty"`
}

// List is used to query the instances of an account by the specified conditions.
// Send GET /v2/{project_id}/instances
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(ResourcePath).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Instances  []Instance `json:"instances"`
	TotalCount int        `json:"instance_num"`
}
