package checkpoint

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOptsBuilder interface {
	ToCheckpointCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// ID of the vault
	VaultID string `json:"vault_id" required:"true"`
	// Checkpoint parameters
	Parameters CheckpointParam `json:"parameters,omitempty"`
}

type Resource struct {
	// ID of the resource to be backed up
	ID string `json:"id"`
	// Name of the resource to be backed up
	Name string `json:"name,omitempty"`
	// Type of the resource to be backed up
	// OS::Nova::Server | OS::Cinder::Volume
	Type string `json:"type,omitempty"`
}

type CheckpointParam struct {
	// Describes whether automatic triggering is enabled
	// Default: false
	AutoTrigger bool `json:"auto_trigger,omitempty"`
	// Backup description
	Description string `json:"description,omitempty"`
	// Whether bacup is incremental or not
	// Default: true
	Incremental bool `json:"incremental,omitempty"`
	// Backup name
	Name string `json:"name,omitempty"`
	// UUID list of resources to be backed up
	Resources []string `json:"resources,omitempty"`
	// Additional information on Resource
	ResourceDetails []Resource `json:"resource_details,omitempty"`
}

func (opts CreateOpts) ToCheckpointCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "checkpoint")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToCheckpointCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rootURL(client), reqBody, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(checkpointURL(client, id), &r.Body, nil)
	return
}
