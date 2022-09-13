package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// ResetMetadata will create multiple new key-value pairs for the given server ID.
// Note: Using this operation will erase any already-existing metadata and
// create the new metadata provided. To keep any already-existing metadata,
// use the UpdateMetadata function.
func ResetMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOpts) (map[string]string, error) {
	b, err := opts.ToMetadataResetMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("servers", id, "metadata"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraMet(err, raw)
}
