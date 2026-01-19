package dashboards

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the options for modifying a dashboard.
type UpdateOpts struct {
	// Specifies the dashboard name.
	// It can contain 1 to 128 characters.
	// Only letters, digits, underscores (_), hyphens (-), and Chinese characters are allowed.
	DashboardName string `json:"dashboard_name,omitempty"`
	// Specifies whether the dashboard is a favorite.
	IsFavorite *bool `json:"is_favorite,omitempty"`
	// Specifies how many graphs will be displayed in each row.
	// Possible values: 0, 1, 2, 3. Default: 3
	// 0 indicates auto layout.
	RowWidgetNum *int `json:"row_widget_num,omitempty"`
}

// Update modifies a dashboard.
func Update(client *golangsdk.ServiceClient, dashboardId string, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/dashboards/{dashboard_id}
	_, err = client.Put(client.ServiceURL("dashboards", dashboardId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
