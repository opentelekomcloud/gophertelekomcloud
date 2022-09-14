package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOpts contains the values used when updating a backup policy.
type UpdateOpts struct {
	Description         string                       `json:"description,omitempty"`
	Name                string                       `json:"name,omitempty"`
	Parameters          PolicyParam                  `json:"parameters,omitempty"`
	Resources           []Resource                   `json:"resources,omitempty"`
	ScheduledOperations []ScheduledOperationToUpdate `json:"scheduled_operations,omitempty"`
}

// Update allows backup policies to be updated. call the Extract method on the UpdateResult.
func Update(c *golangsdk.ServiceClient, policyId string, opts UpdateOpts) (*CreateBackupPolicy, error) {
	b, err := golangsdk.BuildRequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	raw, err := c.Put(c.ServiceURL("policies", policyId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
