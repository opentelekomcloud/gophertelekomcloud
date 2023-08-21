package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateReferenceTableOpts struct {
	// Reference table name. The value can contain a maximum of 64 characters.
	// Only digits, letters, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Reference table type. For details, see the enumeration values as followed.
	// Enumeration values:
	// url
	// params
	// ip
	// cookie
	// referer
	// user-agent
	// header
	// response_code
	// response_header
	// response_body
	Type string `json:"type" required:"true"`
	// Value of the reference table.
	Values []string `json:"values"`
}

// CreateReferenceTable will create a reference table on the values in CreateOpts.
func CreateReferenceTable(client *golangsdk.ServiceClient, opts CreateReferenceTableOpts) (*ReferenceTable, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/valuelist
	raw, err := client.Post(client.ServiceURL("waf", "valuelist"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res ReferenceTable
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ReferenceTable struct {
	// ID of a reference table.
	ID string `json:"id"`
	// Reference table name.
	Name string `json:"name"`
	// Type
	Type string `json:"type"`
	// Reference table timestamp.
	CreatedAt string `json:"timestamp"`
	// Value of the reference table.
	Values []string `json:"values"`
	// Reference table description.
	Description string `json:"description"`
	Producer    int    `json:"producer"`
}
