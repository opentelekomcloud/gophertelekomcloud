package message_template

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, domainId string) ([]MessageTemplateResponse, error) {
	// GET /v2/{project_id}/{domain_id}/lts/events/notification/templates
	raw, err := client.Get(client.ServiceURL(domainId, "lts", "events", "notification", "templates"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []MessageTemplateResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type MessageTemplateResponse struct {
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
	// Creation time (timestamp in milliseconds).
	CreatedAt int64 `json:"create_time"`
	// Update time (timestamp in milliseconds).
	UpdatedAt int64 `json:"modify_time"`
	// Project ID
	ProjectId string `json:"project_id"`
}
