package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DisableEIP(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("apigw", "instances", id, "nat-eip"), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return err
	}

	return nil
}
