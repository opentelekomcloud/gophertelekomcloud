package loggroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Log group name.
	// The configuration rules are as follows:
	// - Must be a string of 1 to 64 characters.
	// - Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	// The name cannot start or end with a period.
	LogGroupName string `json:"log_group_name" required:"true"`
	// Log expiration time. The value is fixed to 7 days.
	TTLInDays int `json:"ttl_in_days"`
}

func Create(client *golangsdk.ServiceClient, ops CreateOpts) (string, error) {
	b, err := build.RequestBody(ops, "")
	if err != nil {
		return "", err
	}

	// POST /v2.0/{project_id}/groups
	raw, err := client.Post(client.ServiceURL("groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"log_group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
