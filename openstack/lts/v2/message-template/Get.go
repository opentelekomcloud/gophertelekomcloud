package message_template

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, domainId string, name string) (*MessageTemplateResponse, error) {
	// GET /v2/{project_id}/{domain_id}/lts/events/notification/template/{template_name}
	raw, err := client.Get(client.ServiceURL(domainId, "lts", "events", "notification", "template", name), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res MessageTemplateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
