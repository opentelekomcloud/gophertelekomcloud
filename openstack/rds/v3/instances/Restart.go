package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type RestartRdsInstanceOpts struct {
	//
	Restart struct{} `json:"restart"`
}

type RestartRdsInstanceBuilder interface {
	ToRestartRdsInstanceMap() (map[string]interface{}, error)
}

func (opts RestartRdsInstanceOpts) ToRestartRdsInstanceMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Restart(client *golangsdk.ServiceClient, opts RestartRdsInstanceBuilder, instanceId string) (r RestartRdsInstanceResult) {
	b, err := opts.ToRestartRdsInstanceMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
