package volumetypes

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete deletes the volume type with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("types", id), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
