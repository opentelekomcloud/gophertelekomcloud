package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Delete(client *golangsdk.ServiceClient, instanceID string) (*DeleteResponse, error) {
	raw, err := client.Delete(client.ServiceURL("instances", instanceID), &golangsdk.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	return &res, extract.Into(raw.Body, &res)
}

type DeleteResponse struct {
	JobId string `json:"job_id"`
}
