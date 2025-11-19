package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetInspectionHistoriesOpts struct {
	InstanceId    string `json:"-" q:"-"`
	Offset        *int   `q:"offset"`
	Limit         *int   `q:"limit"`
	Order         string `q:"order"`
	SortField     string `q:"sort_field"`
	TargetVersion string `q:"target_version"`
	IsAvailable   *bool  `q:"is_available"`
}

type InspectionReport struct {
	Id             string `json:"id"`
	CheckTime      string `json:"check_time"`
	ExpirationTime string `json:"expiration_time"`
	TargetVersion  string `json:"target_version"`
	Result         string `json:"result"`
	Detail         string `json:"detail"`
}

type InspectionHistories struct {
	TotalCount        int                `json:"total_count"`
	InspectionReports []InspectionReport `json:"inspection_reports"`
}

func GetInspectionHistories(client *golangsdk.ServiceClient, opts GetInspectionHistoriesOpts) (*InspectionHistories, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", opts.InstanceId, "major-version", "inspection-histories").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res InspectionHistories
	err = extract.Into(raw.Body, &res)
	return &res, err
}
