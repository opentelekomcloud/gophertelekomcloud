package tracker

import "github.com/opentelekomcloud/gophertelekomcloud"

func Update(client *golangsdk.ServiceClient, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTrackerUpdateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL("tracker", "system"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
