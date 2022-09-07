package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update will modify the mutable properties of a security group, notably its
// name and description.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSecGroupUpdateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(resourceURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
