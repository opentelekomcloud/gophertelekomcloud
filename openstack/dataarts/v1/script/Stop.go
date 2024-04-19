package script

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const stopURL = "stop"

// Stop is used to stop executing a script instance.
// Send request POST /v1/{project_id}/scripts/{script_name}/instances/{instance_id}/stop
func Stop(client *golangsdk.ServiceClient, scriptName, instanceId, workspace string) error {
	reqOpts := &golangsdk.RequestOpts{
		OkCodes: []int{204},
	}

	if workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	_, err := client.Post(client.ServiceURL(scriptsURL, scriptName, instancesURL, instanceId, stopURL), nil, nil, reqOpts)
	if err != nil {
		return err
	}
	return nil
}
