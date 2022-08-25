package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

func BeginDetaching(client *golangsdk.ServiceClient, id string) (r BeginDetachingResult) {
	b := map[string]interface{}{"os-begin_detaching": make(map[string]interface{})}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
