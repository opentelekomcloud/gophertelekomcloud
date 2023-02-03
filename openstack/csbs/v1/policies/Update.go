package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contains the values used when updating a backup policy.
type UpdateOpts struct {
	// Backup policy description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Backup policy name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`
	// Backup objects
	Resources []Resource `json:"resources,omitempty"`
	// Scheduling period. A backup policy has only one backup period.
	ScheduledOperations []ScheduledOperationToUpdate `json:"scheduled_operations,omitempty"`
}

type Resource struct {
	// ID of the object to be backed up
	Id string `json:"id" required:"true"`
	// Entity object type of backup objects
	// The value is fixed at OS::Nova::Server (ECSs).
	Type string `json:"type" required:"true"`
	// Backup object name
	Name string `json:"name" required:"true"`
	// Additional information about the backup object
	ExtraInfo interface{} `json:"extra_info,omitempty"`
}

type ScheduledOperationToUpdate struct {
	// Scheduling period description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Whether the backup policy is enabled
	// The default value is true. If it is set to false, automatic scheduling is disabled but manual scheduling is supported.
	Enabled bool `json:"enabled"`
	// Scheduling period name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name,omitempty"`
	// Scheduling period parameter
	OperationDefinition OperationDefinition `json:"operation_definition,omitempty"`
	// Scheduling policy
	Trigger Trigger `json:"trigger,omitempty"`
	// Scheduling period ID
	Id string `json:"id" required:"true"`
}

// Update allows backup policies to be updated. call the Extract method on the UpdateResult.
func Update(c *golangsdk.ServiceClient, policyId string, opts UpdateOpts) (*BackupPolicyResponse, error) {
	b, err := build.RequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	// PUT https://{endpoint}/v1/{project_id}/policies/{policy_id}
	raw, err := c.Put(c.ServiceURL("policies", policyId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
