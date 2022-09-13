package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListAddresses makes a request against the API to list the servers IP addresses.
func ListAddresses(client *golangsdk.ServiceClient, id string) (map[string][]Address, error) {
	raw, err := client.Get(client.ServiceURL("servers", id, "ips"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Addresses map[string][]Address `json:"addresses"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Addresses, err
}
