package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular tag based on its unique ID.
func Get(c *golangsdk.ServiceClient, serverId string) (*Tags, error) {
	raw, err := c.Get(c.ServiceURL("servers", serverId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Tags
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Tags struct {
	// Specifies the tags of a BMS
	Tags []string `json:"tags"`
}
