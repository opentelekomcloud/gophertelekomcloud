package dump

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteTransferTaskOpts struct {
	// Name of the stream.
	StreamName string
	// Name of the dump task to be deleted.
	TaskName string
}

func DeleteTransferTask(client *golangsdk.ServiceClient, opts DeleteTransferTaskOpts) (err error) {
	// DELETE /v2/{project_id}/streams/{stream_name}/transfer-tasks/{task_name}
	_, err = client.Delete(client.ServiceURL("streams", opts.StreamName, "transfer-tasks", opts.TaskName), nil)
	return
}
