package attachinterfaces

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new interface attachment on the server.
func Create(client *golangsdk.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAttachInterfacesCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(createInterfaceURL(client, serverID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
