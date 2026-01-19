package dashboards

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for creating or copying a dashboard.
type CreateOpts struct {
	// Specifies the dashboard name.
	// It can contain 1 to 128 characters.
	// Only letters, digits, underscores (_), hyphens (-), and Chinese characters are allowed.
	DashboardName string `json:"dashboard_name" required:"true"`
	// Specifies the enterprise project ID.
	// The value can be a UUID or "0".
	EnterpriseId string `json:"enterprise_id,omitempty"`
	// Specifies the dashboard ID to copy from.
	// If this parameter is specified, the dashboard is copied.
	// The value starts with "db" and is followed by 22 characters.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Specifies how many graphs will be displayed in each row.
	// Possible values: 0, 1, 2, 3. Default: 0
	// 0 indicates auto layout.
	RowWidgetNum *int `json:"row_widget_num,omitempty"`
}

// Create creates a new dashboard or copies an existing one.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/dashboards
	raw, err := client.Post(client.ServiceURL("dashboards"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		DashboardId string `json:"dashboard_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.DashboardId, err
}
