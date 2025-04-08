package cloud_structuring

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Log group ID.
	LogGroupId string `json:"log_group_id" required:"true"`
	// Log stream ID
	LogStreamId string `json:"log_stream_id" required:"true"`
	// Template ID.
	// When the system template is used, the current attribute can be empty.
	TemplateId *string `json:"template_id,omitempty"`
	// Template name, which cannot be empty and will be verified.
	Name string `json:"template_name" required:"true"`
	// Type of the template. The value can be `built_in` (system templates) or `custom` (custom templates).
	// For details about system template types,
	// see section "Log Search and Analysis" > "Cloud Structuring Parsing" > "Structuring Templates" in the LTS User Guide.
	Type string `json:"template_type" required:"true"`
	// Example field array
	// . You only need to enter the fields whose status is different from that of `is_analysis` in the template.
	DemoFields []Field `json:"demo_fields,omitempty"`
	// Tag field array. You only need to enter the fields whose status is different from that of `is_analysis` in the template.
	TagFields []Field `json:"tag_fields,omitempty"`
	// Indicates whether to enable quick analysis for demo_fields and tag_fields.
	// If this parameter is set to true, quick analysis is enabled for all fields.
	// If this parameter is left blank or set to false, is_analysis in the template is used to determine
	// whether to enable quick analysis.
	QuickAnalysis *bool `json:"quick_analysis,omitempty"`
}

type Field struct {
	// Field name. A log event can be split into multiple fields with customizable names.
	Name string `json:"field_name" required:"true"`
	// Whether quick analysis is enabled.
	IsAnalysis *bool `json:"is_analysis"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v3/{project_id}/lts/struct/template
	_, err = client.Post(client.ServiceURL("lts", "struct", "template"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return err
	}
	return err
}
