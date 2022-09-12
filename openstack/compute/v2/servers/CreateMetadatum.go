package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateMetadatum will create or update the key-value pair with the given key
// for the given server ID.
func CreateMetadatum(client *golangsdk.ServiceClient, id string, opts MetadatumOptsBuilder) (r CreateMetadatumResult) {
	b, key, err := opts.ToMetadatumCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL("servers", id, "metadata", key), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
