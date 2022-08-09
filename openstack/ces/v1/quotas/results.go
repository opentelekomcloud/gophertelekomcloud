package quotas

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ShowQuotasResponse struct {
	Quotas Quotas `json:"quotas,omitempty"`
}

type Quotas struct {
	// Specifies the resource quota list.
	Resources []Resource `json:"resources"`
}

type Resource struct {
	// Specifies the quota type.
	Type string `json:"type"`
	// Specifies the used amount of the quota.
	Used int `json:"used"`
	// Specifies the quota unit.
	Unit string `json:"unit"`
	// Specifies the total amount of the quota.
	Quota int `json:"quota"`
}

type ShowQuotasResult struct {
	golangsdk.Result
}

func (r ShowQuotasResult) Extract() (*ShowQuotasResponse, error) {
	var s = ShowQuotasResponse{}
	if r.Err != nil {
		return nil, r.Err
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, fmt.Errorf("failed to extract Show Quotas Response")
	}
	return &s, nil
}
