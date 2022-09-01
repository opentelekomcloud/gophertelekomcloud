package nodes

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetJobDetails retrieves a particular job based on its unique ID
func GetJobDetails(client *golangsdk.ServiceClient, jobID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("jobs", jobID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
