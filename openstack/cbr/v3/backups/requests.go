package backups

import (
	"reflect"

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
	ID           string
	CheckpointID string `q:"checkpoint_id"`
	ImageType    string `q:"image_type"`
	Limit        string `q:"limit"`
	Marker       string `q:"marker"`
	Name         string `q:"name"`
	Offset       string `q:"offset"`
	ParentID     string `q:"parent_id"`
	ResourceAZ   string `q:"resource_az"`
	ResourceID   string `q:"resource_id"`
	ResourceName string `q:"resource_name"`
	ResourceType string `q:"resource_type"`
	Sort         string `q:"sort"`
	Status       string `q:"status"`
	VaultID      string `q:"vault_id"`
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

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	url := listURL(client)

	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	url += q.String()

	pages, err := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()
	if err != nil {
		return nil, err
	}

	allBackups, err := ExtractBackups(pages)
	if err != nil {
		return nil, err
	}

	return FilterBackups(allBackups, opts)
}

func FilterBackups(backups []Backup, opts ListOpts) ([]Backup, error) {
	var refinedBackups []Backup
	var matched bool
	m := map[string]interface{}{}
	if opts.ID != "" {
		m["ID"] = opts.ID
	}

	if len(m) > 0 && len(backups) > 0 {
		for _, backup := range backups {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&backup, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedBackups = append(refinedBackups, backup)
			}
		}

	} else {
		refinedBackups = backups
	}

	return refinedBackups, nil
}

func getStructField(v *Backup, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
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
		r.Err = err
		return
	}
	_, r.Err = client.Post(restoreURL(client, backupID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
