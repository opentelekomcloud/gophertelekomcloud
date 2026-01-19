package dashboards

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchDeleteOpts contains the options for batch deleting dashboards.
type BatchDeleteOpts struct {
	// Specifies the list of dashboard IDs to delete.
	// A maximum of 30 dashboard IDs are supported.
	DashboardIds []string `json:"dashboard_ids,omitempty"`
}

// BatchDeleteResult represents the result of deleting a single dashboard.
type BatchDeleteResult struct {
	// Specifies the dashboard ID.
	DashboardId string `json:"dashboard_id"`
	// Specifies the deletion result.
	// Possible values: successful, error
	RetStatus string `json:"ret_status"`
	// Specifies the error message if deletion failed.
	ErrorMsg string `json:"error_msg,omitempty"`
}

// BatchDelete batch deletes dashboards.
func BatchDelete(client *golangsdk.ServiceClient, opts BatchDeleteOpts) ([]BatchDeleteResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/dashboards/batch-delete
	raw, err := client.Post(client.ServiceURL("dashboards", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Dashboards []BatchDeleteResult `json:"dashboards"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Dashboards, err
}
