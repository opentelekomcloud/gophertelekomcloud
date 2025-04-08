package cloud_structuring

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListBrief(client *golangsdk.ServiceClient) ([]BriefCustomTemplate, error) {
	// GET /v3/{project_id}/lts/struct/customtemplate/list
	raw, err := client.Get(client.ServiceURL("lts", "struct", "customtemplate", "list"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []BriefCustomTemplate
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type BriefCustomTemplate struct {
	// Template creation/update time.
	CreatedAt int64 `json:"create_time"`
	// Template ID.
	ID string `json:"id"`
	// Template name.
	Name string `json:"template_name"`
	// Structuring type. Currently, regular expression, JSON, delimiters, and Nginx are supported.
	Type string `json:"template_type"`
	// Project ID.
	ProjectId string `json:"project_id"`
}
