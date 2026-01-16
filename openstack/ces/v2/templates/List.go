package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the options for querying alarm templates.
type ListOpts struct {
	// Specifies the pagination offset. Default: 0, Range: 0-10000
	Offset int `q:"offset,omitempty"`
	// Specifies the number of records on each page. Default: 100, Range: 1-100
	Limit int `q:"limit,omitempty"`
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	Namespace string `q:"namespace"`
	// Specifies the resource dimension.
	// Multiple values can be separated by commas.
	DimName string `q:"dim_name"`
	// Specifies the alarm template type.
	// Possible values: system, custom, system_event, custom_event, system_custom_event
	TemplateType string `q:"template_type"`
	// Specifies the alarm template name. Fuzzy matching is supported.
	// Enter 1 to 128 characters.
	TemplateName string `q:"template_name"`
}

// ListResponse contains the response from the List request.
type ListResponse struct {
	// Specifies the list of alarm templates.
	AlarmTemplates []Template `json:"alarm_templates"`
	// Specifies the total number of alarm templates.
	Count int `json:"count"`
}

// Template represents an alarm template in the list response.
type Template struct {
	// Specifies the alarm template ID.
	TemplateId string `json:"template_id"`
	// Specifies the alarm template name.
	TemplateName string `json:"template_name"`
	// Specifies the alarm template type.
	// Possible values: system, custom
	TemplateType string `json:"template_type"`
	// Specifies the time when the alarm template was created.
	// The value is in UTC format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	CreateTime string `json:"create_time"`
	// Provides supplementary information about the alarm template.
	TemplateDescription string `json:"template_description"`
}

// List returns a list of alarm templates.
func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("alarm-templates").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarm-templates
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
