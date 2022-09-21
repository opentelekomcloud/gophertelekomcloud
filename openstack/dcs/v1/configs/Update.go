package configs

import "github.com/opentelekomcloud/gophertelekomcloud"

func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToConfigUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(client.ServiceURL("instances", instanceID, "configs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
