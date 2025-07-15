package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	// List of LTS configurations to be disabled.
	// To disable multiple log configurations for an instance, you need to specify multiple items.
	LtsConfigs []LtsConfig `json:"lts_configs" required:"true"`
}

type LtsConfig struct {
	// DDS Instance ID, which can be obtained by calling the API for querying instances and details.
	// If there are no instances available, create one by calling the API used for creating an instance.
	InstanceID string `json:"instance_id" required:"true"`
	// LTS log type. This parameter cannot be left empty.
	// The only supported option is audit_log.
	LogType string `json:"log_type" required:"true"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// DELETE https://{Endpoint}/v3/{project_id}/instances/logs/lts-configs
	_, err = client.DeleteWithBody(client.ServiceURL("instances", "logs", "lts-configs"), b, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return err
	}
	return
}
