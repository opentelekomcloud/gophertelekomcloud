package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateSilenceRuleOpts struct {
	// Time when the silence rule was created (UTC timestamp, in milliseconds).
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
	Timezone string `json:"timezone,omitempty"`
	// Time when the silence rule was modified (UTC timestamp, in milliseconds).
	UpdateTime int64 `json:"update_time,omitempty"`
	// User ID.
	UserId string `json:"user_id,omitempty"`
}

type SilenceMatch struct {
	// Key specified in the metadata for matching.
	Key string `json:"key" required:"true"`
	// Match mode.
	// EXIST: Exist.
	// REGEX: Regular expression.
	// EQUALS: Equal to.
	Operate string `json:"operate" required:"true"`
	// Value corresponding to the key.
	// If operate is set to EXIST, this parameter is left blank.
	Value []string `json:"value,omitempty"`
}

type SilenceMuteConfig struct {
	// End time of the silence rule (UTC timestamp, in milliseconds).
	// If the silence rule is permanent, set this parameter to 0.
	// No omitempty/required: 0 is a meaningful value (permanent rule).
	EndsAt int64 `json:"ends_at"`
	// When type is set to WEEKLY or MONTHLY, scope cannot be empty.
	Scope []int `json:"scope,omitempty"`
	// Start time of the silence rule (UTC timestamp, in milliseconds).
	// No omitempty/required: 0 is a meaningful value (e.g. DAILY starting at midnight).
	StartsAt int64 `json:"starts_at"`
	// Type of the time for the silence rule to take effect.
	// FIXED: Fixed time.
	// DAILY: Certain time every day.
	// WEEKLY: Certain time every week.
	// MONTHLY: Certain time every month.
	Type string `json:"type" required:"true"`
}

// CreateSilenceRule is used to add a silence rule.
func CreateSilenceRule(client *golangsdk.ServiceClient, opts CreateSilenceRuleOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alert/mute-rules
	_, err = client.Post(client.ServiceURL("alert", "mute-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
