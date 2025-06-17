package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, templateId string) (*Template, error) {
	// GET /v2/{project_id}/notifications/message_template/{message_template_id}
	raw, err := client.Get(client.ServiceURL("message_template", templateId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Template
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Template struct {
	MessageTemplateID string   `json:"message_template_id"`
	Name              string   `json:"message_template_name"`
	Protocol          string   `json:"protocol"`
	TagNames          []string `json:"tag_names"`
	CreateTime        string   `json:"create_time"`
	UpdateTime        string   `json:"update_time"`
	Content           string   `json:"content,omitempty"`
	RequestID         string   `json:"request_id,omitempty"`
}
