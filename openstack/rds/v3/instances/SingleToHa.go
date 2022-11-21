package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type SingleToHaRdsOpts struct {
	SingleToHa *SingleToHaRds `json:"single_to_ha" required:"true"`
}

type SingleToHaRds struct {
	AzCodeNewNode string `json:"az_code_new_node" required:"true"`
	Password      string `json:"password,omitempty"`
}

type SingleToRdsHaBuilder interface {
	ToSingleToRdsHaMap() (map[string]interface{}, error)
}

func (opts SingleToHaRdsOpts) ToSingleToRdsHaMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func SingleToHa(client *golangsdk.ServiceClient, opts SingleToRdsHaBuilder, instanceId string) (r SingleToHaRdsInstanceResult) {
	b, err := opts.ToSingleToRdsHaMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("instances", instanceId, "action"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})

	return
}
