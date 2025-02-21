package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to delete a firewall. It takes effect only for pay-per-use firewalls.
func Delete(client *golangsdk.ServiceClient, resourceId string) (*string, error) {
	// DELETE /v2/{project_id}/firewall/{resource_id}
	raw, err := client.Delete(client.ServiceURL("firewall", resourceId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})

	if err != nil {
		return nil, err
	}

	var res DeleteJob
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}

type DeleteJob struct {
	// ID of a firewall deletion task.
	JobId string `json:"data"`
}
