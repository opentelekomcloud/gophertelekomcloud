package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func UnBindEIP(client *golangsdk.ServiceClient, nodeId string) (*string, error) {
	raw, err := client.Post(client.ServiceURL("nodes", nodeId, "unbind-eip"), nil, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return extractJob(err, raw)
}
