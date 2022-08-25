package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Volume, error) {
	b, err := golangsdk.BuildRequestBody(opts, "volume")
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
	err = extract.Into(raw.Body, &res)
	return &res, err
}
