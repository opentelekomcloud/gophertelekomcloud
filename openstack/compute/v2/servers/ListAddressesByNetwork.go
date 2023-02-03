package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListAddressesByNetwork makes a request against the API to list the servers IP addresses for the given network.
func ListAddressesByNetwork(client *golangsdk.ServiceClient, id, network string) ([]Address, error) {
	raw, err := client.Get(client.ServiceURL("servers", id, "ips", network), nil, nil)
	if err != nil {
		return nil, err
	}

	var res map[string][]Address
	err = extract.Into(raw.Body, &res)

	var key string
	for k := range res {
		key = k
	}

	return res[key], err
}
