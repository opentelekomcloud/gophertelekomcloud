package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	VolumeID    string            `json:"volume_id" required:"true"`
	Force       bool              `json:"force,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Snapshot, error) {
	b, err := golangsdk.BuildRequestBody(opts, "snapshot")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("snapshots"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Snapshot Snapshot `json:"snapshot"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Snapshot, err
}
