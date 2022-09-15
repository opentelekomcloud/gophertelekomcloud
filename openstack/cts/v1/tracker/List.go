package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

// List returns collection of Tracker. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Tracker, error) {
	var r ListResult
	_, r.Err = client.Get(client.ServiceURL("tracker"), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	allTracker, err := r.ExtractTracker()
	if err != nil {
		return nil, err
	}

	return FilterTracker(allTracker, opts)
}
