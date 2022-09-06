package extensions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	common "github.com/opentelekomcloud/gophertelekomcloud/openstack/common/extensions"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ExtractExtensions interprets a Page as a slice of Extensions.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	return common.ExtractExtensions(page)
}

// Get retrieves information for a specific extension using its alias.
func Get(client *golangsdk.ServiceClient, alias string) common.GetResult {
	return common.Get(client, alias)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return common.List(client)
}
