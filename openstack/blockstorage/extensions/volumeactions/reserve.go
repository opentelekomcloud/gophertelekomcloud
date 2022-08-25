package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

func Reserve(client *golangsdk.ServiceClient, id string) (r ReserveResult) {
	b := map[string]interface{}{"os-reserve": make(map[string]interface{})}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
