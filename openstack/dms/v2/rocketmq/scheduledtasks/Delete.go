package scheduledtasks

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a specified scheduled task of an instance.
// Send DELETE to /v2/{project_id}/instances/{instance_id}/scheduled-tasks/{task_id}
func Delete(client *golangsdk.ServiceClient, instanceID, taskID string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "scheduled-tasks", taskID).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
