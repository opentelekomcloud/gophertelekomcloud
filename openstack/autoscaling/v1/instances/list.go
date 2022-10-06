package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOptsBuilder interface {
	ToInstancesListQuery() (string, error)
}

type ListOpts struct {
	LifeCycleStatus string `q:"life_cycle_state"`
	HealthStatus    string `q:"health_status"`
}

func List(client *golangsdk.ServiceClient, groupID string, opts ListOpts) ([]Instance, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("scaling_group_instance", groupID, "list")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Instance
	err = extract.IntoSlicePtr(raw.Body, &res, "scaling_group_instances")
	return res, err
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
}
