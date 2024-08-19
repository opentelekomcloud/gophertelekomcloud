package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RecyclePolicyOpts struct {
	// CIDR block where the client is located
	RecyclePolicy *RecyclePolicy `json:"recycle_policy" required:"true"`
}

type RecyclePolicy struct {
	// The recycling policy is enabled and cannot be disabled.
	// true: The recycling policy is enabled.
	Enabled *bool `json:"enabled" required:"true"`
	// Policy retention duration (1 to 7 days).
	// The value is a positive integer. If this parameter is left blank, the policy is retained for 7 days by default.
	RetentionPeriodInDays int `json:"retention_period_in_days,omitempty"`
}

func SetRecyclePolicy(client *golangsdk.ServiceClient, instanceId string, opts RecyclePolicyOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/recycle-policy
	_, err = client.Put(client.ServiceURL("instances", "recycle-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
