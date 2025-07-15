package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name     string `json:"message_template_name" required:"true"`
	Content  string `json:"content" required:"true"`
	Protocol string `json:"protocol" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/notifications/message_template
	raw, err := client.Post(client.ServiceURL("message_template"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateResponse struct {
	RequestID         string `json:"request_id"`
	MessageTemplateID string `json:"message_template_id"`
}
