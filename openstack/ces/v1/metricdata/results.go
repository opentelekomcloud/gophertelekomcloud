package metricdata

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

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
}

type BatchListMetricDataResult struct {
	golangsdk.Result
}

func (r BatchListMetricDataResult) Extract() ([]BatchMetricData, error) {
	var s struct {
		MetricDatas []BatchMetricData `json:"metrics"`
	}
	err := r.ExtractInto(&s)
	return s.MetricDatas, err
}

type CreateMetricDataResult struct {
	golangsdk.ErrResult
}

// --------------------------------------------------------------------------------------

type EventDataInfo struct {
	// Specifies the event type, for example, instance_host_info.
	Type string `json:"type"`
	// Specifies when the event is reported. It is a UNIX timestamp and the unit is ms.
	Timestamp int64 `json:"timestamp"`
	// Specifies the host configuration information.
	Value string `json:"value"`
}

type ShowEventDataResponse struct {
	Datapoints []EventDataInfo `json:"datapoints"`
}

type ShowEventDataResult struct {
	golangsdk.Result
}

func (r ShowEventDataResult) Extract() (*ShowEventDataResponse, error) {
	var s *ShowEventDataResponse
	err := r.ExtractInto(&s)
	return s, err
}

// ------------------------------------------------------------------------------------------

type ShowMetricDataResponse struct {
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

type ShowMetricDataResult struct {
	golangsdk.Result
}

func (r ShowMetricDataResult) Extract() (*ShowMetricDataResponse, error) {
	var s *ShowMetricDataResponse
	err := r.ExtractInto(&s)
	return s, err
}
