package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ListAlarmsRequest struct {
	// The value ranges from 1 to 100, and is 100 by default.
	// This parameter is used to limit the number of query results.
	Limit int `q:"limit"`
	// Specifies the result sorting method, which is sorted by timestamp.
	// The default value is desc.
	// asc: The query results are displayed in the ascending order.
	// desc: The query results are displayed in the descending order.
	Order string `q:"order"`
	// Specifies the first queried alarm to be displayed on a page.
	Start string `q:"start"`
}

func ListAlarms(client *golangsdk.ServiceClient, req ListAlarmsRequest) (*ListAlarmsResponse, error) {
	url := alarmsURL(client)
	query, err := golangsdk.BuildQueryString(req)
	if err != nil {
		return nil, err
	}

	raw := golangsdk.Result{}
	url += query.String()

	_, err = client.Get(url, &raw.Body, nil)
	if err != nil {
		return nil, err
	}

	var s ListAlarmsResponse
	err = raw.ExtractInto(&s)
	return &s, err
}

// ------------------------------------------------------------------------------------------------

func ShowAlarm(client *golangsdk.ServiceClient, id string) ([]MetricAlarms, error) {
	raw := golangsdk.Result{}
	_, err := client.Get(alarmIdURL(client, id), &raw.Body, nil)
	if err != nil {
		return nil, err
	}

	var res []MetricAlarms
	err = raw.ExtractIntoSlicePtr(&res, "metric_alarms")
	return res, err
}

// ------------------------------------------------------------------------------------------------

type ModifyAlarmActionRequest struct {
	// Specifies whether the alarm rule is enabled.
	AlarmEnabled bool `json:"alarm_enabled"`
}

func UpdateAlarmAction(client *golangsdk.ServiceClient, id string, req ModifyAlarmActionRequest) (err error) {
	reqBody, err := golangsdk.BuildRequestBody(req, "")
	if err != nil {
		return
	}

	_, err = client.Put(alarmActionURL(client, id), reqBody, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// ------------------------------------------------------------------------------------------------

func DeleteAlarm(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(alarmIdURL(client, id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// ------------------------------------------------------------------------------------------------

type CreateAlarmRequest struct {
	// Specifies the alarm rule name.
	// Enter 1 to 128 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed.
	AlarmName string `json:"alarm_name"`
	// Provides supplementary information about the alarm rule. Enter 0 to 256 characters.
	AlarmDescription string `json:"alarm_description,omitempty"`
	// Specifies the alarm metric.
	Metric MetricForAlarm `json:"metric"`
	// Specifies the alarm triggering condition.
	Condition Condition `json:"condition"`
	// Specifies whether to enable the alarm.
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`
	// Specifies whether to enable the action to be triggered by an alarm. The default value is true.
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// Specifies the alarm severity. Possible values are 1, 2 (default), 3 and 4,
	// indicating critical, major, minor, and informational, respectively.
	AlarmLevel int `json:"alarm_level,omitempty"`
	// Specifies the action to be triggered by an alarm.
	AlarmActions []AlarmActions `json:"alarm_actions,omitempty"`
	// Specifies the action to be triggered after the alarm is cleared.
	OkActions []AlarmActions `json:"ok_actions,omitempty"`
}

type MetricForAlarm struct {
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	// service and item each must start with a letter and contain only letters, digits, and underscores (_).
	Namespace string `json:"namespace"`
	// Specifies the metric name.
	// Start with a letter. Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	MetricName string `json:"metric_name"`
	// Specifies the list of metric dimensions.
	Dimensions []MetricsDimension `json:"dimensions,omitempty"`
	// Specifies the resource group ID selected during the alarm rule creation.
	ResourceGroupId string `json:"resource_group_id,omitempty"`
}

func CreateAlarm(client *golangsdk.ServiceClient, req CreateAlarmRequest) (string, error) {
	reqBody, err := golangsdk.BuildRequestBody(req, "")
	if err != nil {
		return "", err
	}

	raw := golangsdk.Result{}
	_, err = client.Post(alarmsURL(client), reqBody, &raw.Body, nil)
	if err != nil {
		return "", err
	}

	var s struct {
		AlarmId string `json:"alarm_id"`
	}
	err = raw.ExtractInto(&s)
	return s.AlarmId, err
}
