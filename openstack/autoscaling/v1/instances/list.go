package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies the AS group ID.
	ScalingGroupId string
	// Specifies the instance lifecycle status in the AS group.
	// INSERVICE: The instance is enabled.
	// PENDING: The instance is being added to the AS group.
	// REMOVING: The instance is being removed from the AS group.
	LifeCycleState string `q:"life_cycle_state,omitempty"`
	// Specifies the instance health status.
	// INITIALIZING: The instance is initializing.
	// NORMAL: The instance is normal.
	// ERROR: The instance is abnormal.
	HealthStatus string `q:"health_status,omitempty"`
	// Specifies the instance protection status.
	// true: Instance protection is enabled.
	// false: Instance protection is disabled.
	ProtectFromScalingDown string `q:"protect_from_scaling_down,omitempty"`
	// Specifies the start line number. The default value is 0. The minimum parameter value is 0.
	StartNumber int32 `q:"start_number,omitempty"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int32 `q:"limit,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListScalingInstancesResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /autoscaling-api/v1/{project_id}/scaling_group_instance/{scaling_group_id}/list
	raw, err := client.Get(client.ServiceURL("scaling_group_instance", opts.ScalingGroupId, "list")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingInstancesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListScalingInstancesResponse struct {
	TotalNumber           int32      `json:"total_number,omitempty"`
	StartNumber           int32      `json:"start_number,omitempty"`
	Limit                 int32      `json:"limit,omitempty"`
	ScalingGroupInstances []Instance `json:"scaling_group_instances,omitempty"`
}

type Instance struct {
	// Specifies the instance ID.
	ID string `json:"instance_id"`
	// Specifies the instance name.
	Name string `json:"instance_name"`
	// Specifies the ID of the AS group to which the instance belongs.
	GroupID string `json:"scaling_group_id"`
	// Specifies the name of the AS group to which the instance belongs.
	// Supports fuzzy search.
	GroupName string `json:"scaling_group_name"`
	// Specifies the instance lifecycle status in the AS group.
	// INSERVICE: The instance is enabled.
	// PENDING: The instance is being added to the AS group.
	// REMOVING: The instance is being removed from the AS group.
	LifeCycleStatus string `json:"life_cycle_state"`
	// Specifies the instance health status.
	// INITIALIZING: The instance is being initialized.
	// NORMAL: The instance is functional.
	// ERROR: The instance is faulty.
	HealthStatus string `json:"health_status"`
	// Specifies the AS configuration name.
	ConfigurationName string `json:"scaling_configuration_name"`
	// Specifies the AS configuration ID.
	// If the returned value is not empty, the instance is an ECS automatically created in a scaling action.
	// If the returned value is empty, the instance is an ECS manually added to the AS group.
	ConfigurationID string `json:"scaling_configuration_id"`
	// Specifies the time when the instance is added to the AS group. The time format complies with UTC.
	CreateTime string `json:"create_time"`
	// Specifies the instance protection status.
	ProtectFromScalingDown bool `json:"protect_from_scaling_down"`
}
