package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new tracker based on the values in CreateOpts. To extract
// the tracker name  from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToTrackerCreateMap()

	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("tracker"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
