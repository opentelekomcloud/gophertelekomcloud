package policies

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type policyResult struct {
	golangsdk.Result
}

type CreateResult struct {
	policyResult
}

type GetResult struct {
	policyResult
}

type UpdateResult struct {
	policyResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type PolicyTriggerPropertiesResp struct {
	Pattern   []string `json:"pattern"`
	StartTime string   `json:"start_time"`
}
type PolicyTriggerResp struct {
	Properties PolicyTriggerPropertiesResp `json:"properties"`
}
type PolicyAssociateVault struct {
	VaultID            string `json:"vault_id"`
	DestinationVaultID string `json:"destination_vault_id"`
}

type Policy struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Enabled             bool                   `json:"enabled"`
	OperationDefinition *PolicyODCreate        `json:"operation_definition"`
	OperationType       OperationType          `json:"operation_type"`
	Trigger             *PolicyTriggerResp     `json:"trigger"`
	AssociatedVaults    []PolicyAssociateVault `json:"associated_vaults"`
}

func (r policyResult) Extract() (*Policy, error) {
	var s struct {
		Policy *Policy `json:"policy"`
	}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("error extracting policy from create response: %s", err)
	}
	return s.Policy, err
}

type PolicyPage struct {
	pagination.SinglePageBase
}

func ExtractPolicies(r pagination.Page) ([]Policy, error) {
	var s []Policy
	err := r.(PolicyPage).Result.ExtractIntoSlicePtr(&s, "policies")
	if err != nil {
		return nil, err
	}
	return s, nil
}
