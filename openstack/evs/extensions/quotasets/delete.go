package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, projectID string) (err error) {
	_, err = client.Delete(client.ServiceURL("os-quota-sets", projectID), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
