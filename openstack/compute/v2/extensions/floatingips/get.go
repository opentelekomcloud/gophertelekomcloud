package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns data about a previously created Floating IP.
func Get(client *golangsdk.ServiceClient, id string) (*FloatingIP, error) {
	raw, err := client.Get(client.ServiceURL("os-floating-ips", id), nil, nil)
	return extra(err, raw)
}
