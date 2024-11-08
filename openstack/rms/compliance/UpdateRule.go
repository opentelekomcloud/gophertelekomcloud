package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type UpdateRuleOpts struct {
	DomainId             string                     `json:"-"`
	PolicyAssignmentId   string                     `json:"-"`
	PolicyAssignmentType string                     `json:"policy_assignment_type" required:"true"`
	Name                 string                     `json:"name" required:"true"`
	Description          string                     `json:"description,omitempty"`
	Period               string                     `json:"period,omitempty"`
	PolicyFilter         PolicyFilterDefinition     `json:"policy_filter,omitempty"`
	PolicyDefinitionID   string                     `json:"policy_definition_id,omitempty"`
	CustomPolicy         *CustomPolicy              `json:"custom_policy,omitempty"`
	Parameters           map[string]PolicyParameter `json:"parameters,omitempty"`
	Tags                 []tags.ResourceTag         `json:"tags,omitempty"`
}

func UpdateRule(client *golangsdk.ServiceClient, opts UpdateRuleOpts) (*PolicyRule, error) {
	// PUT /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("resource-manager", "domains", opts.DomainId, "policy-assignments", opts.PolicyAssignmentId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PolicyRule

	err = extract.Into(raw.Body, &res)
	return &res, err
}
