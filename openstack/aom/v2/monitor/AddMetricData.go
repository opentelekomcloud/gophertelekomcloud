package monitor

import "github.com/opentelekomcloud/gophertelekomcloud"

type AddMetricDataOpts struct {
	// Metric data.
	Metric MetricAdd `json:"metric" required:"true"`
	Values []Value   `json:"values" required:"true"`
	// Data collection time, which must meet the following requirement:
	//
	// Current UTC time - Data collection time <= 24 hours, or Data collection time - Current UTC time <= 30 minutes
	//
	// Value Range
	//
	// UNIX timestamp, in ms.
	CollectTime *int64 `json:"collect_time" required:"true"`
}

type MetricAdd struct {
	// Metric namespace.
	//
	// Value Range
	//
	// Namespace, which must be in the format of service.item. The value must be 3 to 32 characters starting with a letter. Only letters, digits, and underscores (_) are allowed. In addition, service cannot start with PAAS or SYS.
	Namespace string `json:"namespace" required:"true"`
	// List of metric dimensions.
	//
	// Value Range
	//
	// Each dimension is a JSON object, and its structure is as follows:
	//
	// dimension.name: 1-32 characters.
	//
	// dimension.value: 1-64 characters.
	Dimensions []Dimension `json:"dimensions" required:"true"`
}

type Value struct {
	// Data unit.
	Unit string `json:"unit,omitempty"`
	// Metric name.
	MetricName string `json:"metric_name" required:"true"`
	// Data type.
	//
	// Value Range
	//
	// "int" or "float"
	Type string `json:"type,omitempty"`
	// Metric value.
	//
	// Value Range
	//
	// Valid numeral type.
	Value *float32 `json:"value" required:"true"`
}

func AddMetricData(client *golangsdk.ServiceClient) {
	// POST /v2/{project_id}/ams/report/metricdata
}

type AddMetricDataResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
}
