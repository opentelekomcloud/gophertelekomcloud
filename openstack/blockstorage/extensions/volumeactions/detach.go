package volumeactions

import "github.com/opentelekomcloud/gophertelekomcloud"

type DetachOptsBuilder interface {
	ToVolumeDetachMap() (map[string]interface{}, error)
}

type DetachOpts struct {
	// AttachmentID is the ID of the attachment between a volume and instance.
	AttachmentID string `json:"attachment_id,omitempty"`
}

func (opts DetachOpts) ToVolumeDetachMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-detach")
}

func Detach(client *golangsdk.ServiceClient, id string, opts DetachOptsBuilder) (r DetachResult) {
	b, err := opts.ToVolumeDetachMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
