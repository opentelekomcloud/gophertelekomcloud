package configs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

func List(client *golangsdk.ServiceClient, instanceID string) (r ListResult) {
	_, r.Err = client.Get(rootURL(client, instanceID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	RedisConfigs []RedisConfig `json:"redis_config" required:"true"`
}

type RedisConfig struct {
	ParamID    string `json:"param_id" required:"true"`
	ParamName  string `json:"param_name" required:"true"`
	ParamValue string `json:"param_value" required:"true"`
}

func (opts UpdateOpts) ToConfigUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToConfigUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(rootURL(client, instanceID), b, nil, openstack.StdRequestOpts())
	return
}
