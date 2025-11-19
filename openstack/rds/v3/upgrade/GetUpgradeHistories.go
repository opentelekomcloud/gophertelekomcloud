package upgrade

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetUpgradeHistoriesOpts struct {
	InstanceId string `json:"-" q:"-"`
	Offset     *int   `q:"offset"`
	Limit      *int   `q:"limit"`
	Order      string `q:"order"`
	SortField  string `q:"sort_field"`
}

type Report struct {
	Id                       string `json:"id"`
	StartTime                string `json:"start_time"`
	EndTime                  string `json:"end_time"`
	SrcInstanceId            string `json:"src_instance_id"`
	SrcDatabaseVersion       string `json:"src_database_version"`
	DstInstanceId            string `json:"dst_instance_id"`
	DstDatabaseVersion       string `json:"dst_database_version"`
	Result                   string `json:"result"`
	IsPrivateIpChanged       bool   `json:"is_private_ip_changed"`
	PrivateIpChangeTime      string `json:"private_ip_change_time"`
	StatisticsCollectionMode string `json:"statistics_collection_mode"`
	Detail                   string `json:"detail"`
}

type Histories struct {
	TotalCount     int      `json:"total_count"`
	UpgradeReports []Report `json:"upgrade_reports"`
}

func GetUpgradeHistories(client *golangsdk.ServiceClient, opts GetUpgradeHistoriesOpts) (*Histories, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", opts.InstanceId, "major-version", "upgrade-histories").
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

	var res Histories
	err = extract.Into(raw.Body, &res)
	return &res, err
}
