package resource_tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListKeysOpts struct {
	// Region ID
	RegionId string `q:"region_id,omitempty"`
	// The maximum queries supported. The value 200 is used by default if this parameter is not set.
	// The value range is 1 to 200.
	Limit *int `q:"limit,omitempty"`
	// Paging location identifier (index).
	// The query starts from the next piece of data indexed by this parameter.
	// When querying the data on the first page, you do not need to specify this parameter.
	// When querying the data on subsequent pages, set this parameter to the value in the response body
	// returned by querying data of the previous page.
	// When the returned next_marker is empty, the last page has been queried.
	Marker string `q:"marker,omitempty"`
}

func ListKeys(client *golangsdk.ServiceClient, opts ListKeysOpts) ([]string, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("tag-keys").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v1.0/tag-keys
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []string
	err = extract.IntoSlicePtr(raw.Body, &res, "keys")
	return res, err
}
