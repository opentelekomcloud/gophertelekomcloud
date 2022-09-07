package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create requests the creation of a new volume attachment on the server.
func Create(client *golangsdk.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeAttachmentCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(createURL(client, serverID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
