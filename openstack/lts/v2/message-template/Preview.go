package message_template

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PreviewOpts struct {
	DomainId string `json:"_" required:"true"`
	// Email template content.
	Template string `json:"templates" required:"true"`
	// Language type, for example, en-us.
	Language string `json:"locale" required:"true"`
	// Source. The value can only be LTS.
	Source string `json:"source" required:"true"`
	// The content of this field is rendered to be used as the title of the message template.
	Subject string `json:"subject,omitempty"`
}

func Preview(client *golangsdk.ServiceClient, opts PreviewOpts) (*PreviewResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/{domain_id}/lts/events/notification/templates/view
	raw, err := client.Post(client.ServiceURL(opts.DomainId, "lts", "events", "notification", "templates", "view"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PreviewResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PreviewResponse struct {
	// The value is an HTML text and needs to be parsed before being displayed.
	Template string `json:"template"`
	// The title displayed after the field is parsed. It is displayed at the top of the returned HTML text.
	Subject string `json:"subject"`
}
