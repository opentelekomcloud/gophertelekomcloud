package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update allows clusters to update description.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
