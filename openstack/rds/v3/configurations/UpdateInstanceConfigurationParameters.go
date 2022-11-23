package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/logs"
)

type UpdateInstanceConfigurationOptsBuilder interface {
	ToUpdateInstanceConfigurationMap() (map[string]interface{}, error)
}

type UpdateInstanceConfigurationOpts struct {
	Values map[string]interface{} `json:"values"`
}

func (opts UpdateInstanceConfigurationOpts) ToUpdateInstanceConfigurationMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func UpdateInstanceConfigurationParameters(client *golangsdk.ServiceClient, instanceID string, opts UpdateInstanceConfigurationOptsBuilder) (r logs.UpdateConfigurationResult) {
	b, err := opts.ToUpdateInstanceConfigurationMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Put(client.ServiceURL("instances", instanceID, "configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
