package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Update(client *golangsdk.ServiceClient, projectID string, opts UpdateOpts) (*QuotaSet, error) {
	b, err := build.RequestBody(opts, "quota_set")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-quota-sets", projectID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res QuotaSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}

type UpdateOpts struct {
	// Volumes is the number of volumes that are allowed for each project.
	Volumes *int `json:"volumes,omitempty"`
	// Snapshots is the number of snapshots that are allowed for each project.
	Snapshots *int `json:"snapshots,omitempty"`
	// Gigabytes is the size (GB) of volumes and snapshots that are allowed for
	// each project.
	Gigabytes *int `json:"gigabytes,omitempty"`
	// PerVolumeGigabytes is the size (GB) of volumes and snapshots that are
	// allowed for each project and the specifed volume type.
	PerVolumeGigabytes *int `json:"per_volume_gigabytes,omitempty"`
	// Backups is the number of backups that are allowed for each project.
	Backups *int `json:"backups,omitempty"`
	// BackupGigabytes is the size (GB) of backups that are allowed for each
	// project.
	BackupGigabytes *int `json:"backup_gigabytes,omitempty"`
	// Groups is the number of groups that are allowed for each project.
	Groups *int `json:"groups,omitempty"`
	// Force will update the quotaset even if the quota has already been used
	// and the reserved quota exceeds the new quota.
	Force bool `json:"force,omitempty"`
}
