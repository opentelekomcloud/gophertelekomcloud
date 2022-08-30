package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

func BeginDetaching(client *golangsdk.ServiceClient, id string) (err error) {
	b := map[string]interface{}{"os-begin_detaching": make(map[string]interface{})}
	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
