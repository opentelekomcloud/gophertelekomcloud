package policies

import (
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// Get will get a single backup policy with specific ID. call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, policyId string) (*GetBackupPolicy, error) {
	// GET https://{endpoint}/v1/{project_id}/policies/{policy_id}
	raw, err := client.Get(client.ServiceURL("policies", policyId), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	var res GetBackupPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}

type GetBackupPolicy struct {
	// Creation time, for example, 2017-04-18T01:21:52.701973
	CreatedAt time.Time `json:"-"`
	// Backup policy description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description"`
	// Backup policy ID
	ID string `json:"id"`
	// Backup policy name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`
	// Parameters of a backup policy
	Parameters PolicyParam `json:"parameters"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id"`
	// Backup object list
	Resources []Resource `json:"resources"`
	// Scheduling period list
	ScheduledOperations []ScheduledOperation `json:"scheduled_operations"`
	// Backup policy status
	// disabled: indicates that the backup policy is unavailable.
	// enabled: indicates that the backup policy is available.
	Status string `json:"status"`
	// Tag list
	// Keys in the tag list must be unique.
	Tags []tags.ResourceTag `json:"tags"`
}
