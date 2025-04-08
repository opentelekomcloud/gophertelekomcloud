package log_collection

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Disable(client *golangsdk.ServiceClient) error {
	// POST /v2/{project_id}/collection/disable
	_, err := client.Post(client.ServiceURL("collection", "disable"), nil, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return err
	}
	return err
}
