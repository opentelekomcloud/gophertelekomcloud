package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains the parameters of the Update function.
type UpdateOpts struct {
	// Rule name.
	// It can contain 1 to 64 characters. Only letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.
	Name string `json:"name,omitempty"`
	// Destination IP address added to a rule.
	IPAddresses []IPInfo `json:"ipaddresses,omitempty"`
}

// Update is used to modify an endpoint rule.
// Send PUT /v2.1/resolverrules/{resolverrule_id}
func Update(client *golangsdk.ServiceClient, ruleID string, opts UpdateOpts) (*ResolverRule, error) {
	if ruleID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "ruleID"}
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(v21URL(client, "resolverrules", ruleID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res ResolverRule
	err = extract.IntoStructPtr(raw.Body, &res, "resolver_rule")
	return &res, err
}
