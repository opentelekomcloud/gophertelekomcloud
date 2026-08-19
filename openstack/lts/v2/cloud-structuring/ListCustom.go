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

	var res []StructTemplateModel
	if err := extract.IntoSlicePtr(raw.Body, &res, "results"); err != nil {
		return nil, err
	}
	return convertStructTemplateModels(res), nil
}

type StructTemplateModel struct {
	// Project ID.
	ProjectId string `json:"project_id"`
	// Template name.
	Name string `json:"template_name"`
	// Structuring type. Currently, regular expression, JSON, delimiters, and Nginx are supported.
	Type string `json:"template_type"`
	// Sample log event.
	DemoLog string `json:"demo_log"`
	// Structured field.
	DemoFields []DemoField `json:"demo_fields"`
	// Keyword details.
	TagFields []TagFieldNew `json:"tag_fields"`
	// Structuring method.
	Rule *TemplateRule `json:"rule"`
	// Attributes of the sample log event.
	DemoLabel string `json:"demo_label"`
	// Template creation/update time.
	CreatedAt int64 `json:"create_time"`
	// Structuring rule ID.
	ID string `json:"id"`
}

type CustomTemplate struct {
	ProjectId  string              `json:"projectId"`
	Name       string              `json:"templateName"`
	Type       string              `json:"template_type"`
	DemoLog    string              `json:"demoLog"`
	DemoFields []FieldFullResponse `json:"demo_fields"`
	TagFields  []FieldResponse     `json:"tag_fields"`
	Rule       *RuleResponse       `json:"rule"`
	DemoLabel  string              `json:"demoLabel"`
	CreatedAt  int64               `json:"create_time"`
	ID         string              `json:"id"`
}

type DemoField struct {
	// Field name.
	Name string `json:"field_name"`
	// Field content.
	Content string `json:"content"`
	// Field data type.
	Type string `json:"type"`
	// Whether parsing is enabled.
	IsAnalysis bool `json:"is_analysis"`
	// Field sequence number.
	Index int `json:"index"`
	// Describes the hierarchical relationship between fields in a multi-level JSON file.
	Relation string `json:"relation"`
	// Custom field alias in JSON and Nginx modes.
	UserDefinedName string `json:"user_defined_name"`
}

type FieldFullResponse struct {
	Name            string `json:"fieldName"`
	Content         string `json:"content"`
	Type            string `json:"type"`
	IsAnalysis      bool   `json:"isAnalysis"`
	Index           int    `json:"index"`
	Relation        string `json:"relation"`
	UserDefinedName string `json:"user_defined_name"`
}

type TagFieldNew struct {
	// Field name.
	Name string `json:"field_name"`
	// Field content.
	Content string `json:"content"`
	// Field data type.
	Type string `json:"type"`
	// Whether parsing is enabled.
	IsAnalysis bool `json:"is_analysis"`
	// Field sequence number.
	Index int `json:"index"`
}

type TemplateRule struct {
	// Structuring type.
	Type string `json:"type"`
	// Type-specific structuring rule.
	Param string `json:"param"`
}

func convertStructTemplateModels(models []StructTemplateModel) []CustomTemplate {
	if models == nil {
		return nil
	}
	result := make([]CustomTemplate, len(models))
	for i, model := range models {
		result[i] = CustomTemplate{
			ProjectId:  model.ProjectId,
			Name:       model.Name,
			Type:       model.Type,
			DemoLog:    model.DemoLog,
			DemoFields: convertDemoFields(model.DemoFields),
			TagFields:  convertTagFields(model.TagFields),
			Rule:       convertTemplateRule(model.Rule),
			DemoLabel:  model.DemoLabel,
			CreatedAt:  model.CreatedAt,
			ID:         model.ID,
		}
	}
	return result
}

func convertDemoFields(fields []DemoField) []FieldFullResponse {
	if fields == nil {
		return nil
	}
	result := make([]FieldFullResponse, len(fields))
	for i, field := range fields {
		result[i] = FieldFullResponse{
			Name:            field.Name,
			Content:         field.Content,
			Type:            field.Type,
			IsAnalysis:      field.IsAnalysis,
			Index:           field.Index,
			Relation:        field.Relation,
			UserDefinedName: field.UserDefinedName,
		}
	}
	return result
}

func convertTagFields(fields []TagFieldNew) []FieldResponse {
	if fields == nil {
		return nil
	}
	result := make([]FieldResponse, len(fields))
	for i, field := range fields {
		result[i] = FieldResponse{
			Name:       field.Name,
			Content:    field.Content,
			Type:       field.Type,
			IsAnalysis: field.IsAnalysis,
			Index:      field.Index,
		}
	}
	return result
}

func convertTemplateRule(rule *TemplateRule) *RuleResponse {
	if rule == nil {
		return nil
	}
	return &RuleResponse{
		Type:  rule.Type,
		Param: rule.Param,
	}
}
