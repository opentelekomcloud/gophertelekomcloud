package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SingleToHaRdsOpts struct {
	InstanceId string
	SingleToHa SingleToHaRds `json:"single_to_ha" required:"true"`
}

type SingleToHaRds struct {
	// Specifies the AZ code of the DB instance node.
	AzCodeNewNode string `json:"az_code_new_node" required:"true"`
	// This parameter is mandatory only when a Microsoft SQL Server DB instance type is changed from single to primary/standby.
	Password string `json:"password,omitempty"`
}

func SingleToHa(client *golangsdk.ServiceClient, opts SingleToHaRdsOpts) (*string, error) {
	b, err := build.RequestBody(&opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/action
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extraJob(err, raw)
}
