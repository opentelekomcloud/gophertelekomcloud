package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type taskIdStr struct {
	Id string `q:"id"`
}

func ShowJobInfo(client *golangsdk.ServiceClient, taskId string) (*GetJobInfoDetail, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("jobs").WithQueryParams(&taskIdStr{Id: taskId}).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/jobs?id={id}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetJobInfoDetail
	err = extract.IntoStructPtr(raw.Body, &res, "job")
	return &res, err
}

type GetJobInfoDetail struct {
	// Task ID
	Id string `json:"id"`
	// Task name
	Name string `json:"name"`
	// Task execution status Value:
	// Running: The task is being executed.
	// Completed: The task is successfully executed.
	// Failed: The task failed to be executed.
	Status string `json:"status"`
	// Creation time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset.
	// For example, for French Winter Time (FWT), the time offset is shown as +0200.
	// The value is empty unless the instance creation is complete.
	Created string `json:"created"`
	// End time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset.
	// For example, for French Winter Time (FWT), the time offset is shown as +0200.
	// The value is empty unless the instance creation is complete.
	Ended string `json:"ended,omitempty"`
	// Task execution progress. The execution progress (such as 60%) is displayed only when the task is being executed. Otherwise, "" is returned.
	Process string `json:"process,omitempty"`
	// Instance information of the task with the specified ID
	Instance JobInstanceInfo `json:"instance"`
	// Displayed information varies depending on tasks.
	Entities JobEntities `json:"entities,omitempty"`
	// Task failure information
	FailReason string `json:"fail_reason,omitempty"`
}

type JobInstanceInfo struct {
	// Instance ID
	Id string `json:"id"`
	// Instance name
	Name string `json:"name"`
}

type JobEntities struct {
	// Instance queried in the task
	Instance JobInstance `json:"instance"`
	// Resource ID involved in a task
	ResourceIds []string `json:"resource_ids"`
}

type JobInstance struct {
	// Instance connection address
	Endpoint string `json:"endpoint"`
	// Instance type
	Type string `json:"type"`
	// Database information.
	Datastore JobDatastore `json:"datastore"`
}

type JobDatastore struct {
	// DB engine
	Type string `json:"type"`
	// DB engine version
	Version string `json:"version"`
}
