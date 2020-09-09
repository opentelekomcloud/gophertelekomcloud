package apiversions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users.
func ListVersions(c *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, apiVersionsURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}

// ListVersionResources lists all of the different API resources for a
// particular API versions. Typical resources for Neutron might be: networks,
// subnets, etc.
func ListVersionResources(c *golangsdk.ServiceClient, v string) pagination.Pager {
	return pagination.NewPager(c, apiInfoURL(c, v), func(r pagination.PageResult) pagination.Page {
		return APIVersionResourcePage{pagination.SinglePageBase(r)}
	})
}
