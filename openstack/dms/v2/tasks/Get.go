package tasks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get sends GET /v2/{project_id}/instances/{instance_id}/tasks/{task_id}
func Get(client *golangsdk.ServiceClient, instanceID, taskID string) (*TasksResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "tasks", taskID).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
