package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateSecondLevelMonitoringOpts struct {
	InstanceId    string `json:"-"`
	MonitorSwitch bool   `json:"monitor_switch" required:"true"`
	Period        int    `json:"period,omitempty"`
}

func UpdateSecondLevelMonitoring(client *golangsdk.ServiceClient, opts UpdateSecondLevelMonitoringOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "monitor-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobID, extract.Into(raw.Body, &res)
}
