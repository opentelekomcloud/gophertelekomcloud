package alarms

type ListAlarmsResponse struct {
	MetricAlarms []MetricAlarms `json:"metric_alarms,omitempty"`
	MetaData     MetaData       `json:"meta_data,omitempty"`
}

type MetaData struct {
	Count  int    `json:"count"`
	Total  int    `json:"total"`
	Marker string `json:"marker"`
}

type MetricAlarms struct {
	// Specifies the alarm rule name.
	AlarmName string `json:"alarm_name"`
	// Provides supplementary information about the alarm rule.
	AlarmDescription string `json:"alarm_description,omitempty"`
	// Specifies the alarm rule type.
	AlarmType string `json:"alarm_type"`
	// Specifies the alarm metric.
	Metric MetricForAlarm `json:"metric"`
	// Specifies the alarm triggering condition.
	Condition Condition `json:"condition"`
	// Specifies whether to enable the alarm rule.
	AlarmEnabled bool `json:"alarm_enabled,omitempty"`
	// Specifies the alarm severity. Possible values are 1, 2, 3 and 4, indicating critical, major, minor, and informational, respectively.
	AlarmLevel int `json:"alarm_level,omitempty"`
	// Specifies whether to enable the action to be triggered by an alarm.
	AlarmActionEnabled bool `json:"alarm_action_enabled,omitempty"`
	// Specifies the action to be triggered by an alarm.
	AlarmActions []AlarmActions `json:"alarm_actions,omitempty"`
	//  Specifies the action to be triggered after the alarm is cleared.
	OkActions []AlarmActions `json:"ok_actions,omitempty"`
	// Specifies the alarm rule ID.
	AlarmId string `json:"alarm_id"`
	// Specifies when the alarm status changed. The value is a UNIX timestamp and the unit is ms.
	UpdateTime int64 `json:"update_time"`
	// Specifies the alarm status. The value can be:
	//
	// ok: The alarm status is normal.
	// alarm: An alarm is generated.
	// insufficient_data: The required data is insufficient.
	AlarmState string `json:"alarm_state"`
}

type MetricsDimension struct {
	// Specifies the dimension. For example, the ECS dimension is instance_id.
	Name string `json:"name"`
	// Specifies the dimension value, for example, an ECS ID.
	Value string `json:"value"`
}

type Condition struct {
	// Specifies the operator of alarm thresholds. Possible values are >, =, <, >=, and <=.
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the number of consecutive occurrence times that the alarm policy was met.
	// The value ranges from 1 to 5.
	Count int `json:"count"`
	// Specifies the data rollup method. The following methods are supported:
	//
	// average: Cloud Eye calculates the average value of metric data within a rollup period.
	// max: Cloud Eye calculates the maximum value of metric data within a rollup period.
	// min: Cloud Eye calculates the minimum value of metric data within a rollup period.
	// sum: Cloud Eye calculates the sum of metric data within a rollup period.
	// variance: Cloud Eye calculates the variance value of metric data within a rollup period.
	Filter string `json:"filter"`
	// Specifies the interval (seconds) for checking whether the configured alarm rules are met.
	Period int `json:"period"`
	// Specifies the data unit. Enter up to 32 characters.
	Unit string `json:"unit,omitempty"`
	// Specifies the alarm threshold. The value ranges from 0 to Number. MAX_VALUE (1.7976931348623157e+108).
	//
	// For detailed thresholds, see the value range of each metric in the appendix.
	// For example, you can set ECS cpu_util in Services Interconnected with Cloud Eye to 80.
	Value float64 `json:"value"`
	// Specifies the interval for triggering an alarm if the alarm persists.
	SuppressDuration int `json:"suppress_duration"`
}

type AlarmActions struct {
	// Specifies the alarm notification type.
	// notification: indicates that a notification will be sent.
	// autoscaling: indicates that a scaling action will be triggered.
	Type string `json:"type"`
	// Specifies the list of objects to be notified if the alarm status changes.
	// You can configure up to 5 object IDs. You can obtain the topicUrn value from SMN in the following format:
	// urn:smn:([a-z]|[A-Z]|[0-9]|\-){1,32}:([a-z]|[A-Z]|[0-9]){32}:([a-z]|[A-Z]|[0-9]|\-|\_){1,256}.
	// If you set type to notification, you must specify notificationList.
	// If you set type to autoscaling, you must set notificationList to [].
	NotificationList []string `json:"notificationList"`
}
