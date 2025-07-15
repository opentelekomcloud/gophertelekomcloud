package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllUserComplianceOpts struct {
	DomainId string `json:"-"`
	// compliance_state
	ComplianceState string `q:"compliance_state"`
	ResourceId      string `q:"resource_id"`
	ResourceName    string `q:"resource_name"`
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
}

func ListAllUserCompliance(client *golangsdk.ServiceClient, opts ListAllUserComplianceOpts) ([]PolicyState, error) {
	// GET /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}/policy-states

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("resource-manager", "domains", opts.DomainId, "policy-states"),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractComplianceState(pages)
}
