package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for creating a custom alarm template.
type CreateOpts struct {
	// Specifies the alarm template name.
	// It must start with a letter and can contain letters, digits, underscores (_),
	// hyphens (-), parentheses, and periods (.).
	// Enter 1 to 128 characters.
	TemplateName string `json:"template_name" required:"true"`
	// Specifies the alarm template type.
	// 0: metric alarm template
	// 2: event alarm template
	TemplateType int `json:"template_type,omitempty"`
	// Provides supplementary information about the alarm template.
	// Enter 0 to 256 characters.
	TemplateDescription string `json:"template_description,omitempty"`
	// Specifies the alarm policies. A maximum of 50 policies are supported.
	Policies []Policy `json:"policies" required:"true"`
}

// Policy represents an alarm policy in the template.
type Policy struct {
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	Namespace string `json:"namespace" required:"true"`
	// Specifies the resource dimension.
	// The value can contain a maximum of 32 characters.
	// Leave this parameter blank for an event alarm template.
	DimensionName string `json:"dimension_name,omitempty"`
	// Specifies the metric name of a resource.
	// It must start with a letter and can contain 1 to 96 characters.
	MetricName string `json:"metric_name" required:"true"`
	// Specifies the aggregation period in seconds.
	// Possible values: 0, 1, 300, 1200, 3600, 14400, 86400
	Period int `json:"period" required:"true"`
	// Specifies the data aggregation method.
	// Possible values: average, max, min, sum, variance
	Filter string `json:"filter" required:"true"`
	// Specifies the comparison operator.
	// Possible values: >, <, >=, <=, =, !=
	ComparisonOperator string `json:"comparison_operator" required:"true"`
	// Specifies the alarm threshold.
	Value float64 `json:"value,omitempty"`
	// Specifies the data unit.
	// Enter 0 to 32 characters.
	Unit string `json:"unit,omitempty"`
	// Specifies the number of consecutive times that the alarm condition is met.
	// For event alarms: 1-180
	// For metric alarms: 1, 2, 3, 4, 5, 10, 15, 30, 60, 90, 120, or 180
	Count int `json:"count" required:"true"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	// Default value: 2
	AlarmLevel int `json:"alarm_level,omitempty"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	// Possible values: 0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400
	SuppressDuration int `json:"suppress_duration" required:"true"`
}

// Create creates a new custom alarm template.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/alarm-templates
	raw, err := client.Post(client.ServiceURL("alarm-templates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		TemplateId string `json:"template_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.TemplateId, err
}
