package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetMonitoringDimensions sends GET /v2/{project_id}/instances/{instance_id}/ces-hierarchy
func GetMonitoringDimensions(client *golangsdk.ServiceClient, instanceId string) (*GetMonitoringDimensionsResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceId, "ces-hierarchy").
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetMonitoringDimensionsResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetMonitoringDimensionsResp struct {
	Dimensions  []Dimension       `json:"dimensions"`
	InstanceIDs []DimensionObject `json:"instance_ids"`
	Nodes       []DimensionObject `json:"nodes"`
	Queues      []DimensionQueue  `json:"queues"`
	Groups      []DimensionGroup  `json:"groups"`
}

type Dimension struct {
	Name      string           `json:"name"`
	Metrics   []string         `json:"metrics"`
	KeyName   []string         `json:"key_name"`
	DimRouter []string         `json:"dim_router"`
	Children  []DimensionChild `json:"children"`
}

type DimensionChild struct {
	Name      string   `json:"name"`
	Metrics   []string `json:"metrics"`
	KeyName   []string `json:"key_name"`
	DimRouter []string `json:"dim_router"`
}

type DimensionObject struct {
	Name string `json:"name"`
}

type DimensionQueue struct {
	Name       string            `json:"name"`
	Partitions []DimensionObject `json:"partitions"`
}

type DimensionGroup struct {
	Name   string           `json:"name"`
	Queues []DimensionQueue `json:"queues"`
}
