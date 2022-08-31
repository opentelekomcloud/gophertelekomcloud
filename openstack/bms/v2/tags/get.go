package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular tag based on its unique ID.
func Get(c *golangsdk.ServiceClient, serverId string) ([]string, error) {
	raw, err := c.Get(c.ServiceURL("servers", serverId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []string
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
