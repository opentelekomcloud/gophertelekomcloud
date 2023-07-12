package services

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List makes a request against the API to list services.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.Pager{
		Client:     client,
		InitialURL: listURL(client),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ServicePage{pagination.SinglePageBase(r)}
		},
	}
}
