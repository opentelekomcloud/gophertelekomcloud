package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListGroupAccessPolicies queries the users granted permissions for a consumer group.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/groups/{group}/accesspolicy
func ListGroupAccessPolicies(client *golangsdk.ServiceClient, instanceID, group string, opts ListAccessPoliciesOpts) (*ListAccessPoliciesResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "groups", group, "accesspolicy").
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

	var res ListAccessPoliciesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
