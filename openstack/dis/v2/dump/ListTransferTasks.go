package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListTransferTasks(client *golangsdk.ServiceClient, streamName string) (*ListTransferTasksResponse, error) {
	// GET /v2/{project_id}/streams/{stream_name}/transfer-tasks
	raw, err := client.Get(client.ServiceURL("streams", streamName, "transfer-tasks"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListTransferTasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListTransferTasksResponse struct {
	// Total number of dump tasks.
	TotalNumber *int `json:"total_number,omitempty"`
	// List of dump tasks.
	Tasks []TransferTask `json:"tasks,omitempty"`
}

type TransferTask struct {
	// Name of the dump task.
	TaskName string `json:"task_name,omitempty"`
	// Id of the dump task
	TaskId string `json:"task_id,omitempty"`
	// Dump task status. Possible values:
	// ERROR: An error occurs.
	// STARTING: The dump task is being started.
	// PAUSED: The dump task has been stopped.
	// RUNNING: The dump task is running.
	// DELETE: The dump task has been deleted.
	// ABNORMAL: The dump task is abnormal.
	// Enumeration values:
	// ERROR
	// STARTING
	// PAUSED
	// RUNNING
	// DELETE
	// ABNORMAL
	State string `json:"state,omitempty"`
	// Dump destination. Possible values:
	// OBS: Data is dumped to OBS.
	// Enumeration values:
	// OBS
	DestinationType string `json:"destination_type,omitempty"`
	// Time when the dump task is created.
	CreateTime *int64 `json:"create_time,omitempty"`
	// Latest dump time of the dump task.
	LastTransferTimestamp *int64 `json:"last_transfer_timestamp,omitempty"`
}
