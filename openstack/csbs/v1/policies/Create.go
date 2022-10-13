package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains the options for create a Backup Policy. This object is passed to policies.Create().
type CreateOpts struct {
	Description         string               `json:"description,omitempty"`
	Name                string               `json:"name" required:"true"`
	Parameters          PolicyParam          `json:"parameters" required:"true"`
	ProviderId          string               `json:"provider_id" required:"true"`
	Resources           []Resource           `json:"resources" required:"true"`
	ScheduledOperations []ScheduledOperation `json:"scheduled_operations" required:"true"`
	Tags                []ResourceTag        `json:"tags,omitempty"`
}

// Create will create a new backup policy based on the values in CreateOpts. To extract
// the Backup object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateBackupPolicy, error) {
	b, err := build.RequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
