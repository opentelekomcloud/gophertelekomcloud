package quotasets

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetUsage(client *golangsdk.ServiceClient, projectID string) (*QuotaUsageSet, error) {
	raw, err := client.Get(fmt.Sprintf("%s?usage=true", client.ServiceURL("os-quota-sets", projectID)), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		QuotaUsageSet QuotaUsageSet `json:"quota_set"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.QuotaUsageSet, err
}

type QuotaUsageSet struct {
	// ID is the project ID associated with this QuotaUsageSet.
	ID string `json:"id"`
	// Volumes is the volume usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Volumes QuotaUsage `json:"volumes"`
	// Snapshots is the snapshot usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Snapshots QuotaUsage `json:"snapshots"`
	// Gigabytes is the size (GB) usage information of volumes and snapshots
	// for this project, including in_use, limit, reserved and allocated
	// attributes. Note: allocated attribute is available only when nested
	// quota is enabled.
	Gigabytes QuotaUsage `json:"gigabytes"`
	// PerVolumeGigabytes is the size (GB) usage information for each volume,
	// including in_use, limit, reserved and allocated attributes. Note:
	// allocated attribute is available only when nested quota is enabled and
	// only limit is meaningful here.
	PerVolumeGigabytes QuotaUsage `json:"per_volume_gigabytes"`
	// Backups is the backup usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Backups QuotaUsage `json:"backups"`
	// BackupGigabytes is the size (GB) usage information of backup for this
	// project, including in_use, limit, reserved and allocated attributes.
	// Note: allocated attribute is available only when nested quota is
	// enabled.
	BackupGigabytes QuotaUsage `json:"backup_gigabytes"`
}

type QuotaUsage struct {
	// InUse is the current number of provisioned resources of the given type.
	InUse int `json:"in_use"`
	// Allocated is the current number of resources of a given type allocated
	// for use.  It is only available when nested quota is enabled.
	Allocated int `json:"allocated"`
	// Reserved is a transitional state when a claim against quota has been made
	// but the resource is not yet fully online.
	Reserved int `json:"reserved"`
	// Limit is the maximum number of a given resource that can be
	// allocated/provisioned.  This is what "quota" usually refers to.
	Limit int `json:"limit"`
}
