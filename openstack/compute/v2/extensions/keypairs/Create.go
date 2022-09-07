package keypairs

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new KeyPair on the server, or to import a
// pre-existing keypair.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToKeyPairCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(createURL(client), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
