package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RestartOpts struct {
	// Restart-related parameter. Value is of the type RestartInstanceInfo
	Restart RestartInstanceInfo `json:"restart,omitempty"`
}

type RestartInstanceInfo struct {
	// Restart type, which can be soft or hard.
	// soft: Only the process is restarted.
	// hard: The instance VM is forcibly restarted.
	// Enumerated values:
	// soft
	// hard
	Type string `json:"type,omitempty"`
}

func RestartInstance(client *golangsdk.ServiceClient, instanceId string, opts RestartOpts) (*ResponseInstance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v1/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var responseInstance ResponseInstance
	return &responseInstance, extract.Into(raw.Body, &responseInstance)

}

type ResponseInstance struct {
	// DDM instance ID
	InstanceId string `json:"instanceId"`
	// DDM instance name
	InstanceName string `json:"instanceName"`
	// Task ID
	JobId string `json:"jobId"`
}
