package compliance

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListResComplianceOpts struct {
	DomainId   string `json:"-"`
	ResourceId string `json:"-"`
	// compliance_state
	ComplianceState string `q:"compliance_state"`
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
}

func ListResCompliance(client *golangsdk.ServiceClient, opts ListResComplianceOpts) ([]PolicyState, error) {
	// GET /v1/resource-manager/domains/{domain_id}/resources/{resource_id}/policy-states

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("resource-manager", "domains", opts.DomainId, "resources", opts.ResourceId, "policy-states"),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractComplianceState(pages)
}

func ExtractComplianceState(r pagination.NewPage) ([]PolicyState, error) {
	var s struct {
		Values []PolicyState `json:"value"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Values, err
}

type PolicyState struct {
	DomainID             string `json:"domain_id"`
	RegionID             string `json:"region_id"`
	ResourceID           string `json:"resource_id"`
	ResourceName         string `json:"resource_name"`
	ResourceProvider     string `json:"resource_provider"`
	ResourceType         string `json:"resource_type"`
	TriggerType          string `json:"trigger_type"`
	ComplianceState      string `json:"compliance_state"`
	PolicyAssignmentID   string `json:"policy_assignment_id"`
	PolicyAssignmentName string `json:"policy_assignment_name"`
	PolicyDefinitionID   string `json:"policy_definition_id"`
	EvaluationTime       string `json:"evaluation_time"`
}
