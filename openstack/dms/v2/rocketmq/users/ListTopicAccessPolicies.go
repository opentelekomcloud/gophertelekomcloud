package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListTopicAccessPolicies queries the users granted permissions for a topic.
// Send GET to /v2/{project_id}/instances/{instance_id}/topics/{topic}/accesspolicy
func ListTopicAccessPolicies(client *golangsdk.ServiceClient, instanceID, topic string, opts ListAccessPoliciesOpts) (*ListAccessPoliciesResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "topics", topic, "accesspolicy").
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
