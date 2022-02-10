package backups

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToBackupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a Backup.
type UpdateOpts struct {
	// Keep Days
	KeepDays *int `json:"keep_days" required:"true"`
	// Start Time
	StartTime string `json:"start_time,omitempty"`
	// Period
	Period string `json:"period,omitempty"`
}

// ToBackupUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToBackupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "backup_policy")
}

// Update accepts a UpdateOpts struct and uses the values to update a Backup.The response code from api is 200
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToBackupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200},
		MoreHeaders: RequestOpts.MoreHeaders}
	_, r.Err = c.Put(resourceURL(c, id), b, nil, reqOpt)
	return
}

type CreateOptsBuilder interface {
	ToBackupCreateMap() (map[string]interface{}, error)
}

type BackupDatabase struct {
	Name string `json:"name"`
}

type CreateOpts struct {
	InstanceID  string           `json:"instance_id" required:"true"`
	Name        string           `json:"name" required:"true"`
	Description string           `json:"description,omitempty"`
	Databases   []BackupDatabase `json:"databases,omitempty"`
}

func (opts CreateOpts) ToBackupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBackupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(baseURL(c), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

type ListOptsBuilder interface {
	ToBackupListQuery() (string, error)
}

type ListOpts struct {
	InstanceID string `q:"instance_id"`
	BackupID   string `q:"backup_id"`
	BackupType string `q:"backup_type"`
	BeginTime  string `q:"begin_time"`
	EndTime    string `q:"end_time"`
}

func (opts ListOpts) ToBackupListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(c)
	if opts != nil {
		q, err := opts.ToBackupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return BackupPage{SinglePageBase: pagination.SinglePageBase(r)}
	})
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

func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(backupURL(c, id), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}

type RestoreType string

const (
	TypeBackup    RestoreType = "backup"
	TypeTimestamp RestoreType = "timestamp"
)

type RestorePoint struct {
	InstanceID  string      `json:"instance_id" required:"true"`
	Type        RestoreType `json:"type" required:"true"`
	BackupID    string      `json:"backup_id,omitempty"`
	RestoreTime int         `json:"restore_time,omitempty"`
}

type RestoreToNewOpts struct {
	Name             string                    `json:"name" required:"true"`
	Ha               *instances.Ha             `json:"ha,omitempty"`
	ConfigurationId  string                    `json:"configuration_id,omitempty"`
	Port             string                    `json:"port,omitempty"`
	Password         string                    `json:"password" required:"true"`
	BackupStrategy   *instances.BackupStrategy `json:"backup_strategy,omitempty"`
	DiskEncryptionId string                    `json:"disk_encryption_id,omitempty"`
	FlavorRef        string                    `json:"flavor_ref" required:"true"`
	Volume           *instances.Volume         `json:"volume" required:"true"`
	AvailabilityZone string                    `json:"availability_zone" required:"true"`
	VpcId            string                    `json:"vpc_id" required:"true"`
	SubnetId         string                    `json:"subnet_id" required:"true"`
	SecurityGroupId  string                    `json:"security_group_id" required:"true"`
	RestorePoint     RestorePoint              `json:"restore_point" required:"true"`
}

type RestorePITROpts struct {
	Source Source `json:"source"`
	Target Target `json:"target"`
}

type Source struct {
	BackupID    string `json:"backup_id" required:"false"`
	InstanceID  string `json:"instance_id" required:"true"`
	RestoreTime int64  `json:"restore_time" required:"false"`
	Type        string `json:"type" required:"true"`
}

type Target struct {
	InstanceID string `json:"instance_id" required:"true"`
}

type RestoreToNewOptsBuilder interface {
	ToBackupRestoreMap() (map[string]interface{}, error)
}

type RestorePITROptsBuilder interface {
	ToPITRRestoreMap() (map[string]interface{}, error)
}

func (opts RestoreToNewOpts) ToBackupRestoreMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func (opts RestorePITROpts) ToPITRRestoreMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func RestoreToNew(c *golangsdk.ServiceClient, opts RestoreToNewOptsBuilder) (r RestoreResult) {
	b, err := opts.ToBackupRestoreMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(instances.CreateURL(c), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}

func RestorePITR(c *golangsdk.ServiceClient, opts RestorePITROptsBuilder) (r RestoreResult) {
	b, err := opts.ToPITRRestoreMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(restoreURL(c), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
