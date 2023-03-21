package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular Configuration based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*IpGroup, error) {
	raw, err := client.Get(client.ServiceURL("ipgroups", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res IpGroup
	err = extract.IntoStructPtr(raw.Body, &res, "ipgroup")
	return &res, err
}
