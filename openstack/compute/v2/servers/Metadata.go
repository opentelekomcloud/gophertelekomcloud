package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Metadata requests all the metadata for the given server ID.
func Metadata(client *golangsdk.ServiceClient, id string) (r GetMetadataResult) {
	raw, err := client.Get(client.ServiceURL("servers", id, "metadata"), nil, nil)
	return
}
