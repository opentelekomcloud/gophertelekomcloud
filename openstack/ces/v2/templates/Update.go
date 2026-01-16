package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the options for modifying a custom alarm template.
type UpdateOpts struct {
	// Specifies the alarm template name.
	// It must start with a letter and can contain letters, digits, underscores (_),
	// hyphens (-), parentheses, and periods (.).
	// Enter 1 to 128 characters.
	TemplateName string `json:"template_name" required:"true"`
	// Specifies the alarm template type.
	// 0: metric alarm template
	// 2: event alarm template
	TemplateType int `json:"template_type,omitempty"`
	// Provides supplementary information about the alarm template.
	// Enter 0 to 256 characters.
	TemplateDescription string `json:"template_description,omitempty"`
	// Specifies the alarm policies. A maximum of 50 policies are supported.
	Policies []Policy `json:"policies" required:"true"`
}

// Update modifies a custom alarm template.
func Update(client *golangsdk.ServiceClient, templateId string, opts UpdateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /v2/{project_id}/alarm-templates/{template_id}
	_, err = client.Put(client.ServiceURL("alarm-templates", templateId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
