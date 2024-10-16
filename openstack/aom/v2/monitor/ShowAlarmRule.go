package monitor

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowAlarmRule(client *golangsdk.ServiceClient, alarmId int64) (*ShowAlarmRuleResponse, error) {
	// GET /v2/{project_id}/ams/alarms/{alarm_id}
	raw, err := client.Get(client.ServiceURL("ams", "alarms", strconv.FormatInt(alarmId, 10)), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowAlarmRuleResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowAlarmRuleResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Threshold rule list.
	Thresholds []Threshold `json:"thresholds"`
}

type Threshold struct {
	// Statistic.
	//
	// Value Range
	//
	// maximum, minimum, average, sum, or sampleCount.
	Statistic string `json:"statistic" required:"true"`
	// Namespace.
	//
	// Value Range
	//
	// 1-255 characters. Colons (:) are not allowed.
	Namespace string `json:"namespace" required:"true"`
	// Metric name.
	//
	// Value Range
	//
	// 1-255 characters.
	MetricName string `json:"metricName" required:"true"`
	// Statistical period.
	//
	// Value Range
	//
	// 20000, 60000, 300000, 900000, 1800000, 3600000, 14400000, or 86400000
	Period *int `json:"period" required:"true"`
	// Alarm severity.
	//
	// Value Range
	//
	// 4, 3, 2, or 1
	AlarmLevel *int `json:"alarmLevel,omitempty"`
	// Number of consecutive periods.
	//
	// Value Range
	//
	// 4, 3, or 2
	EvaluationPeriods *int `json:"evaluationPeriods" required:"true"`
	// Comparison operator.
	//
	// Value Range
	//
	// >, >=, <, or <=
	ComparisonOperator string `json:"comparisonOperator" required:"true"`
	// Threshold.
	//
	// Value Range
	//
	// Non-null value that can be converted to a value of the double type. Once converted, the value cannot be null, or be a positive or negative infinity.
	Threshold *int64 `json:"threshold" required:"true"`
	// Threshold name.
	//
	// Value Range
	//
	// 1-255 characters. Special characters are not allowed.
	AlarmName string `json:"alarmName" required:"true"`
	// Metric dimension.
	//
	// Value Range
	//
	// Non-null; number of dimensions < 100
	Dimensions []Dimension `json:"dimensions" required:"true"`
	// Metric unit.
	//
	// Value Range
	//
	// Number of characters < 32
	Unit string `json:"unit,omitempty"`
	// Whether to enable the alarm function.
	ActionEnabled *bool `json:"actionEnabled,omitempty"`
	// Alarm action.
	//
	// Value Range
	//
	// Size <= 5
	AlarmActions []string `json:"alarmActions,omitempty"`
	// Suggestion.
	//
	// Value Range
	//
	// Number of characters < 255
	AlarmAdvice string `json:"alarmAdvice,omitempty"`
	// Threshold rule description.
	//
	// Value Range
	//
	// Number of characters < 255
	AlarmDescription string `json:"alarmDescription,omitempty"`
	// Action to be taken when data is insufficient.
	//
	// Value Range
	//
	// Size <= 5
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty"`
	// Recovery action.
	//
	// Value Range
	//
	// Size <= 5
	OkActions []string `json:"okActions,omitempty"`
	// ?
	StateValue string `json:"stateValue,omitempty"`
	// ?
	StateReason string `json:"stateReason,omitempty"`
	// ?
	StateUpdatedTimestamp *int64 `json:"stateUpdatedTimestamp,omitempty"`
}
