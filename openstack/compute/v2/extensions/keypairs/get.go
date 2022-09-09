package keypairs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously uploaded KeyPair.
func Get(client *golangsdk.ServiceClient, name string) (*KeyPair, error) {
	raw, err := client.Get(client.ServiceURL("os-keypairs", name), nil, nil)
	return extra(err, raw)
}
