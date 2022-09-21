package configs

import "github.com/opentelekomcloud/gophertelekomcloud"

func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOptsBuilder) (err error) {
	b, err := opts.ToConfigUpdateMap()
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL("instances", instanceID, "configs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
