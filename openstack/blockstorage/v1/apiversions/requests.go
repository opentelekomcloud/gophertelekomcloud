package apiversions

import (
	"net/url"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List lists all the Cinder API versions available to end-users.
func List(c *golangsdk.ServiceClient) pagination.Pager {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return pagination.NewPager(c, u.String(), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}

// Get will retrieve the volume type with the provided ID. To extract the volume
// type from the result, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, v string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL(strings.TrimRight(v, "/")+"/"), nil, nil)
	return
}
