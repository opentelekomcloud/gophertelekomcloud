package quotasets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, projectID string) (*QuotaSet, error) {
	// GET /v3/{project_id}/os-quota-sets/{target_project_id}?usage=True
	raw, err := client.Get(client.ServiceURL("os-quota-sets", projectID)+"?usage=True", nil, nil)
	if err != nil {
		return nil, err
	}

	var res QuotaSet
	err = extract.IntoStructPtr(raw.Body, &res, "quota_set")
	return &res, err
}

type Detail struct {
	Reserved int `json:"reserved"`
	Limit    int `json:"limit"`
	InUse    int `json:"in_use"`
}

type QuotaSet struct {
	// Specifies the size (GB) reserved for common I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	GigabytesSAS Detail `json:"gigabytes_SAS"`
	// Specifies the number of reserved common I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	VolumesSATA Detail `json:"volumes_SATA"`
	// Specifies the total size (GB) of disks and snapshots allowed. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	Gigabytes Detail `json:"gigabytes"`
	// Specifies the backup size (GB). Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	BackupGigabytes Detail `json:"backup_gigabytes"`
	// Specifies the number of snapshots reserved for high I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	SnapshotsSAS Detail `json:"snapshots_SAS"`
	// Specifies the number of reserved ultra-high I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	VolumesSSD Detail `json:"volumes_SSD"`
	// Specifies the number of snapshots. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	Snapshots Detail `json:"snapshots"`
	// Specifies the tenant ID. The tenant ID is actually the project ID.
	Id string `json:"id"`
	// Specifies the number of reserved high I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	VolumesSAS Detail `json:"volumes_SAS"`
	// Specifies the number of snapshots reserved for ultra-high I/O disks. Sub-parameters include reserved
	// (reserved quota), limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	SnapshotsSSD Detail `json:"snapshots_SSD"`
	// Specifies the number of disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	Volumes Detail `json:"volumes"`
	// Specifies the backup size (GB). Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	GigabytesSATA Detail `json:"gigabytes_SATA"`
	// Specifies the number of backups. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	Backups Detail `json:"backups"`
	// Specifies the size (GB) reserved for ultra-high I/O disks. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	GigabytesSSD Detail `json:"gigabytes_SSD"`
	// Specifies the number of snapshots reserved for common I/O disks. Sub-parameters include reserved
	// (reserved quota), limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	SnapshotsSATA Detail `json:"snapshots_SATA"`
	// Specifies the capacity quota of each EVS disk. Sub-parameters include reserved (reserved quota),
	// limit (maximum quota), and in_use (used quota), and are made up of key-value pairs.
	PerVolumeGigabytes Detail `json:"per_volume_gigabytes"`
}
