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
	ID                string `json:"instance_id"`
	Name              string `json:"instance_name"`
	GroupID           string `json:"scaling_group_id"`
	GroupName         string `json:"scaling_group_name"`
	LifeCycleStatus   string `json:"life_cycle_state"`
	HealthStatus      string `json:"health_status"`
	ConfigurationName string `json:"scaling_configuration_name"`
	ConfigurationID   string `json:"scaling_configuration_id"`
	CreateTime        string `json:"create_time"`
}
