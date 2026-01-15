package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying resources in an alarm rule.
type ListOpts struct {
	// Specifies the pagination offset. Default: 0
	Offset int `q:"offset"`
	// Specifies the number of records on each page. Default: 10, Max: 100
	Limit int `q:"limit"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the list of resources.
	Resources [][]Dimension `json:"resources"`
	// Specifies the total number of resources.
	Count int `json:"count"`
}

// List returns a list of resources in an alarm rule.
func List(client *golangsdk.ServiceClient, alarmId string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("alarms", alarmId, "resources").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarms/{alarm_id}/resources
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
