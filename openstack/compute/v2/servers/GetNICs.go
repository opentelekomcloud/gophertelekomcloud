package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

func GetNICs(client *golangsdk.ServiceClient, id string) (r GetNICResult) {
	raw, err := client.Get(client.ServiceURL("servers", id, "os-interface"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
