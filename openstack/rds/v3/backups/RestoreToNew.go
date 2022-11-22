package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
)

type RestoreToNewOpts struct {
	//
	Name string `json:"name" required:"true"`
	//
	Ha *instances.Ha `json:"ha,omitempty"`
	//
	ConfigurationId string `json:"configuration_id,omitempty"`
	//
	Port string `json:"port,omitempty"`
	//
	Password string `json:"password" required:"true"`
	//
	BackupStrategy *instances.BackupStrategy `json:"backup_strategy,omitempty"`
	//
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	//
	FlavorRef string `json:"flavor_ref" required:"true"`
	//
	Volume *instances.Volume `json:"volume" required:"true"`
	//
	AvailabilityZone string `json:"availability_zone" required:"true"`
	//
	VpcId string `json:"vpc_id" required:"true"`
	//
	SubnetId string `json:"subnet_id" required:"true"`
	//
	SecurityGroupId string `json:"security_group_id" required:"true"`
	//
	RestorePoint RestorePoint `json:"restore_point" required:"true"`
}

type RestoreType string

const (
	TypeBackup    RestoreType = "backup"
	TypeTimestamp RestoreType = "timestamp"
)

type RestorePoint struct {
	//
	InstanceID string `json:"instance_id" required:"true"`
	//
	Type RestoreType `json:"type" required:"true"`
	//
	BackupID string `json:"backup_id,omitempty"`
	//
	RestoreTime int `json:"restore_time,omitempty"`
}

type RestoreToNewOptsBuilder interface {
	ToBackupRestoreMap() (map[string]interface{}, error)
}

func (opts RestoreToNewOpts) ToBackupRestoreMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func RestoreToNew(c *golangsdk.ServiceClient, opts RestoreToNewOptsBuilder) (r RestoreResult) {
	b, err := opts.ToBackupRestoreMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(c.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	return
}
