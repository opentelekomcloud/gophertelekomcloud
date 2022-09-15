package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("clusters", id), nil, nil)
	return
}
