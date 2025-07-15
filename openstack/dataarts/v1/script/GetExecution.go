package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetExecutionResult is used to obtain the execution status and result of a script instance.
// Send request GET /v1/{project_id}/scripts/{script_name}/instances/{instance_id}
func GetExecutionResult(client *golangsdk.ServiceClient, scriptName, instanceId, workspace string) (*Script, error) {

	var opts golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(scriptsEndpoint, scriptName, instancesEndpoint, instanceId), nil, &opts)
	if err != nil {
		return nil, err
	}

	var res Script
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetExecutionResp struct {
	// Status is an execution status.
	// Can be:
	// LAUNCHING: The script instance is being submitted.
	// RUNNING: The script instance is running.
	// FINISHED: The script instance is successfully run.
	// FAILED: The script instance fails to be run.
	Status string `json:"status"`
	// Results is an execution result of the script instance.
	Results []*Result `json:"results"`
}

type Result struct {
	// Message Execution failure message.
	Message string `json:"message,omitempty"`
	// Duration is a time of executing the script instance. Unit: second
	Duration float64 `json:"duration,omitempty"`
	// RowCount is a number of the result rows.
	RowCount int `json:"rowCount,omitempty"`
	// Rows is a result data.
	Rows [][]interface{} `json:"rows,omitempty"`
	// Schema is a format definition of the result data.
	Schema []map[string]string `json:"schema,omitempty"`
}
