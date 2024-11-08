package lifecycle

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete an instance by id
func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("instances", id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
