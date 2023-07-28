package metricdata

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ShowMetricDataOpts struct {
	// Specifies the namespace of a service.
	Namespace string `q:"namespace"`
	// Specifies the metric name.
	MetricName string `q:"metric_name"`
	// Currently, a maximum of three metric dimensions are supported,
	// and the dimensions are numbered from 0 in the dim.{i}=key,value format.
	// The key cannot exceed 32 characters and the value cannot exceed 256 characters.
	// The following dimensions are only examples.
	// For details about whether multiple dimensions are supported,
	// see the dimension description in the monitoring indicator description of each service.
	// Single dimension: dim.0=instance_id,i-12345
	// Multiple dimensions: dim.0=instance_id,i-12345&dim.1=instance_name,i-1234
	Dim0 string `q:"dim.0"`
	Dim1 string `q:"dim.1"`
	Dim2 string `q:"dim.2"`
	// Specifies the data rollup method. The following methods are supported:
	//
	// average: Cloud Eye calculates the average value of metric data within a rollup period.
	// max: Cloud Eye calculates the maximum value of metric data within a rollup period.
	// min: Cloud Eye calculates the minimum value of metric data within a rollup period.
	// sum: Cloud Eye calculates the sum of metric data within a rollup period.
	// variance: Cloud Eye calculates the variance value of metric data within a rollup period.
	Filter string `q:"filter"`
	// Specifies how often Cloud Eye aggregates data.
	//
	// Possible values are:
	// 1: Cloud Eye performs no aggregation and displays raw data.
	// 300: Cloud Eye aggregates data every 5 minutes.
	// 1200: Cloud Eye aggregates data every 20 minutes.
	// 3600: Cloud Eye aggregates data every 1 hour.
	// 14400: Cloud Eye aggregates data every 4 hours.
	// 86400: Cloud Eye aggregates data every 24 hours.
	Period int `q:"period"`
	// Specifies the start time of the query.
	// The value is a UNIX timestamp and the unit is ms.
	// Set the value of from to at least one period earlier than the current time.
	// Rollup aggregates the raw data generated within a period to the start time of the period.
	// Therefore, if values of from and to are within a period,
	// the query result will be empty due to the rollup failure.
	// Take the 5-minute period as an example. If it is 10:35 now,
	// the raw data generated between 10:30 and 10:35 will be aggregated to 10:30.
	// Therefore, in this example, if the value of period is 5 minutes,
	// the value of from should be 10:30 or earlier.
	From string `q:"from"`
	// Specifies the end time of the query.
	To string `q:"to"`
}

func ShowMetricData(client *golangsdk.ServiceClient, opts ShowMetricDataOpts) (*MetricData, error) {
	var opts2 interface{} = &opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /V1.0/{project_id}/metric-data
	raw, err := client.Get(client.ServiceURL("metric-data")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res MetricData
	err = extract.Into(raw.Body, &res)

	return &res, err
}

type MetricData struct {
	// Specifies the metric data list. For details, see Table 4.
	// Since Cloud Eye rounds up the value of from based on the level of granularity for data query,
	// datapoints may contain more data points than expected.
	Datapoints []Datapoint `json:"datapoints"`
	// Specifies the metric ID. For example, if the monitoring metric of an ECS is CPU usage, metric_name is cpu_util.
	MetricName string `json:"metric_name"`
}

type Datapoint struct {
	// Specifies the maximum value of metric data within a rollup period.
	Max float64 `json:"max,omitempty"`
	// Specifies the minimum value of metric data within a rollup period.
	Min float64 `json:"min,omitempty"`
	// Specifies the average value of metric data within a rollup period.
	Average float64 `json:"average,omitempty"`
	// Specifies the sum of metric data within a rollup period.
	Sum float64 `json:"sum,omitempty"`
	// Specifies the variance of metric data within a rollup period.
	Variance float64 `json:"variance,omitempty"`
	// Specifies when the metric is collected. It is a UNIX timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
	// Specifies the metric unit.
	Unit string `json:"unit,omitempty"`
}
