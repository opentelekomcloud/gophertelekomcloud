package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

func Update(client *golangsdk.ServiceClient, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTrackerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("tracker", "system"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
