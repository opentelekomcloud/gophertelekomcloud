package loggroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Specifies the log group name.
	LogGroupName string `json:"log_group_name" required:"true"`
	// Specifies the log expiration time. The value is fixed to 7 days.
	TTL int `json:"ttl_in_days,omitempty"`
}

// Create a log group with given parameters.
func Create(client *golangsdk.ServiceClient, ops CreateOpts) (string, error) {
	b, err := build.RequestBody(ops, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("log-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})

	var res struct {
		ID string `json:"log_group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
