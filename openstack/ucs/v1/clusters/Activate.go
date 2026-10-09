package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Activate(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Put(client.ServiceURL("clusters", id, "activation"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
