package util

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListStatsOpts struct {
	Filter string `q:"filter" required:"true"`
	Period string `q:"period"`
	Option string `q:"option"`
	Limit  string `q:"limit"`
	Marker string `q:"marker"`
}

func ListStats(client *golangsdk.ServiceClient, opts ListStatsOpts) (*ListStatsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("fgs", "functions", "statistics").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListStatsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListStatsResponse struct {
	Count      []MonthUsed `json:"count"`
	Gbs        []MonthUsed `json:"gbs"`
	GpuGbs     []MonthUsed `json:"gpu_gbs"`
	Statistics *StatResp   `json:"monitor_data"`
}

type MonthUsed struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

type StatResp struct {
	Count               []SlaReport `json:"count"`
	Duration            []SlaReport `json:"duration"`
	FailCount           []SlaReport `json:"fail_count"`
	MaxDuration         []SlaReport `json:"max_duration"`
	MinDuration         []SlaReport `json:"min_duration"`
	RejectCount         []SlaReport `json:"reject_count"`
	FuncErrorCount      []SlaReport `json:"function_error_count"`
	SysErrorCount       []SlaReport `json:"system_error_count"`
	ReservedInstanceNum []SlaReport `json:"reserved_instance_num"`
	ConcurrencyNum      []SlaReport `json:"concurrency_num"`
}

type SlaReport struct {
	TimeStamp int     `json:"timestamp"`
	Value     float64 `json:"value"`
}
