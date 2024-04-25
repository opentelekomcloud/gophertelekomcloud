package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const executeEndpoint = "execute"

type ExecuteReq struct {
	Params map[string]string `json:"params,omitempty"`
}

// Execute is used to execute a specific script, which can be a DWS SQL, DLI SQL, RDS SQL, Flink SQL, Hive SQL, Presto SQL, or Spark SQL script.
// A script instance is generated each time the script is executed. You can call the API Querying the Execution Result of a Script Instance to obtain script execution results.
// Send request POST /v1/{project_id}/scripts/{script_name}/execute
func Execute(client *golangsdk.ServiceClient, scriptName, workspace string, opts *ExecuteReq) (*ExecuteResp, error) {
	var err error
	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}

	var b *build.Body
	if opts != nil {
		b, err = build.RequestBody(opts, "")
		if err != nil {
			return nil, err
		}

		reqOpts.MoreHeaders[HeaderContentType] = ApplicationJson
	}

	if workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = workspace
	}

	raw, err := client.Post(client.ServiceURL(scriptsEndpoint, scriptName, executeEndpoint), b, nil, reqOpts)
	if err != nil {
		return nil, err
	}

	var res *ExecuteResp
	err = extract.Into(raw.Body, res)
	return res, err
}

type ExecuteResp struct {
	InstanceID string `json:"instanceId"`
}
