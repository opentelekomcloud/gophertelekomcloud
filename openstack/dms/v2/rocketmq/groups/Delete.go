package groups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a specified consumer group of a RocketMQ instance.
// Send DELETE to /v2/{project_id}/instances/{instance_id}/groups/{group}
func Delete(client *golangsdk.ServiceClient, instanceID, group string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "groups", group).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
