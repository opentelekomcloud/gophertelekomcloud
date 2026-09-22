package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Message engine type.
	//    rocketmq: RocketMQ engine.
	//    reliability: RocketMQ engine alias.
	Engine string `q:"engine,required"`
	// Instance name.
	Name string `q:"name,omitempty"`
	// Instance ID.
	InstanceID string `q:"instance_id,omitempty"`
	// Instance status.
	Status string `q:"status,omitempty"`
	// Whether to return the number of instances that fail to be created.
	// Options: 'true', 'false'.
	IncludeFailure string `q:"include_failure,omitempty"`
	// Whether to search for the instance that precisely matches a specified
	// instance name. Options: 'true', 'false'. Default: 'false'.
	ExactMatchName string `q:"exact_match_name,omitempty"`
	// Enterprise project ID.
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
	// Maximum number of instances that can be returned in a query.
	// Range: 1-50. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Instance list.
	Instances []Instance `json:"instances"`
	// Number of instances.
	InstanceNum int `json:"instance_num"`
}

// List the RocketMQ instances of an account by the specified conditions.
// Send GET to /v2/{project_id}/instances
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
