package compliance

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func ListAllPolicies(client *golangsdk.ServiceClient) ([]PolicyDefinition, error) {
	// GET /v1/resource-manager/policy-definitions

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("resource-manager", "policy-definitions"),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractResources(pages)
}

type ResPage struct {
	pagination.NewSinglePageBase
}

func ExtractResources(r pagination.NewPage) ([]PolicyDefinition, error) {
	var s struct {
		Values []PolicyDefinition `json:"value"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Values, err
}

type PolicyDefinition struct {
	ID                   string                               `json:"id"`
	Name                 string                               `json:"name"`
	DisplayName          string                               `json:"display_name"`
	PolicyType           string                               `json:"policy_type"`
	Description          string                               `json:"description"`
	PolicyRuleType       string                               `json:"policy_rule_type"`
	PolicyRule           interface{}                          `json:"policy_rule"`
	TriggerType          string                               `json:"trigger_type"`
	Keywords             []string                             `json:"keywords"`
	DefaultResourceTypes []DefaultResourceType                `json:"default_resource_types"`
	Parameters           map[string]PolicyParameterDefinition `json:"parameters"`
}

type DefaultResourceType struct {
	Provider string `json:"provider"`
	Type     string `json:"type"`
}

type PolicyParameterDefinition struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	AllowedValues []interface{} `json:"allowed_values"`
	DefaultValue  string        `json:"default_value"`
	Minimum       float64       `json:"minimum"`
	Maximum       float64       `json:"maximum"`
	MinItems      int           `json:"min_items"`
	MaxItems      int           `json:"max_items"`
	MinLength     int           `json:"min_length"`
	MaxLength     int           `json:"max_length"`
	Pattern       string        `json:"pattern"`
	Type          string        `json:"type"`
}
