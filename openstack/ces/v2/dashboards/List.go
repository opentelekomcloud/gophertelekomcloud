package dashboards

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying dashboards.
type ListOpts struct {
	// Specifies whether to filter by favorite status.
	// This parameter requires enterprise_id to be specified.
	IsFavorite *bool `q:"is_favorite,omitempty"`
	// Specifies the dashboard name for filtering.
	// It can contain 1 to 128 characters.
	DashboardName string `q:"dashboard_name,omitempty"`
	// Specifies the dashboard ID for filtering.
	// The value starts with "db" and is followed by 22 characters.
	DashboardId string `q:"dashboard_id,omitempty"`
	// Specifies the enterprise project ID.
	EnterpriseId string `q:"enterprise_id,omitempty"`
}

// Dashboard represents a dashboard in the list response.
type Dashboard struct {
	// Specifies the dashboard ID.
	DashboardId string `json:"dashboard_id"`
	// Specifies the dashboard name.
	DashboardName string `json:"dashboard_name"`
	// Specifies the enterprise project ID.
	EnterpriseId string `json:"enterprise_id"`
	// Specifies the name of the user who created the dashboard.
	CreatorName string `json:"creator_name"`
	// Specifies the time when the dashboard was created.
	// The value is a UNIX timestamp in milliseconds.
	CreateTime int64 `json:"create_time"`
	// Specifies how many graphs will be displayed in each row.
	// Possible values: 0, 1, 2, 3. Default: 3
	RowWidgetNum int `json:"row_widget_num"`
	// Specifies whether the dashboard is a favorite.
	IsFavorite bool `json:"is_favorite"`
}

// List queries dashboards.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Dashboard, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("dashboards").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/dashboards
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Dashboards []Dashboard `json:"dashboards"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Dashboards, err
}
