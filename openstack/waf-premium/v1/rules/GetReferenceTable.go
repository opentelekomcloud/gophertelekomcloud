package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetReferenceTable is used to query a reference table by ID.
func GetReferenceTable(client *golangsdk.ServiceClient, tableId string) (*ReferenceTable, error) {
	// GET /v1/{project_id}/waf/valuelist/{table_id}
	raw, err := client.Get(client.ServiceURL("waf", "valuelist", tableId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ReferenceTable
	return &res, extract.Into(raw.Body, &res)
}
