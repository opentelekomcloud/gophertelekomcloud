package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetRuleStatus(client *golangsdk.ServiceClient, domainId, id string) (*PolicyStatus, error) {
	// GET /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}/policy-states/evaluation-state
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "domains", domainId, "policy-assignments", id, "policy-states", "evaluation-state"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res PolicyStatus
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PolicyStatus struct {
	PolicyAssignmentId string `json:"policy_assignment_id"`
	State              string `json:"state"`
	StartTime          string `json:"start_time"`
	EnTime             string `json:"end_time"`
	ErrorMessage       string `json:"error_message"`
}
