package cloud_structuring

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	ID string `q:"id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]CustomTemplate, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("lts", "struct", "customtemplate").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v3/{project_id}/lts/struct/customtemplate
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res []CustomTemplate
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}

type CustomTemplate struct {
	// Project ID.
	ProjectId string `json:"projectId"`
	// Template name.
	Name string `json:"templateName"`
	// Structuring type. Currently, regular expression, JSON, delimiters, and Nginx are supported.
	Type string `json:"template_type"`
	// Sample log event.
	DemoLog string `json:"demoLog"`
	// Structured field.
	DemoFields []FieldFullResponse `json:"demo_fields"`
	// Keyword details.
	TagFields []FieldResponse `json:"tag_fields"`
	// Structuring method.
	Rule *RuleResponse `json:"rule"`
	// Attributes of the sample log event.
	DemoLabel string `json:"demoLabel"`
	// Template creation/update time.
	CreatedAt int64 `json:"create_time"`
	// Structuring rule ID.
	ID string `json:"id"`
}

type FieldFullResponse struct {
	// Field name.
	Name string `json:"fieldName"`
	// Field content.
	Content string `json:"content"`
	// Field data type.
	Type string `json:"type"`
	// Whether parsing is enabled.
	IsAnalysis bool `json:"isAnalysis"`
	// Field sequence number.
	Index int `json:"index"`
	// Describes the hierarchical relationship between fields in a multi-level JSON file.
	Relation string `json:"relation"`
	// Custom field alias in JSON and Nginx modes.
	UserDefinedName string `json:"user_defined_name"`
}
