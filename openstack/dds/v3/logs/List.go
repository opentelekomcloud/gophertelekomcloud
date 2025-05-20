package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Index offset. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The default value is 0, indicating that the query starts from the first piece of data.
	// The value must be a positive integer.
	Offset *int `q:"offset"`
	// Number of records to be queried. The value ranges from 0 to 50.
	// If this parameter is not transferred,
	// the log configurations of the first 50 DB instances are queried by default.
	Limit *int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListConfigs, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", "logs", "lts-configs").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/logs/lts-configs
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConfigs
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConfigs struct {
	// LTS log configuration and brief information about each instance.
	InstanceLtsConfigs []InstanceLtsConfigs `json:"instance_lts_configs"`
	// Total number of cloud service log configurations that can be queried,
	// which is equal to the total number of DDS instances.
	TotalCount int `json:"total_count"`
}

type InstanceLtsConfigs struct {
	// Brief information about an instance.
	Instance *Instance `json:"instance"`
	// LTS log configuration details. If no LTS log stream is configured, no response is returned for this field.
	LtsConfigs []LtsConfigs `json:"lts_configs"`
}

type Instance struct {
	// DDS Instance ID, which can be obtained by calling the API for querying instances and details.
	// If there are no instances available, create one by calling the API used for creating an instance.
	ID string `json:"id"`
	// DDS Instance name.
	Name string `json:"name"`
	// Instance type, which can be single node, replica set, or cluster.
	// Enumerated values:
	// ReplicaSingle
	// ReplicaSet
	// Sharding
	Mode string `json:"mode"`
	// DB engine and version of the DB instance.
	Datastore *Datastore `json:"datastore"`
	// Instance status.
	Status string `json:"status"`
	// ID of the enterprise project to which the instance belongs.
	// For the default enterprise project, the value is 0.
	// For other enterprise projects, see Enterprise Management User Guide.
	EnterpriseProjectId string `json:"enterprise_project_id"`
	// All actions that are being executed on an instance.
	Actions []string `json:"actions"`
}

type LtsConfigs struct {
	// LTS log type. This parameter cannot be left empty.
	// The only supported option is audit_log.
	LogType string `json:"log_type"`
	// LTS log group ID.
	LtsGroupId string `json:"lts_group_id"`
	// LTS log stream ID.
	LtsStreamId string `json:"lts_stream_id"`
	// Indicates whether to upload logs to LTS.
	Enabled bool `json:"enabled"`
}

type Datastore struct {
	// DB engine. The value is mongodb.
	Type string `json:"type"`
	// Database major version.
	Version string `json:"version"`
}
