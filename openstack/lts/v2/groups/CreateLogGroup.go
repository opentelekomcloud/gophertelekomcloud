package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Name of the log group to be created.
	// Minimum length: 1 character
	// Maximum length: 64 characters
	// Enumerated value:
	// lts-group-01nh
	LogGroupName string `json:"log_group_name" required:"true"`
	// Log retention duration, in days (fixed to 7 days).
	TTLInDays int `json:"ttl_in_days"`
}

func CreateLogGroup(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/groups
	raw, err := client.Post(client.ServiceURL("groups"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"log_group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
