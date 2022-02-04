package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// UpdateResult represents the result of a update operation.
type UpdateResult struct {
	golangsdk.ErrResult
}

type BackupStatus string

const (
	StatusBuilding  BackupStatus = "BUILDING"
	StatusCompleted BackupStatus = "COMPLETED"
	StatusFailed    BackupStatus = "FAILED"
	StatusDeleting  BackupStatus = "DELETING"
	StatusDeleted   BackupStatus = "DELETED"
)

type Backup struct {
	ID         string              `json:"id"`
	InstanceID string              `json:"instance_id"`
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	Size       int                 `json:"size"`
	Databases  []BackupDatabase    `json:"databases"`
	BeginTime  string              `json:"begin_time"`
	EndTime    string              `json:"end_time"`
	Datastore  instances.Datastore `json:"datastore"`
	Status     BackupStatus        `json:"status"`
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
