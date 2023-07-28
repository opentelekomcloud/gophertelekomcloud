package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
)

type ListScalingActivityLogsOpts struct {
	ScalingGroupId string
	// Specifies the scaling action log ID.
	LogId string `q:"log_id,omitempty"`
	// Specifies the start time that complies with UTC for querying scaling action logs.
	// The format of the start time is yyyy-MM-ddThh:mm:ssZ.
	StartTime string `q:"start_time,omitempty"`
	// Specifies the end time that complies with UTC for querying scaling action logs.
	// The format of the end time is yyyy-MM-ddThh:mm:ssZ.
	EndTime string `q:"end_time,omitempty"`
	// Specifies the start line number. The default value is 0. The minimum parameter value is 0.
	StartNumber int32 `q:"start_number,omitempty"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int32 `q:"limit,omitempty"`
	// Specifies the types of the scaling actions to be queried. Different types are separated by commas (,).
	//
	// NORMAL: indicates a common scaling action.
	// MANUAL_REMOVE: indicates manually removing instances from an AS group.
	// MANUAL_DELETE: indicates manually removing and deleting instances from an AS group.
	// MANUAL_ADD: indicates manually adding instances to an AS group.
	// ELB_CHECK_DELETE: indicates that instances are removed from an AS group and deleted based on the ELB health check result.
	// AUDIT_CHECK_DELETE: indicates that instances are removed from an AS group and deleted based on the OpenStack audit.
	// DIFF: indicates that the number of expected instances is different from the actual number of instances.
	// MODIFY_ELB: indicates the load balancer migration.
	Type string `q:"type,omitempty"`
	// Specifies the status of the scaling action.
	//
	// SUCCESS: The scaling action has been performed.
	// FAIL: Performing the scaling action failed.
	// DOING: The scaling action is being performed.
	Status string `q:"status,omitempty"`
}

func ListScalingActivityLogs(client *golangsdk.ServiceClient, opts ListScalingActivityLogsOpts) (*ListScalingActivityLogsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /autoscaling-api/v2/{project_id}/scaling_activity_log/{scaling_group_id}
	raw, err := client.Get(client.ServiceURL("scaling_activity_log", opts.ScalingGroupId)+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingActivityLogsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListScalingActivityLogsResponse struct {
	TotalNumber        int32                `json:"total_number,omitempty"`
	StartNumber        int32                `json:"start_number,omitempty"`
	Limit              int32                `json:"limit,omitempty"`
	ScalingActivityLog []ScalingActivityLog `json:"scaling_activity_log,omitempty"`
}

type ScalingActivityLog struct {
	// Specifies the status of the scaling action.
	//
	// SUCCESS: The scaling action has been performed.
	// FAIL: Performing the scaling action failed.
	// DOING: The scaling action is being performed.
	Status string `json:"status,omitempty"`
	// Specifies the start time of the scaling action. The time format must comply with UTC.
	StartTime string `json:"start_time,omitempty"`
	// Specifies the end time of the scaling action. The time format must comply with UTC.
	EndTime string `json:"end_time,omitempty"`
	// Specifies the scaling action log ID.
	Id string `json:"id,omitempty"`
	// Specifies names of the ECSs that are removed from the AS group in a scaling action.
	InstanceRemovedList []ScalingInstance `json:"instance_removed_list,omitempty"`
	// Specifies names of the ECSs that are removed from the AS group and deleted in a scaling action.
	InstanceDeletedList []ScalingInstance `json:"instance_deleted_list,omitempty"`
	// Specifies names of the ECSs that are added to the AS group in a scaling action.
	InstanceAddedList []ScalingInstance `json:"instance_added_list,omitempty"`
	// Specifies the ECSs for which a scaling action fails.
	InstanceFailedList []ScalingInstance `json:"instance_failed_list,omitempty"`
	// Specifies the ECSs that are set to standby mode or for which standby mode is canceled in a scaling action.
	InstanceStandbyList []ScalingInstance `json:"instance_standby_list,omitempty"`
	// Specifies the number of added or deleted instances during the scaling.
	ScalingValue string `json:"scaling_value,omitempty"`
	// Specifies the description of the scaling action.
	Description string `json:"description,omitempty"`
	// Specifies the number of instances in the AS group.
	InstanceValue int32 `json:"instance_value,omitempty"`
	// Specifies the expected number of instances for the scaling action.
	DesireValue int32 `json:"desire_value,omitempty"`
	// Specifies the load balancers that are bound to the AS group.
	LbBindSuccessList []ModifyLb `json:"lb_bind_success_list,omitempty"`
	// Specifies the load balancers that failed to be bound to the AS group.
	LbBindFailedList []ModifyLb `json:"lb_bind_failed_list,omitempty"`
	// Specifies the load balancers that are unbound from the AS group.
	LbUnbindSuccessList []ModifyLb `json:"lb_unbind_success_list,omitempty"`
	// Specifies the load balancers that failed to be unbound from the AS group.
	LbUnbindFailedList []ModifyLb `json:"lb_unbind_failed_list,omitempty"`
	// Specifies the type of the scaling action.
	Type string `json:"type,omitempty"`
}

type ScalingInstance struct {
	// Specifies the ECS name.
	InstanceName string `json:"instance_name,omitempty"`
	// Specifies the ECS ID.
	InstanceId string `json:"instance_id,omitempty"`
	// Specifies the cause of the instance scaling failure.
	FailedReason string `json:"failed_reason,omitempty"`
	// Specifies details of the instance scaling failure.
	FailedDetails string `json:"failed_details,omitempty"`
	// Specifies the information about instance configurations.
	InstanceConfig string `json:"instance_config,omitempty"`
}

type ModifyLb struct {
	// Specifies information about an enhanced load balancer.
	LbaasListener groups.LBaaSListener `json:"lbaas_listener,omitempty"`
	// Specifies information about a classic load balancer.
	Listener string `json:"listener,omitempty"`
	// Specifies the cause of a load balancer migration failure.
	FailedReason string `json:"failed_reason,omitempty"`
	// Specifies the details of a load balancer migration failure.
	FailedDetails string `json:"failed_details,omitempty"`
}
