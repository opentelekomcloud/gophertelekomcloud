package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// TemplateDetail represents the detailed information of an alarm template.
type TemplateDetail struct {
	// Specifies the alarm template ID.
	TemplateId string `json:"template_id"`
	// Specifies the alarm template name.
	TemplateName string `json:"template_name"`
	// Specifies the alarm template type.
	// Possible values: system, custom
	TemplateType string `json:"template_type"`
	// Specifies the time when the alarm template was created.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	CreateTime string `json:"create_time"`
	// Provides supplementary information about the alarm template.
	TemplateDescription string `json:"template_description"`
	// Specifies the alarm policies.
	Policies []PolicyResp `json:"policies"`
}

// PolicyResp represents an alarm policy in the response.
type PolicyResp struct {
	// Specifies the namespace of a service.
	Namespace string `json:"namespace"`
	// Specifies the resource dimension.
	DimensionName string `json:"dimension_name"`
	// Specifies the metric name of a resource.
	MetricName string `json:"metric_name"`
	// Specifies the aggregation period in seconds.
	Period int `json:"period"`
	// Specifies the data aggregation method.
	Filter string `json:"filter"`
	// Specifies the comparison operator.
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the alarm threshold.
	Value float64 `json:"value"`
	// Specifies the data unit.
	Unit string `json:"unit"`
	// Specifies the number of consecutive times that the alarm condition is met.
	Count int `json:"count"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	AlarmLevel int `json:"alarm_level"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	SuppressDuration int `json:"suppress_duration"`
}

// Get retrieves details of an alarm template.
func Get(client *golangsdk.ServiceClient, templateId string) (*TemplateDetail, error) {
	// GET /v2/{project_id}/alarm-templates/{template_id}
	raw, err := client.Get(client.ServiceURL("alarm-templates", templateId), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res TemplateDetail
	err = extract.Into(raw.Body, &res)
	return &res, err
}
