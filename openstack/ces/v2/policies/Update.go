package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Policy represents an alarm policy.
type Policy struct {
	// Specifies the metric name.
	MetricName string `json:"metric_name" required:"true"`
	// Specifies the rollup period in seconds.
	// Possible values: 0, 1, 300, 1200, 3600, 14400, 86400
	Period int `json:"period" required:"true"`
	// Specifies the data rollup method.
	// Possible values: average, max, min, sum, variance
	Filter string `json:"filter" required:"true"`
	// Specifies the comparison operator.
	// Possible values: >, <, >=, <=, =, !=
	ComparisonOperator string `json:"comparison_operator" required:"true"`
	// Specifies the alarm threshold.
	Value float64 `json:"value,omitempty"`
	// Specifies the data unit.
	Unit string `json:"unit,omitempty"`
	// Specifies the number of consecutive times that the alarm condition is met.
	Count int `json:"count" required:"true"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	// Possible values: 0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400
	SuppressDuration int `json:"suppress_duration,omitempty"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	Level int `json:"level,omitempty"`
}

// UpdateOpts contains the options for modifying policies in an alarm rule.
type UpdateOpts struct {
	// Specifies the list of alarm policies.
	// Array length: 1-50 items.
	Policies []Policy `json:"policies" required:"true"`
}

// UpdateResponse contains the response from the Update request.
type UpdateResponse struct {
	// Specifies the list of updated alarm policies.
	Policies []Policy `json:"policies"`
}

// Update modifies policies in an alarm rule.
func Update(client *golangsdk.ServiceClient, alarmId string, opts UpdateOpts) (*UpdateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/alarms/{alarm_id}/policies
	raw, err := client.Put(client.ServiceURL("alarms", alarmId, "policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
