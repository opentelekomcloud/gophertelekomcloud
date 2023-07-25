package extensions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Get retrieves information for a specific extension using its alias.
func Get(c *golangsdk.ServiceClient, alias string) (r GetResult) {
	_, r.Err = c.Get(ExtensionURL(c, alias), &r.Body, nil)
	return
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *golangsdk.ServiceClient) pagination.Pager {
	return pagination.Pager{
		Client:     c,
		InitialURL: ListExtensionURL(c),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return ExtensionPage{pagination.SinglePageBase{r}}
		},
	}
}
