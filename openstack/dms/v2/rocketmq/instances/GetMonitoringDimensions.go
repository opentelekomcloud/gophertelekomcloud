package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// MonitoringDimensions is returned by GetMonitoringDimensions.
type MonitoringDimensions struct {
	// Monitoring dimensions.
	Dimensions []Dimension `json:"dimensions"`
	// Instance information.
	InstanceIDs []DimensionObject `json:"instance_ids"`
	// Node information.
	Nodes []DimensionObject `json:"nodes"`
	// Topic information.
	Topics []DimensionObject `json:"topics"`
	// Dead letter queues.
	DLQ []DimensionObject `json:"dlq"`
	// Consumer group information.
	Groups []DimensionGroup `json:"groups"`
}

type Dimension struct {
	// Monitoring dimension name.
	Name string `json:"name"`
	// Metric names.
	Metrics []string `json:"metrics"`
	// Keys used for monitoring query.
	KeyName []string `json:"key_name"`
	// Monitoring dimension route.
	DimRouter []string `json:"dim_router"`
	// Secondary dimensions.
	Children []Dimension `json:"children"`
}

type DimensionObject struct {
	// Instance ID, node name, topic name or dead letter queue name.
	Name string `json:"name"`
}

type DimensionGroup struct {
	// Consumer group name.
	Name string `json:"name"`
	// Subscribed topics.
	Topics []DimensionObject `json:"topics"`
}

// GetMonitoringDimensions queries the monitoring dimensions of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/ces-hierarchy
func GetMonitoringDimensions(client *golangsdk.ServiceClient, id string) (*MonitoringDimensions, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", id, "ces-hierarchy").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res MonitoringDimensions
	err = extract.Into(raw.Body, &res)
	return &res, err
}
