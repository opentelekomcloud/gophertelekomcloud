package global_connection_bandwidth

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (*GlobalConnectionBandwidth, error) {
	raw, err := client.Get(client.ServiceURL(client.DomainID, "gcb", "gcbandwidths", id), nil, nil)
	if err != nil {
		return nil, err
	}

	return extra(raw)
}
