package backups

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
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
	//
	Size int `json:"size"`
	//
	// Indicates a list of self-built Microsoft SQL Server databases that are partially backed up. (Only Microsoft SQL Server support partial backups.)
	Databases []BackupDatabase `json:"databases"`
	// Indicates the backup start time in the "yyyy-mm-ddThh:mm:ssZ" format, where "T" indicates the start time of the time field, and "Z" indicates the time zone offset.
	BeginTime string `json:"begin_time"`
	//
	EndTime string `json:"end_time"`
	//
	Datastore instances.Datastore `json:"datastore"`
	// Indicates the backup status. Value:
	//
	// BUILDING: Backup in progress
	// COMPLETED: Backup completed
	// FAILED: Backup failed
	// DELETING: Backup being deleted
	Status BackupStatus `json:"status"`
}

type CreateResult struct {
	golangsdk.Result
}

type BackupPage struct {
	pagination.SinglePageBase
}

func (p BackupPage) IsEmpty() (bool, error) {
	bs, err := ExtractBackups(p)
	return len(bs) == 0, err
}

func ExtractBackups(r pagination.Page) ([]Backup, error) {
	var bks []Backup
	err := r.(BackupPage).ExtractIntoSlicePtr(&bks, "backups")
	if err != nil {
		return nil, err
	}
	return bks, nil
}

func (r CreateResult) Extract() (*Backup, error) {
	backup := new(Backup)
	err := r.ExtractIntoStructPtr(backup, "backup")
	if err != nil {
		return nil, err
	}
	return backup, nil
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type RestoreResult struct {
	instances.CreateResult
}

func WaitForBackup(c *golangsdk.ServiceClient, instanceID, backupID string, status BackupStatus) error {
	return golangsdk.WaitFor(1200, func() (bool, error) {
		pages, err := List(c, ListOpts{InstanceID: instanceID, BackupID: backupID}).AllPages()
		if err != nil {
			return false, fmt.Errorf("error listing backups: %w", err)
		}
		backupList, err := ExtractBackups(pages)
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
