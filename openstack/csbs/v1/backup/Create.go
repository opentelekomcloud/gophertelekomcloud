package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts contains the options for create a Backup. This object is passed to backup.Create().
type CreateOpts struct {
	BackupName   string             `json:"backup_name,omitempty"`
	Description  string             `json:"description,omitempty"`
	ResourceType string             `json:"resource_type,omitempty"`
	Incremental  *bool              `json:"incremental,omitempty"`
	Tags         []tags.ResourceTag `json:"tags,omitempty"`
	ExtraInfo    interface{}        `json:"extra_info,omitempty"`
}

// Create will create a new backup based on the values in CreateOpts. To extract
// the checkpoint object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, resourceID string, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "protect")
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(client.ServiceURL("providers", providerID, "resources", resourceID, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
