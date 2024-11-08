package lifecycle

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	// Indicates the name of an instance.
	// An instance name starts with a letter,
	// consists of 4 to 64 characters,
	// and supports only letters, digits, and hyphens (-).
	Name string `json:"name,omitempty"`

	// Indicates the description of an instance.
	// It is a character string containing not more than 1024 characters.
	Description *string `json:"description,omitempty"`

	// Indicates the time at which a maintenance time window starts.
	// Format: HH:mm:ss
	MaintainBegin string `json:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm:ss
	MaintainEnd string `json:"maintain_end,omitempty"`

	// Indicates the ID of a security group.
	SecurityGroupID string `json:"security_group_id,omitempty"`

	// Indicates the action to be taken when the memory usage reaches the disk capacity threshold. Options:
	// time_base: Automatically delete the earliest messages.
	// produce_reject: Stop producing new messages.
	RetentionPolicy string `json:"retention_policy,omitempty"`
}

// Update is a method which can be able to update the instance
// via accessing to the service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*InstanceCreate, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", id), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return nil, err
	}

	var res InstanceCreate
	err = extract.Into(raw.Body, &res)
	return &res, err
}
