package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying alarm rules.
type ListOpts struct {
	// Specifies the alarm rule ID.
	AlarmId string `q:"alarm_id"`
	// Specifies the alarm rule name.
	Name string `q:"name"`
	// Specifies the namespace of a service.
	Namespace string `q:"namespace"`
	// Specifies the resource ID.
	ResourceId string `q:"resource_id"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId string `q:"enterprise_project_id"`
	// Specifies the pagination offset. Default: 0
	Offset int `q:"offset"`
	// Specifies the number of records on each page. Default: 10, Max: 100
	Limit int `q:"limit"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the list of alarm rules.
	Alarms []Alarm `json:"alarms"`
	// Specifies the total number of alarm rules.
	Count int `json:"count"`
}

// Alarm represents an alarm rule.
type Alarm struct {
	// Specifies the alarm rule ID.
	AlarmId string `json:"alarm_id"`
	// Specifies the alarm rule name.
	Name string `json:"name"`
	// Specifies the alarm rule description.
	Description string `json:"description"`
	// Specifies the namespace of the service.
	Namespace string `json:"namespace"`
	// Specifies the alarm policies.
	Policies []PolicyResp `json:"policies"`
	// Specifies the resource list.
	Resources []ResourceResp `json:"resources"`
	// Specifies the alarm rule type.
	Type string `json:"type"`
	// Specifies whether to generate alarms.
	Enabled bool `json:"enabled"`
	// Specifies whether notification is enabled.
	NotificationEnabled bool `json:"notification_enabled"`
	// Specifies the action triggered by an alarm.
	AlarmNotifications []Notification `json:"alarm_notifications"`
	// Specifies the action triggered after an alarm is cleared.
	OkNotifications []Notification `json:"ok_notifications"`
	// Specifies the time when the alarm notification was enabled.
	NotificationBeginTime string `json:"notification_begin_time"`
	// Specifies the time when the alarm notification was disabled.
	NotificationEndTime string `json:"notification_end_time"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// Specifies the alarm template ID.
	AlarmTemplateId string `json:"alarm_template_id"`
}

// PolicyResp represents an alarm policy in the response.
type PolicyResp struct {
	// Specifies the metric name.
	MetricName string `json:"metric_name"`
	// Specifies the rollup period in seconds.
	Period int `json:"period"`
	// Specifies the data rollup method.
	Filter string `json:"filter"`
	// Specifies the comparison operator.
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the alarm threshold.
	Value float64 `json:"value"`
	// Specifies the hierarchical alarm threshold.
	HierarchicalValue map[string]interface{} `json:"hierarchical_value"`
	// Specifies the data unit.
	Unit string `json:"unit"`
	// Specifies the number of consecutive times.
	Count int `json:"count"`
	// Specifies the interval for triggering an alarm if the alarm persists.
	SuppressDuration int `json:"suppress_duration"`
	// Specifies the alarm severity.
	Level int `json:"level"`
}

// ResourceResp represents a resource in the response.
type ResourceResp struct {
	// Specifies the resource group ID.
	ResourceGroupId string `json:"resource_group_id"`
	// Specifies the resource group name.
	ResourceGroupName string `json:"resource_group_name"`
	// Specifies the metric dimensions.
	Dimensions []Dimension `json:"dimensions"`
}

// List returns a list of alarm rules.
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("alarms").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarms
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
