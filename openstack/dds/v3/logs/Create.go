package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// Each item indicates an LTS configuration for the instance.
	LtsConfigs []Configs `json:"lts_configs" required:"true"`
}

type Configs struct {
	// DDS Instance ID, which can be obtained by calling the API for querying instances and details.
	// If there are no instances available, create one by calling the API used for creating an instance.
	InstanceID string `json:"instance_id" required:"true"`
	// LTS log type. This parameter cannot be left empty.
	// The only supported option is audit_log.
	LogType string `json:"log_type" required:"true"`
	// LTS log group ID.
	// You can obtain the value using the LTS API for querying all log groups under an account.
	LtsGroupId string `json:"lts_group_id" required:"true"`
	// LTS log stream ID.
	// You can obtain the value using the LTS API for querying all log streams in a specified log group.
	LtsStreamId string `json:"lts_stream_id" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/logs/lts-configs
	_, err = client.Post(client.ServiceURL("instances", "logs", "lts-configs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return err
	}
	return
}
