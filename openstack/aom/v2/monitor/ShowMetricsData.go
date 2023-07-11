package monitor

import "github.com/opentelekomcloud/gophertelekomcloud"

type ShowMetricsDataQuery struct {
	// Filled value for breakpoints in monitoring data. Default value: -1.
	//
	// -1: Breakpoints are filled with -1.
	//
	// 0: Breakpoints are filled with 0.
	//
	// null: Breakpoints are filled with null.
	//
	// average: Breakpoints are filled with the average value of adjacent valid data. If there is no valid data, breakpoints are filled with null.
	FillValue string `q:"fillValue"`
}

type ShowMetricsDataOpts struct {
	// List of metrics.
	//
	// Value Range
	//
	// A maximum of 20 metrics are supported.
	Metrics []MetricQuery `json:"metrics" required:"true"`
	// Data monitoring granularity.
	//
	// Value Range
	//
	// Enumerated value. Options:
	//
	// 60: The data monitoring granularity is 1 minute.
	//
	// 300: The data monitoring granularity is 5 minutes.
	//
	// 900: The data monitoring granularity is 15 minutes.
	//
	// 3600: The data monitoring granularity is 1 hour.
	Period *int `json:"period" required:"true"`
	// Query time period. For example, -1.-1.5 indicates the last 5 minutes. 1501545600000.1501632000000.1440 indicates the fixed time period from 08:00:00 on August 1, 2017 to 08:00:00 on August 2, 2017.
	//
	// Time range/Period <= 1440
	//
	// The timerange and period must use the same unit.
	//
	// Value Range
	//
	// Format: start time (UTC, in ms).end time (UTC, in ms).number of minutes in the time period. When the start time and end time are -1, it indicates the latest N minutes. N indicates the time period by the granularity of minute.
	Timerange string `json:"timerange" required:"true"`
	// Statistic.
	//
	// Value Range
	//
	// maximum, minimum, sum, average, or sampleCount.
	Statistics []string `json:"statistics" required:"true"`
}

type MetricQuery struct {
	// Metric namespace.
	//
	// Value Range
	//
	// PAAS.CONTAINER, PAAS.NODE, PAAS.SLA, PAAS.AGGR, and CUSTOMMETRICS.
	Namespace string `json:"namespace" required:"true"`
	// Metric name.
	//
	// Value Range
	//
	// 1-255 characters.
	MetricName string `json:"metricName" required:"true"`
	// Metric dimension.
	//
	// dimensions.name: dimension name. Example: appName.
	//
	// dimensions.value: dimension value, such as a specific application name.
	//
	// Value Range
	//
	// Neither the array nor the name/value of any dimension in the array can be left blank.
	Dimensions []Dimension `json:"dimensions" required:"true"`
}

func ShowMetricsData(client *golangsdk.ServiceClient, opts ShowMetricsDataOpts, q ShowMetricsDataQuery) (*ShowMetricsDataResponse, error) {
	// POST /v2/{project_id}/ams/metricdata
}

type ShowMetricsDataResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Metric list.
	Metrics []Metrics `json:"metrics"`
}

type Metrics struct {
	// Query parameters.
	Metric Metric `json:"metric"`
	// Key metric.
	DataPoints []DataPoint `json:"dataPoints"`
}

type DataPoint struct {
	// Timestamp
	Timestamp int64 `json:"timestamp"`
	// Time series unit.
	Unit string `json:"unit"`
	// Statistic.
	Statistics []Statistic `json:"statistics"`
}

type Statistic struct {
	// Statistic.
	Statistic string `json:"statistic"`
	// Statistical result.
	Value float64 `json:"value"`
}
