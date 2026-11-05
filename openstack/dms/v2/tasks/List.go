package tasks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Start     int    `q:"start,omitempty"`
	Limit     int    `q:"limit,omitempty"`
	BeginTime string `q:"begin_time,omitempty"`
	EndTime   string `q:"end_time,omitempty"`
}

// List sends GET /v2/{project_id}/instances/{instance_id}/tasks
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*TasksResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "tasks").
		WithQueryParams(&opts).
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

type TasksResponse struct {
	TaskCount string `json:"task_count"`
	Tasks     []Task `json:"tasks"`
}

type Task struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	UserID    string `json:"user_id"`
	Params    string `json:"params"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
