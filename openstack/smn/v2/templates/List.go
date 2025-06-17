package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset   string `q:"offset,omitempty"`
	Limit    int    `q:"limit,omitempty"`
	Name     string `q:"message_template_name,omitempty"`
	Protocol string `q:"protocol,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("message_template").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET v2/{project_id}/notifications/message_template
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	RequestId            string     `json:"request_id"`
	MessageTemplateCount int        `json:"message_template_count"`
	MessageTemplates     []Template `json:"message_templates"`
}
