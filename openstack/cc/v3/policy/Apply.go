package policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Apply applies a central network policy.
func Apply(client *golangsdk.ServiceClient, centralNetworkId, policyId string) (*ApplyResp, error) {
	raw, err := client.Post(client.ServiceURL(client.DomainID, "gcn", "central-network", centralNetworkId, "policies", policyId, "apply"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res ApplyResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ApplyResp struct {
	// Request ID.
	RequestId string `json:"request_id"`
	// Central network policy.
	CentralNetworkPolicy CentralNetworkPolicy `json:"central_network_policy"`
	// List of central network policy changes.
	CentralNetworkPolicyChangeSet []ElementChangeEntry `json:"central_network_policy_change_set"`
}
