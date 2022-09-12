package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateImage makes a request against the nova API to schedule an image to be
// created of the server
func CreateImage(client *golangsdk.ServiceClient, id string, opts CreateImageOptsBuilder) (r CreateImageResult) {
	b, err := opts.ToServerCreateImageMap()
	if err != nil {
		return nil, err
	}
	resp, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}
	r.Header = resp.Header
	return
}
