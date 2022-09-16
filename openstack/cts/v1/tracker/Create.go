package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new tracker based on the values in CreateOpts. To extract
// the tracker name  from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("tracker"), b, nil, nil)
	return
}
