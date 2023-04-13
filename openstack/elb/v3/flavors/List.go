package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorListMap() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through the API.
type ListOpts struct {
	// Specifies the id.
	ID []string `q:"id"`
	// Specifies the name.
	Name []string `q:"name"`
	// Specifies whether shared.
	Shared *bool `q:"shared"`
	// Specifies the type.
	Type []string `q:"type"`
}

// ToFlavorListMap formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListMap() (string, error) {
	s, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return s.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// flavors.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("flavors")
	if opts != nil {
		queryString, err := opts.ToFlavorListMap()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += queryString
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return FlavorPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
