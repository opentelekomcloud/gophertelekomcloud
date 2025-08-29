package snatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePrivateSnatOpts struct {
	// Provides supplementary information about the SNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the IDs of the transit IP addresses.
	// Constraints: A maximum number of 20 IDs is allowed.
	TransitIpIds []string `json:"transit_ip_ids,omitempty"`
}

// This function is used to update an SNAT rule.
func Update(client *golangsdk.ServiceClient, ruleId string, opts UpdatePrivateSnatOpts) (*SnatCommonResponse, error) {
	b, err := build.RequestBody(opts, "snat_rule")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/private-nat/snat-rules/{snat_rule_id}
	raw, err := client.Put(client.ServiceURL("private-nat", "snat-rules", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SnatCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
