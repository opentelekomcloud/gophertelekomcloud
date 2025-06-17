package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	TemplateID string `json:"-"`
	Content    string `json:"content" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*UpdateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/notifications/message_template/{message_template_id}
	raw, err := client.Put(client.ServiceURL("message_template", opts.TemplateID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateResponse struct {
	RequestID string `json:"request_id"`
}
