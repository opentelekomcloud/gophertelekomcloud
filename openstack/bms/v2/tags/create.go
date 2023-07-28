package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains all the values needed to create a new tag.
type CreateOpts struct {
	Tag []string `json:"tags" required:"true"`
}

// Create will create a new Tag based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, serverId string, opts CreateOpts) ([]string, error) {
	b, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := c.Put(c.ServiceURL("servers", serverId, "tags"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []string
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
