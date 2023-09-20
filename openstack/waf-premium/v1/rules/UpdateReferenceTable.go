package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateReferenceTableOpts struct {
	// Reference table name. The value can contain a maximum of 64 characters.
	// Only digits, letters, hyphens (-), underscores (_), and periods (.) are allowed
	Name string `json:"name" required:"true"`
	// Reference table type.
	Type string `json:"type" required:"true"`
	// Value of the reference table.
	Values []string `json:"values"`
	// Reference table description.
	Description string `json:"description"`
}

// UpdateReferenceTable is used to modify a reference table.
func UpdateReferenceTable(client *golangsdk.ServiceClient, tableId string, opts UpdateReferenceTableOpts) (*ReferenceTable, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/valuelist/{table_id}
	raw, err := client.Put(client.ServiceURL("waf", "valuelist", tableId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res ReferenceTable
	return &res, extract.Into(raw.Body, &res)
}
