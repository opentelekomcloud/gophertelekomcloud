package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DisableFederation(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("clustergroups", id, "federations"), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
