package message_template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Update(client *golangsdk.ServiceClient, opts CreateOpts) (*MessageTemplateCreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/{domain_id}/lts/events/notification/templates
	raw, err := client.Put(client.ServiceURL(opts.DomainId, "lts", "events", "notification", "templates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res MessageTemplateCreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
