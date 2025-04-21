package message_template

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	DomainId string `json:"_" required:"true"`
	// Notification rule name, which is mandatory.
	// The value can contain only digits, letters, underscores (_), and hyphens (-),
	// and cannot start or end with special characters such as underscores.
	// The value can contain 1 to 100 characters and cannot be changed after being created.
	Name string `json:"name" required:"true"`
	// Notification method.
	Methods []string `json:"type,omitempty"`
	// Template description, which is mandatory.
	// The value can contain only digits, letters, and underscores (_), and cannot start or end with an underscore.
	// The value can contain 0 to 1,024 characters.
	Description string `json:"desc,omitempty"`
	// Template source. Currently, this parameter must be set to LTS. Otherwise, the template cannot be filtered.
	Source string `json:"source" required:"true"`
	// Language type, for example, en-us.
	Language string `json:"locale" required:"true"`
	// Template body is an array.
	Templates []Templates `json:"templates" required:"true"`
}

type Templates struct {
	// Template subtype, for example, sms or email.
	Type string `json:"sub_type" required:"true"`
	// Sub-template body. A variable following a dollar symbol ($) can only be one of the following variables.
	// The supported variables vary according to alarm types.
	// Currently, the variables supported for keyword alarms are as follows:
	//
	// Severity: ${event_severity};
	// Occurred: ${starts_at};
	// Alarm source: $event.metadata.resource_provider;
	// Resource type: $event.metadata.resource_type;
	// Resource ID: ${resources};
	// Statistical type: by keyword;
	// Expression: $event.annotations.condition_expression;
	// Current value: $event.annotations.current_value;
	// Statistical period: $event.annotations.frequency;
	// Query time: $event.annotations.results[0].time;
	// Query log: $event.annotations.results[0].raw_results;
	Content string `json:"content" required:"true"`
	// Email subject. This parameter is valid only when sub_type is set to email.
	Topic string `json:"topic,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*MessageTemplateCreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/{domain_id}/lts/events/notification/templates
	raw, err := client.Post(client.ServiceURL(opts.DomainId, "lts", "events", "notification", "templates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res MessageTemplateCreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type MessageTemplateCreateResponse struct {
	// Notification rule name.
	Name string `json:"name"`
	// Notification method.
	Type []string `json:"type"`
	// Template description.
	Description string `json:"desc"`
	// Template source.
	Source string `json:"source"`
	// Language.
	Language string `json:"locale"`
	// Template body, which is an array.
	Templates []TemplateResponse `json:"templates"`
}

type TemplateResponse struct {
	// Template subtype, for example, sms or email.
	Type string `json:"sub_type"`
	// Sub-template body. A variable following a dollar symbol ($) can only be one of the following variables.
	// The supported variables vary according to alarm types (keyword or SQL).
	Content string `json:"content"`
	// Email subject. This parameter is valid only when sub_type is set to email.
	Topic string `json:"topic"`
}
