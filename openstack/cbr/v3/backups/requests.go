package backups

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}

type ListOptsBuilder interface {
	ToBackupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the flavor attributes you want to see returned.
type ListOpts struct {
	CheckpointID   string `q:"checkpoint_id"`
	DedicatedCloud bool   `q:"dec"`
	EndTime        string `q:"end_time"`
	ImageType      string `q:"image_type"`
	Limit          string `q:"limit"`
	Marker         string `q:"marker"`
	MemberStatus   string `q:"member_status"`
	Name           string `q:"name"`
	Offset         string `q:"offset"`
	OwnType        string `q:"own_type"`
	ParentID       string `q:"parent_id"`
	ResourceAZ     string `q:"resource_az"`
	ResourceID     string `q:"resource_id"`
	ResourceName   string `q:"resource_name"`
	ResourceType   string `q:"resource_type"`
	Sort           string `q:"sort"`
	StartTime      string `q:"start_time"`
	Status         string `q:"status"`
	UserPercent    string `q:"user_percent"`
	VaultID        string `q:"vault_id"`
}

type RestoreBackupStruct struct {
	Mappings []BackupRestoreServer `json:"mappings,omitempty"`
	PowerOn  bool                  `json:"power_on,omitempty"`
	ServerID string                `json:"server_id,omitempty"`
	VolumeID string                `json:"volume_id,omitempty"`
}

type BackupRestoreServer struct {
	BackupID string `json:"backup_id"`
	VolumeID string `json:"volume_id"`
}

func (opts ListOpts) ToBackupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToBackupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupPage{pagination.SinglePageBase(r)}
	})
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(singleURL(client, id), nil)
	return
}

type RestoreResourcesOptsBuilder interface {
	ToRestoreBackup() (map[string]interface{}, error)
}

type RestoreBackupOpts struct {
	Restore RestoreBackupStruct `json:"restore"`
}

func (opts RestoreBackupOpts) ToRestoreBackup() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func RestoreBackup(client *golangsdk.ServiceClient, backupID string, opts RestoreResourcesOptsBuilder) (r RestoreBackupResult) {
	b, err := opts.ToRestoreBackup()
	if err != nil {
		r.Err = fmt.Errorf("failed to restore backup: %s", err)
		return
	}
	_, r.Err = client.Post(restoreURL(client, backupID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
