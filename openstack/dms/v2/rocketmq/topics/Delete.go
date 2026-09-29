package topics

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a specified topic of a RocketMQ instance.
// Send DELETE to /v2/{project_id}/instances/{instance_id}/topics/{topic}
func Delete(client *golangsdk.ServiceClient, instanceID, topic string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics", topic).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
