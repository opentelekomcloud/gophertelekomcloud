package metricdata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchListMetricDataOpts struct {
	// Specifies the metric data. The maximum length of the array is 10.
	Metrics []Metric `json:"metrics" required:"true"`
	// Specifies the start time of the query.
	// The value is a UNIX timestamp and the unit is ms.
	// Set the value of from to at least one period earlier than the current time.
	// Rollup aggregates the raw data generated within a period to the start time of the period.
	// Therefore, if values of from and to are within a period,
	// the query result will be empty due to the rollup failure.
	// You are advised to set from to be at least one period earlier than the current time.
	// Take the 5-minute period as an example. If it is 10:35 now,
	// the raw data generated between 10:30 and 10:35 will be aggregated to 10:30.
	// Therefore, in this example, if the value of period is 5 minutes,
	// the value of from should be 10:30 or earlier.
	From int64 `json:"from" required:"true"`
	// Specifies the end time of the query.
	// The value is a UNIX timestamp and the unit is ms.
	// The value of parameter from must be earlier than that of parameter to.
	To int64 `json:"to" required:"true"`
	// Specifies how often Cloud Eye aggregates data.
	//
	// Possible values are:
	// 1: Cloud Eye performs no aggregation and displays raw data.
	// 300: Cloud Eye aggregates data every 5 minutes.
	// 1200: Cloud Eye aggregates data every 20 minutes.
	// 3600: Cloud Eye aggregates data every 1 hour.
	// 14400: Cloud Eye aggregates data every 4 hours.
	// 86400: Cloud Eye aggregates data every 24 hours.
	Period string `json:"period" required:"true"`
	// Specifies the data rollup method. The following methods are supported:
	//
	// average: Cloud Eye calculates the average value of metric data within a rollup period.
	// max: Cloud Eye calculates the maximum value of metric data within a rollup period.
	// min: Cloud Eye calculates the minimum value of metric data within a rollup period.
	// sum: Cloud Eye calculates the sum of metric data within a rollup period.
	// variance: Cloud Eye calculates the variance value of metric data within a rollup period.
	// The value of filter does not affect the query result of raw data. (The period is 1.)
	Filter string `json:"filter" required:"true"`
}

type Metric struct {
	// Specifies the metric namespace. Its value must be in the service.item format and can contain 3 to 32 characters.
	// service and item each must be a string that starts with a letter and contains only letters, digits, and underscores (_).
	Namespace string `json:"namespace" required:"true"`

	// Specifies the metric name. Start with a letter.
	// Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	MetricName string `json:"metric_name" required:"true"`

	// Specifies the list of the metric dimensions.
	Dimensions []MetricsDimension `json:"dimensions" required:"true"`
}

func BatchListMetricData(client *golangsdk.ServiceClient, opts BatchListMetricDataOpts) ([]BatchMetricData, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /V1.0/{project_id}/batch-query-metric-data
	raw, err := client.Post(client.ServiceURL("batch-query-metric-data"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []BatchMetricData
	err = extract.IntoSlicePtr(raw.Body, &res, "metrics")

	return res, err
}

type BatchMetricData struct {
	// Specifies the metric namespace.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	Namespace string `json:"namespace"`
	// Specifies the metric name. Start with a letter. Enter 1 to 64 characters.
	MetricName string `json:"metric_name"`
	// Specifies the list of metric dimensions.
	Dimensions []MetricsDimension `json:"dimensions"`
	// Specifies the metric data list.
	// Since Cloud Eye rounds up the value of from based on the level of granularity for data query,
	// datapoints may contain more data points than expected.
	Datapoints []DatapointForBatchMetric `json:"datapoints"`
	// Specifies the metric unit.
	Unit string `json:"unit"`
}

type DatapointForBatchMetric struct {
	// Specifies the maximum value of metric data within a rollup period.
	Max float64 `json:"max"`
	// Specifies the minimum value of metric data within a rollup period.
	Min float64 `json:"min"`
	// Specifies the average value of metric data within a rollup period.
	Average float64 `json:"average"`
	// Specifies the sum of metric data within a rollup period.
	Sum float64 `json:"sum"`
	// Specifies the variance of metric data within a rollup period.
	Variance float64 `json:"variance"`
	// Specifies when the metric is collected. It is a UNIX timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
}
