package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListLogsOpts struct {
	// Specifies the AS policy ID.
	ScalingPolicyId string
	// Specifies the ID of an AS policy execution log.
	LogId string `q:"log_id,omitempty"`
	// Specifies the scaling resource type.
	// AS group: SCALING_GROUP
	// Bandwidth: BANDWIDTH
	ScalingResourceType string `q:"scaling_resource_type,omitempty"`
	// Specifies the scaling resource ID.
	ScalingResourceId string `q:"scaling_resource_id,omitempty"`
	// Specifies the AS policy execution type.
	// SCHEDULED: automatically triggered at a specified time point
	// RECURRENCE: automatically triggered at a specified time period
	// ALARM: alarm-triggered
	// MANUAL: manually triggered
	ExecuteType string `q:"execute_type,omitempty"`
	// Specifies the start time that complies with UTC for querying AS policy execution logs.
	// The format of the start time is yyyy-MM-ddThh:mm:ssZ.
	StartTime string `q:"start_time,omitempty"`
	// Specifies the end time that complies with UTC for querying AS policy execution logs.
	// The format of the end time is yyyy-MM-ddThh:mm:ssZ.
	EndTime string `q:"end_time,omitempty"`
	// Specifies the start line number. The default value is 0. The minimum parameter value is 0.
	StartNumber int32 `q:"start_number,omitempty"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int32 `q:"limit,omitempty"`
}

func ListScalingPolicyExecuteLogs(client *golangsdk.ServiceClient, opts ListLogsOpts) (*ListScalingPolicyExecuteLogsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_policy_execute_log", opts.ScalingPolicyId).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /autoscaling-api/v1/{project_id}/scaling_policy_execute_log/{scaling_policy_id}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingPolicyExecuteLogsResponse
	err = extract.IntoStructPtr(raw.Body, &res, "scaling_policy")
	return &res, err
}

type ListScalingPolicyExecuteLogsResponse struct {
	TotalNumber             int32                         `json:"total_number,omitempty"`
	StartNumber             int32                         `json:"start_number,omitempty"`
	Limit                   int32                         `json:"limit,omitempty"`
	ScalingPolicyExecuteLog []ScalingPolicyExecuteLogList `json:"scaling_policy_execute_log,omitempty"`
}

type ScalingPolicyExecuteLogList struct {
	// Specifies the AS policy execution status.
	// SUCCESS: The AS policy has been executed.
	// FAIL: Executing the AS policy failed.
	// EXECUTING: The AS policy is being executed.
	Status string `json:"status,omitempty"`
	// Specifies the AS policy execution failure.
	FailedReason string `json:"failed_reason,omitempty"`
	// Specifies the AS policy execution type.
	// SCHEDULED: automatically triggered at a specified time point
	// RECURRENCE: automatically triggered at a specified time period
	// ALARM: alarm-triggered
	// MANUAL: manually triggered
	ExecuteType string `json:"execute_type,omitempty"`
	// Specifies the time when an AS policy was executed. The time format complies with UTC.
	ExecuteTime string `json:"execute_time,omitempty"`
	// Specifies the ID of an AS policy execution log.
	Id string `json:"id,omitempty"`
	// Specifies the project ID.
	TenantId string `json:"tenant_id,omitempty"`
	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id,omitempty"`
	// Specifies the scaling resource type.
	// AS group: SCALING_GROUP
	// Bandwidth: BANDWIDTH
	ScalingResourceType string `json:"scaling_resource_type,omitempty"`
	// Specifies the scaling resource ID.
	ScalingResourceId string `json:"scaling_resource_id,omitempty"`
	// Specifies the source value.
	OldValue string `json:"old_value,omitempty"`
	// Specifies the target value.
	DesireValue string `json:"desire_value,omitempty"`
	// Specifies the operation restrictions.
	// If scaling_resource_type is set to BANDWIDTH and operation is not SET, this parameter takes effect and the unit is Mbit/s.
	// In this case:
	// If operation is set to ADD, this parameter indicates the maximum bandwidth allowed.
	// If operation is set to REDUCE, this parameter indicates the minimum bandwidth allowed.
	LimitValue string `json:"limit_value,omitempty"`
	// Specifies the AS policy execution type.
	// ADD: indicates adding instances.
	// REMOVE: indicates reducing instances.
	// SET: indicates setting the number of instances to a specified value.
	Type string `json:"type,omitempty"`
	// Specifies the tasks contained in a scaling action based on an AS policy.
	JobRecords []JobRecords `json:"job_records,omitempty"`

	MetaData EipMetaData `json:"meta_data,omitempty"`
}

type JobRecords struct {
	// Specifies the task name.
	JobName string `json:"job_name,omitempty"`
	// Specifies the record type.
	// API: API calling type
	// MEG: message type
	RecordType string `json:"record_type,omitempty"`
	// Specifies the record time.
	RecordTime string `json:"record_time,omitempty"`
	// Specifies the request body. This parameter is valid only if record_type is set to API.
	Request string `json:"request,omitempty"`
	// Specifies the response body. This parameter is valid only if record_type is set to API.
	Response string `json:"response,omitempty"`
	// Specifies the returned code. This parameter is valid only if record_type is set to API.
	Code string `json:"code,omitempty"`
	// Specifies the message. This parameter is valid only if record_type is set to MEG.
	Message string `json:"message,omitempty"`
	// Specifies the execution status of the task.
	// SUCCESS: The task is successfully executed.
	// FAIL: The task failed to be executed.
	JobStatus string `json:"job_status,omitempty"`
}

type EipMetaData struct {
	// Specifies the bandwidth sharing type in the bandwidth scaling policy.
	MetadataBandwidthShareType string `json:"metadata_bandwidth_share_type,omitempty"`
	// Specifies the EIP ID for the bandwidth in the bandwidth scaling policy.
	MetadataEipId string `json:"metadata_eip_id,omitempty"`
	// Specifies the EIP for the bandwidth in the bandwidth scaling policy.
	MetadataeipAddress string `json:"metadataeip_address,omitempty"`
}
