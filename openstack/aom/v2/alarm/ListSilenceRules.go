package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListSilenceRules is used to obtain the silence rule list.
func ListSilenceRules(client *golangsdk.ServiceClient) ([]SilenceRule, error) {
	// GET /v2/{project_id}/alert/mute-rules
	raw, err := client.Get(client.ServiceURL("alert", "mute-rules"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	// the response body is a bare JSON array of MuteRule objects
	var res []SilenceRule
	err = extract.Into(raw.Body, &res)
	return res, err
}

// SilenceRule represents a silence rule returned by the list API.
type SilenceRule struct {
	// Creation time (UTC timestamp, in milliseconds).
	CreateTime int64 `json:"create_time"`
	// Rule description.
	Desc string `json:"desc"`
	// Match condition of the rule.
	// Up to 10 serial or parallel conditions can be created.
	Match [][]SilenceMatch `json:"match"`
	// Time when the rule takes effect.
	MuteConfig SilenceMuteConfig `json:"mute_config"`
	// Rule name.
	// Enter a maximum of 100 characters and do not start or end with an underscore (_).
	// Only letters, digits, and underscores are allowed.
	Name string `json:"name"`
	// Time zone.
	Timezone string `json:"timezone"`
	// Modification time (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time"`
	// User ID.
	UserId string `json:"user_id"`
}
