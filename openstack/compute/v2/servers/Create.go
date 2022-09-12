package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests a server to be provisioned to the user in the current tenant.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToServerCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("servers"), reqBody, nil, nil)
	return
}
