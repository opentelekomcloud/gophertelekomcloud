package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type MigrateFollowerOpts struct {
	InstanceId string `json:"-"`
	// Specifies the ID of the standby DB instance.
	NodeId string `json:"nodeId"`
	// Specifies the code of the AZ to which the standby DB instance is to be migrated.
	AzCode string `json:"azCode"`
}

func MigrateFollower(client *golangsdk.ServiceClient, opts MigrateFollowerOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/migrateslave
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "migrateslave"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		WorkflowId string `json:"workflowId"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.WorkflowId, err
}
