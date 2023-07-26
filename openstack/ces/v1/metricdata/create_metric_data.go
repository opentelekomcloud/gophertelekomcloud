package metricdata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

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

func CreateMetricData(client *golangsdk.ServiceClient, items []MetricDataItem) error {
	b := make([]map[string]any, len(items))

	for i, opt := range items {
		opt, err := golangsdk.BuildRequestBody(opt, "")
		if err != nil {
			return err
		}
		b[i] = opt
	}

	// POST /V1.0/{project_id}/metric-data
	_, err := client.Post(client.ServiceURL("metric-data"), &b, nil, nil)
	return err
}
