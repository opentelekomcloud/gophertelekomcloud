package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain log configurations.
func GetLogConfig(client *golangsdk.ServiceClient, opts QueryParameters) (*LogConfig, error) {
	// GET /v1/{project_id}/cfw/logs/configuration
	url, err := golangsdk.NewURLBuilder().WithEndpoints("cfw", "logs", "configuration").WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

type GetResponse struct {
	// Log configurations
	Data LogConfig `json:"data"`
}

type LogConfig struct {
	// Firewall ID.
	FWInstanceID string `json:"fw_instance_id"`

	// Whether to enable LTS: 1 (yes), 0 (no).
	LtsEnable int `json:"lts_enable"`

	// Log Tank Service (LTS) log group ID.
	LtsLogGroupID string `json:"lts_log_group_id"`

	// Attack log stream ID.
	LtsAttackLogStreamID string `json:"lts_attack_log_stream_id"`

	// Whether to enable the attack log stream: 1 (yes), 0 (no).
	LtsAttackLogStreamEnable int `json:"lts_attack_log_stream_enable"`

	// Access control log stream ID.
	LtsAccessLogStreamID string `json:"lts_access_log_stream_id"`

	// Whether to enable the access control stream: 1 (yes), 0 (no).
	LtsAccessLogStreamEnable int `json:"lts_access_log_stream_enable"`

	// Traffic log ID.
	LtsFlowLogStreamID string `json:"lts_flow_log_stream_id"`

	// Whether to enable the traffic log function: 1 (yes), 0 (no).
	LtsFlowLogStreamEnable int `json:"lts_flow_log_stream_enable"`
}
