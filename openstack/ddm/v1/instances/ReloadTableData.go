package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func ReloadTableData(client *golangsdk.ServiceClient, instanceId string) error {
	// POST /v1/{project_id}/instances/{instance_id}/reload-config
	_, err := client.Post(client.ServiceURL("instances", instanceId, "reload-config"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return err
	}
	return nil
}
