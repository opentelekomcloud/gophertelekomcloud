package dump

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type TransferTaskActionOpts struct {
	// Name of the stream to be queried.
	// Maximum: 60
	StreamName string `json:"stream_name"`
	// Dump task operation.
	// Currently, only the following operation is supported:
	// - stop: The dump task is stopped.
	// Enumeration values:
	//  stop
	// - start: The dump task is started.
	// Enumeration values:
	//  start
	Action string `json:"action"`
	// List of dump tasks to be paused.
	Tasks []BatchTransferTask `json:"tasks"`
}

type BatchTransferTask struct {
	// Dump task ID.
	Id string `json:"id"`
}

func TransferTaskAction(client *golangsdk.ServiceClient, opts TransferTaskActionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/streams/{stream_name}/transfer-tasks/action
	_, err = client.Post(client.ServiceURL("streams", opts.StreamName, "transfer-tasks", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
