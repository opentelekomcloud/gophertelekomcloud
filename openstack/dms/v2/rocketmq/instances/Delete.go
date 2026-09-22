package instances

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a RocketMQ instance to release all the resources occupied by it.
// Send DELETE to /v2/{project_id}/instances/{instance_id}
func Delete(client *golangsdk.ServiceClient, id string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", id).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
