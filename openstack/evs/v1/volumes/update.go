package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string            `json:"display_name,omitempty"`
	Description string            `json:"display_description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Volume, error) {
	b, err := build.RequestBodyMap(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("volumes", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.IntoStructPtr(raw.Body, &res, "volume")
	return &res, err
}
