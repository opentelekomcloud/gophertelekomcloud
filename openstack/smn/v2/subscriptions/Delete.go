package subscriptions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, subscriptionUrn string) (err error) {
	_, err = client.Delete(client.ServiceURL("subscriptions", subscriptionUrn), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	return
}
