package backup

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the attributes you want to see returned. Marker and Limit are used for pagination.
type ListOpts struct {
	Status       string `q:"status"`
	Limit        string `q:"limit"`
	Marker       string `q:"marker"`
	Sort         string `q:"sort"`
	AllTenants   string `q:"all_tenants"`
	Name         string `q:"name"`
	ResourceId   string `q:"resource_id"`
	ResourceName string `q:"resource_name"`
	PolicyId     string `q:"policy_id"`
	VmIp         string `q:"ip"`
	CheckpointId string `q:"checkpoint_id"`
	ID           string
	ResourceType string `q:"resource_type"`
}

func FilterBackupsById(backups []Backup, filterId string) ([]Backup, error) {
	var refinedBackups []Backup

	for _, backup := range backups {

		if filterId == backup.Id {
			refinedBackups = append(refinedBackups, backup)
		}
	}

	return refinedBackups, nil
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToBackupCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains the options for create a Backup. This object is
// passed to backup.Create().
type CreateOpts struct {
	BackupName   string             `json:"backup_name,omitempty"`
	Description  string             `json:"description,omitempty"`
	ResourceType string             `json:"resource_type,omitempty"`
	Incremental  *bool              `json:"incremental,omitempty"`
	Tags         []tags.ResourceTag `json:"tags,omitempty"`
	ExtraInfo    interface{}        `json:"extra_info,omitempty"`
}

// ToBackupCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToBackupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "protect")
}

// ResourceBackupCapabilityOptsBuilder allows extensions to add additional parameters to the
// QueryResourceBackupCapability request.
type ResourceBackupCapabilityOptsBuilder interface {
	ToQueryResourceCreateMap() (map[string]interface{}, error)
}

// ResourceBackupCapOpts contains the options for querying whether resources can be backed up. This object is
// passed to backup.QueryResourceBackupCapability().
type ResourceBackupCapOpts struct {
	CheckProtectable []ResourceCapQueryParams `json:"check_protectable" required:"true"`
}

type ResourceCapQueryParams struct {
	ResourceId   string `json:"resource_id" required:"true"`
	ResourceType string `json:"resource_type" required:"true"`
}

// ToQueryResourceCreateMap assembles a request body based on the contents of a
// ResourceBackupCapOpts.
func (opts ResourceBackupCapOpts) ToQueryResourceCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// QueryResourceBackupCapability will query whether resources can be backed up based on the values in ResourceBackupCapOpts. To extract
// the ResourceCap object from the response, call the ExtractQueryResponse method on the
// QueryResult.
func QueryResourceBackupCapability(client *golangsdk.ServiceClient, opts ResourceBackupCapabilityOptsBuilder) (r QueryResult) {
	b, err := opts.ToQueryResourceCreateMap()
	if err != nil {
		r.Err = err

		return
	}
	_, r.Err = client.Post(client.ServiceURL("providers", providerID, "resources", "action"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
