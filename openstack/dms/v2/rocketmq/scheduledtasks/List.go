package scheduledtasks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// No. of the scheduled task to start the query. Must be at least 1.
	Start int `q:"start,omitempty"`
	// Number of scheduled tasks to be queried.
	Limit int `q:"limit,omitempty"`
	// Time of a scheduled task where the query starts, in the YYYYMMDDHHmmss format.
	BeginTime string `q:"begin_time,omitempty"`
	// Time of a scheduled task where the query ends, in the YYYYMMDDHHmmss format.
	EndTime string `q:"end_time,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Total number of tasks.
	JobCount int `json:"job_count"`
	// Task list.
	Jobs []ScheduledTask `json:"jobs"`
}

type ScheduledTask struct {
	// Task ID.
	ID string `json:"id"`
	// Task name.
	Name string `json:"name"`
	// Username.
	UserName string `json:"user_name"`
	// User ID.
	UserID string `json:"user_id"`
	// Task parameters.
	Params string `json:"params"`
	// Task status: CREATED, SUCCESS, FAILED, DELETED, EXECUTING or CANCELLED.
	Status string `json:"status"`
	// Creation time.
	CreatedAt string `json:"created_at"`
	// Update time.
	UpdatedAt string `json:"updated_at"`
	// Time when the scheduled task is executed.
	ScheduleAt string `json:"schedule_at"`
}

// List the scheduled tasks of an instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/scheduled-tasks
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "scheduled-tasks").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
