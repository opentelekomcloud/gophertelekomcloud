package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateSilenceRuleOpts struct {
	// Creation time (UTC timestamp, in milliseconds).
	CreateTime int64 `json:"create_time,omitempty"`
	// Rule description.
	Desc string `json:"desc,omitempty"`
	// Match condition of the rule.
	// Up to 10 serial or parallel conditions can be created.
	Match [][]SilenceMatch `json:"match" required:"true"`
	// Time when the rule takes effect.
	MuteConfig SilenceMuteConfig `json:"mute_config" required:"true"`
	// Rule name.
	// Enter a maximum of 100 characters and do not start or end with an underscore (_).
	// Only letters, digits, and underscores are allowed.
	Name string `json:"name" required:"true"`
	// Time zone.
	Timezone string `json:"timezone" required:"true"`
	// Modification time (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time,omitempty"`
	// User ID.
	UserId string `json:"user_id,omitempty"`
}

// UpdateSilenceRule is used to modify a silence rule.
// The rule is identified by name in the request body.
func UpdateSilenceRule(client *golangsdk.ServiceClient, opts UpdateSilenceRuleOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/alert/mute-rules
	_, err = client.Put(client.ServiceURL("alert", "mute-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
