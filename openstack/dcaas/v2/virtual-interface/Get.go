package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular virtual gateway based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (*VirtualInterface, error) {
	raw, err := c.Get(c.ServiceURL("dcaas", "virtual-interfaces", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res VirtualInterface
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_interfaces")
	return &res, err
}
