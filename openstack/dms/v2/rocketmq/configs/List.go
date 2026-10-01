package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of configurations.
	Total int `json:"total"`
	// Offset of the next page.
	NextOffset int `json:"next_offset"`
	// Offset of the previous page.
	PreviousOffset int `json:"previous_offset"`
	// RocketMQ configurations.
	Configs []Config `json:"rocketmq_configs"`
}

// Config contains the details of a RocketMQ configuration.
type Config struct {
	// Configuration name.
	Name string `json:"name"`
	// Current value.
	Value string `json:"value"`
	// Configuration type: dynamic or static.
	ConfigType string `json:"config_type"`
	// Default value.
	DefaultValue string `json:"default_value"`
	// Value range.
	ValidValues string `json:"valid_values"`
	// Value type: integer or boolean.
	ValueType string `json:"value_type"`
}

// List the configurations of a RocketMQ instance.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/configs
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "configs").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
