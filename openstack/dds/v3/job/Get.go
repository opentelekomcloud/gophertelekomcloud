package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get TODO: merge to common job
func Get(client *golangsdk.ServiceClient, id string) (*JobDDSInstance, error) {
	// GET /v3/{project_id}/jobs
	raw, err := client.Get(client.ServiceURL("jobs?id="+id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res JobDDSInstance
	err = extract.IntoStructPtr(raw.Body, &res, "job")
	return &res, err
}

type JobDDSInstance struct {
	// Task ID
	Id string `json:"id"`
	// Task name
	Name string `json:"name"`
	// Task execution status
	//
	// Valid value:
	// Running: The task is being executed.
	// Completed: The task is successfully executed.
	// Failed: The task fails to be executed.
	Status string `json:"status"`
	// Creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	//
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	Created string `json:"created"`
	// End time in the "yyyy-mm-ddThh:mm:ssZ" format.
	//
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	Ended string `json:"ended"`
	// Task execution progress
	//
	// NOTE:
	// The execution progress (such as "60%", indicating the task execution progress is 60%) is displayed only when the task is being executed. Otherwise, "" is returned.
	Progress string `json:"progress"`
	// Task failure information.
	FailReason string `json:"fail_reason"`
	// Instance on which the task is executed.
	Instance Instance `json:"instance"`
}

type Instance struct {
	// Instance ID
	Id string `json:"id"`
	// DB instance name
	Name string `json:"name"`
}
