package keypairs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously uploaded KeyPair.
func Get(client *golangsdk.ServiceClient, name string) (r GetResult) {
	raw, err := client.Get(getURL(client, name), nil, nil)
	return
}
