package direct_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular direct connect based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (*DirectConnect, error) {
	raw, err := c.Get(c.ServiceURL("dcaas", "direct-connects", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DirectConnect
	err = extract.IntoStructPtr(raw.Body, &res, "direct_connect")
	return &res, err
}
