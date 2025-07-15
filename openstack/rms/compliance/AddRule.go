package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type AddRuleOpts struct {
	DomainId             string                     `json:"-"`
	PolicyAssignmentType string                     `json:"policy_assignment_type,omitempty"`
	Name                 string                     `json:"name" required:"true"`
	Description          string                     `json:"description,omitempty"`
	Period               string                     `json:"period,omitempty"`
	PolicyFilter         PolicyFilterDefinition     `json:"policy_filter,omitempty"`
	PolicyDefinitionID   string                     `json:"policy_definition_id,omitempty"`
	CustomPolicy         *CustomPolicy              `json:"custom_policy,omitempty"`
	Parameters           map[string]PolicyParameter `json:"parameters,omitempty"`
	Tags                 []tags.ResourceTag         `json:"tags,omitempty"`
}

type PolicyFilterDefinition struct {
	RegionID         string `json:"region_id,omitempty"`
	ResourceProvider string `json:"resource_provider,omitempty"`
	ResourceType     string `json:"resource_type,omitempty"`
	ResourceID       string `json:"resource_id,omitempty"`
	TagKey           string `json:"tag_key,omitempty"`
	TagValue         string `json:"tag_value,omitempty"`
}

type CustomPolicy struct {
	FunctionUrn string                 `json:"function_urn" required:"true"`
	AuthType    string                 `json:"auth_type" required:"true"`
	AuthValue   map[string]interface{} `json:"auth_value,omitempty"`
}

type PolicyParameter struct {
	Value interface{} `json:"value,omitempty"`
}

func AddRule(client *golangsdk.ServiceClient, opts AddRuleOpts) (*PolicyRule, error) {
	// PUT /v1/resource-manager/domains/{domain_id}/policy-assignments
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("resource-manager", "domains", opts.DomainId, "policy-assignments"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PolicyRule

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PolicyRule struct {
	PolicyAssignmentType string                     `json:"policy_assignment_type"`
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Description          string                     `json:"description"`
	PolicyFilter         *PolicyFilterDefinition    `json:"policy_filter,omitempty"`
	Period               string                     `json:"period"`
	State                string                     `json:"state"`
	Created              string                     `json:"created"`
	Updated              string                     `json:"updated"`
	PolicyDefinitionID   string                     `json:"policy_definition_id"`
	CustomPolicy         *CustomPolicy              `json:"custom_policy,omitempty"`
	Parameters           map[string]PolicyParameter `json:"parameters"`
	Tags                 []ResourceTag              `json:"tags"`
	CreatedBy            string                     `json:"created_by"`
	TargetType           string                     `json:"target_type"`
	TargetID             string                     `json:"target_id"`
}

type ResourceTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
