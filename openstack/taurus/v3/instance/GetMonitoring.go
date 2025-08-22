package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetSecondLevelMonitoring(client *golangsdk.ServiceClient, instanceId string) (*SecondLevelMonitoring, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "monitor-policy"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res SecondLevelMonitoring
	return &res, extract.Into(raw.Body, &res)
}

type SecondLevelMonitoring struct {
	MonitorSwitch bool `json:"monitor_switch"`
	Period        int  `json:"period"`
}
