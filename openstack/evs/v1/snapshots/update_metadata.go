package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateMetadataOpts struct {
	Metadata map[string]any `json:"metadata,omitempty"`
}

func UpdateMetadata(client *golangsdk.ServiceClient, id string, opts UpdateMetadataOpts) (*map[string]any, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("snapshots", id, "metadata"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Metadata map[string]any `json:"metadata"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Metadata, err
}
