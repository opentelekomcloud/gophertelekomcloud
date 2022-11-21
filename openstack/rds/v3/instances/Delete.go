package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteInstance struct {
	InstanceId string `json:"instance_id" required:"true"`
}

type DeleteInstanceBuilder interface {
	ToInstancesDeleteMap() (map[string]interface{}, error)
}

func (opts DeleteInstance) ToInstancesDeleteMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Delete(client *golangsdk.ServiceClient, instanceId string) (r DeleteInstanceRdsResult) {
	url := client.ServiceURL("instances", instanceId)

	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{JSONResponse: &r.Body, MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	return
}
