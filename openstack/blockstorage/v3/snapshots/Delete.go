package snapshots

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("snapshots", id), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
