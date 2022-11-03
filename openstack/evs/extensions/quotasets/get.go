package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, projectID string) (*QuotaSet, error) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", projectID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		QuotaSet QuotaSet `json:"quota_set"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.QuotaSet, err
}

func GetDefaults(client *golangsdk.ServiceClient, projectID string) (*QuotaSet, error) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", projectID, "defaults"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotaSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}

type QuotaSet struct {
	// ID is project associated with this QuotaSet.
	ID string `json:"id"`
	// Volumes is the number of volumes that are allowed for each project.
	Volumes int `json:"volumes"`
	// Snapshots is the number of snapshots that are allowed for each project.
	Snapshots int `json:"snapshots"`
	// Gigabytes is the size (GB) of volumes and snapshots that are allowed for
	// each project.
	Gigabytes int `json:"gigabytes"`
	// PerVolumeGigabytes is the size (GB) of volumes and snapshots that are
	// allowed for each project and the specifed volume type.
	PerVolumeGigabytes int `json:"per_volume_gigabytes"`
	// Backups is the number of backups that are allowed for each project.
	Backups int `json:"backups"`
	// BackupGigabytes is the size (GB) of backups that are allowed for each
	// project.
	BackupGigabytes int `json:"backup_gigabytes"`
}
