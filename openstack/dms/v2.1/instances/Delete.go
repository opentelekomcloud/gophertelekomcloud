package instances

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete an instance by id
// Send DELETE /v2/{project_id}/instances/{instance_id}
func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL(ResourcePath, id), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
