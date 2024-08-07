package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	coordinatorsPath = "coordinators"
)

// GetCoordinator is used to query coordinator details of a Kafka instance.
// Send GET /v2/{project_id}/instances/{instance_id}/management/coordinators
func GetCoordinator(client *golangsdk.ServiceClient, instanceId string) (*GetCoordinatorResp, error) {
	raw, err := client.Get(client.ServiceURL(instances.ResourcePath, instanceId, managementPath, coordinatorsPath), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetCoordinatorResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetCoordinatorResp struct {
	// List of coordinators of all consumer groups.
	Coordinators []*Coordinator `json:"coordinators"`
}

type Coordinator struct {
	// Consumer group ID.
	GroupId string `json:"group_id"`
	// Broker ID of the coordinator.
	Id int `json:"id"`
	// Address of the coordinator.
	Host string `json:"host"`
	// Port number.
	Port int `json:"port"`
}
