package central_network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Get(client *golangsdk.ServiceClient, centralNetworkId string) (*CentralNetwork, error) {
	raw, err := client.Get(client.ServiceURL(client.DomainID, "gcn", "central-networks", centralNetworkId), nil, nil)
	if err != nil {
		return nil, err
	}

	return extra(raw)
}
