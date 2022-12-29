package instances

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, instanceId string) (*string, error) {
	raw, err := client.Delete(client.ServiceURL("instances", instanceId), nil)
	return extractJob(err, raw)
}
