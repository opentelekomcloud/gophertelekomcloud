package groups_hcs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update is a method which can be able to update the group via accessing to the
// autoscaling service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	body, err := opts.ToGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(client.ServiceURL("scaling_group", id), body, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
