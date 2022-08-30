package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type DetachOpts struct {
	// AttachmentID is the ID of the attachment between a volume and instance.
	AttachmentID string `json:"attachment_id,omitempty"`
}

func Detach(client *golangsdk.ServiceClient, id string, opts DetachOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "os-detach")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("volumes", id, "action"), b, nil, nil)
	return
}
