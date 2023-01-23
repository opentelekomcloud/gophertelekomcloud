package instances

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Switchover(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "switchover"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
