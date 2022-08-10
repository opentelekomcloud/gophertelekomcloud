package metricdata

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchListMetricDataRequest struct {
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

func BatchListMetricData(client *golangsdk.ServiceClient, opts BatchListMetricDataRequest) ([]BatchMetricData, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(batchQueryMetricDataURL(client), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []BatchMetricData
	err = extract.IntoSlicePtr(raw.Body, &res, "metrics")

	return res, err
}

// ----------------------------------------------------------------------------

type MetricDataItem struct {
	// Specifies the metric data.
	Metric MetricInfo `json:"metric" required:"true"`
	// Specifies the data validity period.
	// The unit is second. The value range is 0–604,800 seconds.
	// If the validity period expires, the data will be automatically deleted.
	Ttl int `json:"ttl" required:"true"`
	// Specifies when the data was collected.
	// The time is UNIX timestamp (ms) format.
	CollectTime int64 `json:"collect_time" required:"true"`
	// Specifies the monitoring metric data to be added.
	// The value can be an integer or a floating point number.
	Value float64 `json:"value" required:"true"`
	// Specifies the data unit.
	// Enter a maximum of 32 characters.
	Unit string `json:"unit,omitempty"`
	// Specifies the enumerated type.
	// Possible values:
	// int
	// float
	Type string `json:"type,omitempty"`
}

type MetricInfo struct {
	// Specifies the metric dimension. A maximum of three dimensions are supported.
	Dimensions []MetricsDimension `json:"dimensions" required:"true"`
	// Specifies the metric ID. For example, if the monitoring metric of an ECS is CPU usage, metric_name is cpu_util.
	MetricName string `json:"metric_name" required:"true"`
	// Query the namespace of a service.
	Namespace string `json:"namespace" required:"true"`
}

type MetricsDimension struct {
	// Specifies the dimension. For example, the ECS dimension is instance_id.
	// For details about the dimension of each service, see the key column in Services Interconnected with Cloud Eye.
	// Start with a letter. Enter 1 to 32 characters.
	// Only letters, digits, underscores (_), and hyphens (-) are allowed.
	Name string `json:"name,omitempty"`
	// Specifies the dimension value, for example, an ECS ID.
	// Start with a letter or a digit. Enter 1 to 256 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed.
	Value string `json:"value,omitempty"`
}

func CreateMetricData(client *golangsdk.ServiceClient, opts []MetricDataItem) error {
	b := make([]map[string]interface{}, len(opts))

	for i, opt := range opts {
		opt, err := golangsdk.BuildRequestBody(opt, "")
		if err != nil {
			return err
		}
		b[i] = opt
	}

	_, err := client.Post(metricDataURL(client), &b, nil, nil)
	return err
}

// ----------------------------------------------------------------------------

type ShowEventDataRequest struct {
	// Query the namespace of a service.
	Namespace string `q:"namespace"`
	// Specifies the dimension. For example, the ECS dimension is instance_id.
	// For details about the dimensions corresponding to the monitoring metrics of each service,
	// see the monitoring metrics description of the corresponding service in Services Interconnected with Cloud Eye.
	//
	// Specifies the dimension. A maximum of three dimensions are supported,
	// and the dimensions are numbered from 0 in dim.{i}=key,value format.
	// The key cannot exceed 32 characters and the value cannot exceed 256 characters.
	Dim string `q:"dim.0"`
	// Specifies the event type.
	Type string `q:"type"`
	// Specifies the start time of the query.
	From string `q:"from"`
	// Specifies the end time of the query.
	To string `q:"to"`
}

func ShowEventData(client *golangsdk.ServiceClient, opts ShowEventDataRequest) ([]EventDataInfo, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(eventDataURL(client)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []EventDataInfo
	err = extract.IntoSlicePtr(raw.Body, &res, "datapoints")

	return res, err
}

// ----------------------------------------------------------------------------

type ShowMetricDataRequest struct {
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
	Dim string `q:"dim.0"`
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

func ShowMetricData(client *golangsdk.ServiceClient, opts ShowMetricDataRequest) (*ShowMetricDataResponse, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(metricDataURL(client)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowMetricDataResponse
	err = extract.Into(raw.Body, &res)

	return &res, err
}
