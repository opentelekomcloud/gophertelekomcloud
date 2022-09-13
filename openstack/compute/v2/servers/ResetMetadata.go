package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// MetadataOpts is a map that contains key-value pairs.
type MetadataOpts map[string]string

// ResetMetadata will create multiple new key-value pairs for the given server ID.
// Note: Using this operation will erase any already-existing metadata and create the new metadata provided.
// To keep any already-existing metadata, use the UpdateMetadata function.
func ResetMetadata(client *golangsdk.ServiceClient, id string, opts map[string]string) (map[string]string, error) {
	raw, err := client.Put(client.ServiceURL("servers", id, "metadata"), map[string]interface{}{"metadata": opts},
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	return extraMet(err, raw)
}
