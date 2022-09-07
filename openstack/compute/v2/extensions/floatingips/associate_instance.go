package floatingips

import "github.com/opentelekomcloud/gophertelekomcloud"

// AssociateInstance pairs an allocated Floating IP with a server.
func AssociateInstance(client *golangsdk.ServiceClient, serverID string, opts AssociateOptsBuilder) (r AssociateResult) {
	b, err := opts.ToFloatingIPAssociateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(associateURL(client, serverID), b, nil, nil)
	return
}
