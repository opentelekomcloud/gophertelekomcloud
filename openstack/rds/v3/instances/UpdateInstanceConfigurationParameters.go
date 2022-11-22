package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
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

func UpdateInstanceConfigurationParameters(client *golangsdk.ServiceClient, instanceID string, opts UpdateInstanceConfigurationOptsBuilder) (r UpdateConfigurationResult) {
	b, err := opts.ToUpdateInstanceConfigurationMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("instances", instanceID, "configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
