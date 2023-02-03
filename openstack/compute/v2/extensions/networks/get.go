package networks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get returns data about a previously created Network.
func Get(client *golangsdk.ServiceClient, id string) (*Network, error) {
	raw, err := client.Get(client.ServiceURL("os-networks", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Network
	err = extract.IntoStructPtr(raw.Body, &res, "network")
	return &res, err
}
