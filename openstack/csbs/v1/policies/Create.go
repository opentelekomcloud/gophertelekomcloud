package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts contains the options for create a Backup Policy. This object is passed to policies.Create().
type CreateOpts struct {
	// Backup policy description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Backup policy name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`
	// Backup parameters
	// For details, see Table 2-24.
	Parameters PolicyParam `json:"parameters" required:"true"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id" required:"true"`
	// Backup object list. The list can be blank.
	// For details, see Table 2-25.
	Resources []Resource `json:"resources" required:"true"`
	// Scheduling period
	ScheduledOperations []ScheduledOperation `json:"scheduled_operations" required:"true"`
	// Tag list
	// This list cannot be an empty list.
	// The list can contain up to 10 keys.
	// Keys in this list must be unique.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

// Create will create a new backup policy based on the values in CreateOpts. To extract
// the Backup object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*BackupPolicy, error) {
	b, err := build.RequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/policies
	raw, err := client.Post(client.ServiceURL("policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
