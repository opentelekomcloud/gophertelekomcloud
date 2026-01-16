package records

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying alarm records.
type ListOpts struct {
	// Specifies the alarm rule ID.
	AlarmId string `q:"alarm_id"`
	// Specifies the alarm rule name. Enter 1 to 128 characters.
	Name string `q:"name"`
	// Specifies the alarm record status.
	// Possible values: ok, alarm, invalid
	Status string `q:"status"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	Level int `q:"level,omitempty"`
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	Namespace string `q:"namespace"`
	// Specifies the resource ID.
	// For multiple dimensions, the format is dim1=value1,dim2=value2.
	ResourceId string `q:"resource_id"`
	// Specifies the start time of the query.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	From string `q:"from"`
	// Specifies the end time of the query.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	To string `q:"to"`
	// Specifies the pagination offset. Default: 0, Max: 999
	Offset int `q:"offset,omitempty"`
	// Specifies the number of records on each page. Default: 100, Range: 1-100
	Limit int `q:"limit,omitempty"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the list of alarm records.
	AlarmHistories []AlarmHistoryItem `json:"alarm_histories"`
	// Specifies the total number of alarm records.
	Count int `json:"count"`
}

// AlarmHistoryItem represents an alarm record.
type AlarmHistoryItem struct {
	// Specifies the alarm record ID.
	RecordId string `json:"record_id"`
	// Specifies the alarm rule ID.
	AlarmId string `json:"alarm_id"`
	// Specifies the alarm rule name.
	Name string `json:"name"`
	// Specifies the alarm record status.
	// Possible values: ok, alarm, invalid
	Status string `json:"status"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	Level int `json:"level"`
	// Specifies the alarm rule type.
	// Possible values: ALL_INSTANCE, RESOURCE_GROUP, MULTI_INSTANCE, EVENT.SYS, EVENT.CUSTOM
	Type string `json:"type"`
	// Specifies whether to send notifications.
	ActionEnabled bool `json:"action_enabled"`
	// Specifies the time when the alarm record was generated (UTC).
	BeginTime string `json:"begin_time"`
	// Specifies the time when the alarm record ended (UTC).
	EndTime string `json:"end_time"`
	// Specifies the time when the first alarm was generated (UTC).
	FirstAlarmTime string `json:"first_alarm_time"`
	// Specifies the time when the last alarm was generated (UTC).
	LastAlarmTime string `json:"last_alarm_time"`
	// Specifies the time when the alarm was cleared (UTC).
	AlarmRecoveryTime string `json:"alarm_recovery_time"`
	// Specifies the metric information.
	Metric Metric `json:"metric"`
	// Specifies the alarm triggering condition.
	Condition AlarmCondition `json:"condition"`
	// Specifies additional information about the alarm.
	AdditionalInfo AdditionalInfo `json:"additional_info"`
	// Specifies the action triggered by an alarm.
	AlarmActions []Notification `json:"alarm_actions"`
	// Specifies the action triggered after an alarm is cleared.
	OkActions []Notification `json:"ok_actions"`
	// Specifies the metric monitoring data.
	DataPoints []DataPointInfo `json:"data_points"`
}

// Metric represents metric information.
type Metric struct {
	// Specifies the namespace of a service.
	Namespace string `json:"namespace"`
	// Specifies the metric name.
	MetricName string `json:"metric_name"`
	// Specifies the list of metric dimensions.
	Dimensions []Dimension `json:"dimensions"`
}

// Dimension represents a metric dimension.
type Dimension struct {
	// Specifies the dimension name.
	Name string `json:"name"`
	// Specifies the dimension value.
	Value string `json:"value"`
}

// AlarmCondition represents the alarm triggering condition.
type AlarmCondition struct {
	// Specifies the rollup period in seconds.
	Period int `json:"period"`
	// Specifies the data rollup method.
	// Possible values: average, max, min, sum, variance
	Filter string `json:"filter"`
	// Specifies the comparison operator.
	// Possible values: >, <, >=, <=, =, !=
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the alarm threshold.
	Value float64 `json:"value"`
	// Specifies the data unit.
	Unit string `json:"unit"`
	// Specifies the number of consecutive times that the alarm condition is met.
	Count int `json:"count"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	SuppressDuration int `json:"suppress_duration"`
}

// AdditionalInfo represents additional information about the alarm.
type AdditionalInfo struct {
	// Specifies the resource ID.
	ResourceId string `json:"resource_id"`
	// Specifies the resource name.
	ResourceName string `json:"resource_name"`
	// Specifies the event ID.
	EventId string `json:"event_id"`
}

// Notification represents an alarm notification.
type Notification struct {
	// Specifies the notification type.
	// Possible values: notification, autoscaling, groupwatch, ecsRecovery, contact, contactGroup, iecAction
	Type string `json:"type"`
	// Specifies the list of objects to be notified.
	NotificationList []string `json:"notification_list"`
}

// DataPointInfo represents a monitoring data point.
type DataPointInfo struct {
	// Specifies the time when the monitoring data was collected (UTC).
	Time string `json:"time"`
	// Specifies the monitoring value.
	Value float64 `json:"value"`
}

// List returns a list of alarm records.
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("alarm-histories").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarm-histories
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
