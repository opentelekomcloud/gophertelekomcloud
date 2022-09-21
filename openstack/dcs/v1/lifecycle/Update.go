package lifecycle

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update is a method which can be able to update the instance
// via accessing to the service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	body, err := opts.ToInstanceUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", id), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
