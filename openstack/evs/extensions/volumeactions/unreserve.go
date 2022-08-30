package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

func Unreserve(client *golangsdk.ServiceClient, id string) (err error) {
	b := map[string]interface{}{"os-unreserve": make(map[string]interface{})}
	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
