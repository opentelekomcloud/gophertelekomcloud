package script

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Update is used to modify the configuration items or script contents of a script.
// When modifying a script, specify the name of the script to be modified. The script name and type cannot be modified.
// Send request PUT /v1/{project_id}/scripts/{script_name}
func Update(client *golangsdk.ServiceClient, scriptName, workspace string, script Script) error {

	b, err := build.RequestBody(script, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes: []int{204},
	}

	if workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	_, err = client.Put(client.ServiceURL(scriptsEndpoint, scriptName), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
