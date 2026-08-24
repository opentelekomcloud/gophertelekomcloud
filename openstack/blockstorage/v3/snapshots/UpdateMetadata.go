package snapshots

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateMetadataOpts struct {
	Metadata map[string]string `json:"metadata" required:"true"`
}

func UpdateMetadata(client *golangsdk.ServiceClient, id string, opts UpdateMetadataOpts) (*map[string]string, error) {
	b, err := build.RequestBody(opts, "")
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
		Metadata map[string]string `json:"metadata"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Metadata, err
}
