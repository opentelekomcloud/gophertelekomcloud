package parameter_configuration

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListTask(client *golangsdk.ServiceClient, clusterId string) ([]Task, error) {
	// GET /v1.0/{project_id}/clusters/{cluster_id}/ymls/joblists
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "ymls", "joblists"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Task
	err = extract.IntoSlicePtr(raw.Body, &res, "configList")
	return res, err
}

type Task struct {
	// Action ID
	ID string `json:"id"`
	// Cluster ID
	ClusterId string `json:"clusterId"`
	// Creation time. Format: Unix timestamp.
	CreatedAt int64 `json:"createAt"`
	// Task execution status.
	// true: The operation is successful.
	// false: The execution failed.
	Status string `json:"status"`
	// End time. If the creation has not been completed, the end time is null. Format: Unix timestamp.
	FinishedAt int64 `json:"finishedAt"`
	// History of parameter configuration modifications
	History string `json:"modifyDeleteReset"`
	// Returned error message. If the status is success, the value of this parameter is null.
	ErrorMsg string `json:"failedMsg"`
}
