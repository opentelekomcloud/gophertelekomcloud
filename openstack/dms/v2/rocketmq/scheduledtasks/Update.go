package scheduledtasks

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type UpdateOpts struct {
	// New execution time of the scheduled task, for example 20240101000000.
	ExecuteAt string `q:"execute_at,omitempty"`
	// New status of the scheduled task: CANCELLED.
	Status string `q:"status,omitempty"`
}

// Update a specified scheduled task of an instance.
// Send PUT to /v2/{project_id}/instances/{instance_id}/scheduled-tasks/{task_id}
func Update(client *golangsdk.ServiceClient, instanceID, taskID string, opts UpdateOpts) error {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "scheduled-tasks", taskID).
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
