package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new Floating IP.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFloatingIPCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("os-floating-ips"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
