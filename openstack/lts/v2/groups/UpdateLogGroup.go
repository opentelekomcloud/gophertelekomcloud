package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateLogGroupOpts struct {
	// Log group ID. For details about how to obtain a log group ID, see Obtaining the AccountID, Project ID, Log Group ID, and Log Stream ID.
	// Default value: None
	// Value length: 36 characters
	LogGroupId string `json:"-" required:"true"`
	// Log retention duration, in days (fixed to 7 days).
	TTLInDays int32 `json:"ttl_in_days" required:"true"`
}

func UpdateLogGroup(client *golangsdk.ServiceClient, opts UpdateLogGroupOpts) (*LogGroup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/groups/{log_group_id}
	raw, err := client.Post(client.ServiceURL("groups", opts.LogGroupId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LogGroup
	err = extract.Into(raw.Body, &res)
	return &res, err
}
