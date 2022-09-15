package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err = client.Delete(client.ServiceURL("clusters", id), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	return
}
