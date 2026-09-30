package users

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a specified user of a RocketMQ instance.
// Send DELETE to /v2/{project_id}/instances/{instance_id}/users/{user_name}
func Delete(client *golangsdk.ServiceClient, instanceID, userName string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "users", userName).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
