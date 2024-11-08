package compliance

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, domainId, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("resource-manager", "domains", domainId, "policy-assignments", id), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
