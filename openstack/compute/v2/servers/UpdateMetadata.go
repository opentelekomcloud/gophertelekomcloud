package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateMetadata updates (or creates) all the metadata specified by opts for
// the given server ID. This operation does not affect already-existing metadata
// that is not specified by opts.
func UpdateMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOpts) (map[string]string, error) {
	b, err := opts.ToMetadataUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers", id, "metadata"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraMet(err, raw)
}
