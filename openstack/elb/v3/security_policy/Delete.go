package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("security-policies", id), &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}
