package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetJobDetails retrieves a particular job based on its unique ID
func GetJobDetails(client *golangsdk.ServiceClient, jobID string) (*Job, error) {
	raw, err := client.Get(client.ServiceURL("jobs", jobID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res Job
	return &res, extract.Into(raw.Body, &res)
}
