package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// DeleteOpts contains the options for deleting resources from an alarm rule.
type DeleteOpts struct {
	// Specifies the list of resources to delete.
	// A maximum of 1000 resources can be deleted at a time.
	Resources [][]Dimension `json:"resources" required:"true"`
}

// Delete deletes resources from an alarm rule in batches.
func Delete(client *golangsdk.ServiceClient, alarmId string, opts DeleteOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v2/{project_id}/alarms/{alarm_id}/resources/batch-delete
	_, err = client.Post(client.ServiceURL("alarms", alarmId, "resources", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
