package nics

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular nic based on its unique ID.
func Get(c *golangsdk.ServiceClient, serverId string, id string) (*Nic, error) {
	raw, err := c.Get(c.ServiceURL("servers", serverId, "os-interface", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Nic
	err = extract.IntoStructPtr(raw.Body, &res, "interfaceAttachment")
	return &res, err
}
