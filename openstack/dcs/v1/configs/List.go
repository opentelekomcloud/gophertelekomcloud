package configs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, instanceID string) (*ConfigParam, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceID, "configs"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigParam
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ConfigParam struct {
	// Current status of a DCS instance.
	Status string `json:"status"`
	// DCS instance ID.
	InstanceID string `json:"instance_id"`
	// Array of configuration items of the DCS instance.
	RedisConfigs []RedisConfigResult `json:"redis_config"`
	// DCS instance status that is being modified or has been modified. Options:
	// UPDATING
	// FAILURE
	// SUCCESS
	ConfigStatus string `json:"config_status"`
	// Time at which the DCS instance is operated on. For example, 2017-03-31T12:24:46.297Z.
	ConfigTime string `json:"config_time"`
	// Instance type. If true is returned, the instance is a Proxy Cluster DCS Redis 3.0 instance.
	// If false is returned, the instance is not a Proxy Cluster DCS Redis 3.0 instance.
	ClusterV1 bool `json:"cluster_v1"`
}

type RedisConfigResult struct {
	// Configuration item description.
	Description string `json:"description"`
	// Configuration parameter ID. For the possible values
	ParamID string `json:"param_id"`
	// Configuration parameter name. For the possible values
	ParamName string `json:"param_name"`
	// Configuration parameter value.
	ParamValue string `json:"param_value"`
	// Default value of the configuration parameter. For the possible values
	DefaultValue string `json:"default_value"`
	// Type of the configuration parameter value. For the possible values
	ValueType string `json:"value_type"`
	// Range of the configuration parameter value. For the possible values
	ValueRange string `json:"value_range"`
	// If null or empty is returned, the node is a default node, that is, the Redis Server node.
	// If proxy is returned, the node is a proxy node.
	NodeRole string `json:"node_role"`
}
