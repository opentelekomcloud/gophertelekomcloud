package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying policies in an alarm rule.
type ListOpts struct {
	// Specifies the pagination offset. Default: 0
	Offset int `q:"offset"`
	// Specifies the number of records on each page. Default: 10, Max: 100
	Limit int `q:"limit"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the list of alarm policies.
	Policies []PolicyResp `json:"policies"`
	// Specifies the total number of policies.
	Count int `json:"count"`
}

// PolicyResp represents an alarm policy in the response.
type PolicyResp struct {
	// Specifies the metric name.
	MetricName string `json:"metric_name"`
	// Specifies extra metric information.
	ExtraInfo *MetricExtraInfo `json:"extra_info,omitempty"`
	// Specifies the rollup period in seconds.
	Period int `json:"period"`
	// Specifies the data rollup method.
	Filter string `json:"filter"`
	// Specifies the comparison operator.
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the alarm threshold.
	Value float64 `json:"value"`
	// Specifies the data unit.
	Unit string `json:"unit"`
	// Specifies the policy type. "auto" for dynamic threshold.
	Type string `json:"type"`
	// Specifies the number of consecutive times.
	Count int `json:"count"`
	// Specifies the interval for triggering an alarm if the alarm persists.
	SuppressDuration int `json:"suppress_duration"`
	// Specifies the alarm severity.
	Level int `json:"level"`
}

// MetricExtraInfo represents extra metric information.
type MetricExtraInfo struct {
	// Specifies the original metric name.
	OriginMetricName string `json:"origin_metric_name,omitempty"`
	// Specifies the metric prefix.
	MetricPrefix string `json:"metric_prefix,omitempty"`
	// Specifies the custom process name.
	CustomProcName string `json:"custom_proc_name,omitempty"`
	// Specifies the metric type.
	MetricType string `json:"metric_type,omitempty"`
}

// List returns a list of policies in an alarm rule.
func List(client *golangsdk.ServiceClient, alarmId string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("alarms", alarmId, "policies").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarms/{alarm_id}/policies
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
