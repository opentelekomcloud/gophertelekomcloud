package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateComplianceOpts struct {
	DomainId             string         `json:"-"`
	PolicyResource       PolicyResource `json:"policy_resource" required:"true"`
	TriggerType          string         `json:"trigger_type" required:"true"`
	ComplianceState      string         `json:"compliance_state" required:"true"`
	PolicyAssignmentID   string         `json:"policy_assignment_id" required:"true"`
	PolicyAssignmentName string         `json:"policy_assignment_name,omitempty"`
	EvaluationTime       int64          `json:"evaluation_time"`
	EvaluationHash       string         `json:"evaluation_hash"`
}

type PolicyResource struct {
	ResourceID       string `json:"resource_id,omitempty"`
	ResourceName     string `json:"resource_name,omitempty"`
	ResourceProvider string `json:"resource_provider,omitempty"`
	ResourceType     string `json:"resource_type,omitempty"`
	RegionID         string `json:"region_id,omitempty"`
	DomainID         string `json:"domain_id,omitempty"`
}

func UpdateCompliance(client *golangsdk.ServiceClient, opts UpdateComplianceOpts) (*PolicyState, error) {
	// PUT /v1/resource-manager/domains/{domain_id}/policy-states
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("resource-manager", "domains", opts.DomainId, "policy-states"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PolicyState

	err = extract.Into(raw.Body, &res)
	return &res, err
}
