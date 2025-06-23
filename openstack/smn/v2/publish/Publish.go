package publish

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PublishOpts struct {
	TopicUrn            string            `json:"-"`
	Message             string            `json:"message,omitempty"`
	MessageStructure    string            `json:"message_structure,omitempty"`
	MessageTemplateName string            `json:"message_template_name,omitempty"`
	Tags                map[string]string `json:"tags,omitempty"`
	Subject             string            `json:"subject,omitempty"`
	TimeToLive          string            `json:"time_to_live,omitempty"`
}

func Publish(client *golangsdk.ServiceClient, opts PublishOpts) (*PublishResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/notifications/topics/{topic_urn}/publish
	raw, err := client.Post(client.ServiceURL("topics", opts.TopicUrn, "publish"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res PublishResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PublishResponse struct {
	RequestID         string `json:"request_id"`
	MessageTemplateID string `json:"message_template_id"`
}
