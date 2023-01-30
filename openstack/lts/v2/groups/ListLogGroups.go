package groups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListLogGroups(client *golangsdk.ServiceClient) ([]LogGroup, error) {
	// GET /v2/{project_id}/groups
	raw, err := client.Get(client.ServiceURL("groups"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []LogGroup
	err = extract.IntoSlicePtr(raw.Body, &res, "log_groups")
	return res, err
}

type LogGroup struct {
	// Time when a log group was created.
	CreationTime int64 `json:"creation_time"`
	// Log group name.
	// Minimum length: 1 character
	// Maximum length: 64 characters
	LogGroupName string `json:"log_group_name"`
	// Log group ID.
	// Value length: 36 characters
	LogGroupId string `json:"log_group_id"`
	// Log retention duration, in days (fixed to 7 days).
	TTLInDays int `json:"ttl_in_days"`
	// Log group tag.
	Tag map[string]string `json:"tag,omitempty"`
}
