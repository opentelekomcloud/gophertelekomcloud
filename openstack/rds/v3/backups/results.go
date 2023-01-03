package backups

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
)

type BackupStatus string

const (
	StatusBuilding  BackupStatus = "BUILDING"
	StatusCompleted BackupStatus = "COMPLETED"
	StatusFailed    BackupStatus = "FAILED"
	StatusDeleting  BackupStatus = "DELETING"
	StatusDeleted   BackupStatus = "DELETED"
)

type Backup struct {
	// Indicates the backup ID.
	ID string `json:"id"`
	// Indicates the DB instance ID.
	InstanceID string `json:"instance_id"`
	// Indicates the backup name.
	Name string `json:"name"`
	// Indicates the backup description.
	Description string `json:"description"`
	// Indicates the backup type. Value:
	//
	// auto: automated full backup
	// manual: manual full backup
	// fragment: differential full backup
	// incremental: automated incremental backup
	Type string `json:"type"`
	// Indicates the backup size in kB.
	Size int `json:"size"`
	// Indicates a list of self-built Microsoft SQL Server databases that are partially backed up. (Only Microsoft SQL Server support partial backups.)
	Databases []BackupDatabase `json:"databases"`
	// Indicates the backup start time in the "yyyy-mm-ddThh:mm:ssZ" format, where "T" indicates the start time of the time field, and "Z" indicates the time zone offset.
	BeginTime string `json:"begin_time"`
	// Indicates the backup end time.
	// In a full backup, it indicates the full backup end time.
	// In a MySQL incremental backup, it indicates the time when the last transaction in the backup file is submitted.
	// The format is yyyy-mm-ddThh:mm:ssZ. T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	EndTime string `json:"end_time"`
	// Indicates the database version.
	Datastore instances.Datastore `json:"datastore"`
	// Indicates the backup status. Value:
	//
	// BUILDING: Backup in progress
	// COMPLETED: Backup completed
	// FAILED: Backup failed
	// DELETING: Backup being deleted
	Status BackupStatus `json:"status"`
}

func WaitForBackup(c *golangsdk.ServiceClient, instanceID, backupID string, status BackupStatus) error {
	return golangsdk.WaitFor(1200, func() (bool, error) {
		backupList, err := List(c, ListOpts{InstanceID: instanceID, BackupID: backupID})
		if err != nil {
			return false, fmt.Errorf("error extracting backups: %w", err)
		}
		if len(backupList) == 0 {
			if status == StatusDeleted { // when deleted, backup is actually always in status "DELETING"
				return true, nil
			}
			return false, fmt.Errorf("backup %s/%s does not exist", instanceID, backupID)
		}
		backup := backupList[0]
		return backup.Status == status, nil
	})
}
