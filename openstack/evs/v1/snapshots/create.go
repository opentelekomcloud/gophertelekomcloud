package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	VolumeID    string                 `json:"volume_id" required:"true"`
	Description string                 `json:"display_description,omitempty"`
	Force       bool                   `json:"force,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Name        string                 `json:"display_name,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Snapshot, error) {
	b, err := golangsdk.BuildRequestBody(opts, "snapshot")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("snapshots"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Snapshot
	err = extract.IntoStructPtr(raw.Body, &res, "snapshot")
	return &res, err
}
