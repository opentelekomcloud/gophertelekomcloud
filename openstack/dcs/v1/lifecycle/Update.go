package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	// DCS instance name.
	// An instance name is a string of 4–64 characters
	// that contain letters, digits, underscores (_), and hyphens (-).
	// An instance name must start with letters.
	Name string `json:"name,omitempty"`
	// Brief description of the DCS instance.
	// A brief description supports up to 1024 characters.
	Description string `json:"description,omitempty"`
	// Backup policy.
	// This parameter is available for master/standby DCS instances.
	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty"`
	// Time at which the maintenance time window starts.
	// Format: HH:mm:ss
	MaintainBegin string `json:"maintain_begin,omitempty"`
	// Time at which the maintenance time window ends.
	// Format: HH:mm:ss
	MaintainEnd string `json:"maintain_end,omitempty"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id,omitempty"`
}

// Update is a method which can be able to update the instance
// via accessing to the service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Put(client.ServiceURL("instances", id), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
