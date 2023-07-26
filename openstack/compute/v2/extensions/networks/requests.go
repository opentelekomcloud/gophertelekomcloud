package networks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns a Pager that allows you to iterate over a collection of Network.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.Pager{
		Client:     client,
		InitialURL: listURL(client),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return NetworkPage{SinglePageBase: pagination.SinglePageBase{PageResult: r}}
		},
	}
}

// Get returns data about a previously created Network.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
