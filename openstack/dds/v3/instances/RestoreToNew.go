package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RestoreOpts struct {
	// Specifies the DB instance name. Instance name, which can be the same as an existing name.
	//
	// The instance name must contain 4 to 64 characters and must start with a letter. It is case sensitive and can contain letters, digits, hyphens (-), and underscores (_). It cannot contain other special characters.
	Name string `json:"name" required:"true"`
	// Specifies the AZ ID. You can select multiple AZs to create a cross-AZ cluster based on az_status returned by the API described in Querying Database Specifications.
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Specifies the VPC ID.
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the network ID of the subnet.
	SubnetId string `json:"subnet_id" required:"true"`
	// Specifies the security group ID.
	SecurityGroupId string `json:"security_group_id" required:"true"`
	// Specifies the database password.
	// The value must be 8 to 32 characters in length and contain uppercase letters (A to Z), lowercase letters (a to z), digits (0 to 9), and special characters, such as ~!@#%^*-_=+?
	// Enter a strong password to improve security, preventing security risks such as brute force cracking.
	Password string `json:"password,omitempty"`
	// Specifies the key ID used for disk encryption. The string must comply with UUID regular expression rules.
	// If this parameter is not transferred, disk encryption is not performed.
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	// Specifies the instance specifications.
	Flavor []Flavor `json:"flavor" required:"true"`
	// Specifies the advanced backup policy.
	BackupStrategy BackupStrategy `json:"backup_strategy,omitempty"`
	// Specifies whether to enable or disable SSL.
	// Valid value:
	// The value 0 indicates that SSL is disabled by default.
	// The value 1 indicates that SSL is enabled by default.
	// If this parameter is not transferred, SSL is enabled by default.
	Ssl string `json:"ssl_option,omitempty"`
	// Specifies the advanced backup policy.
	RestorePoint RestorePoint `json:"restore_point" required:"true"`
}

type RestorePoint struct {
	// Specifies the instance ID, which can be obtained by calling the API for querying instances.
	// If you do not have an instance, you can call the API used for creating an instance.
	// This parameter is optional when type is set to backup.
	// This parameter is mandatory when type is set to timestamp.
	InstanceId string `json:"instance_id,omitempty"`
	// Specifies the recovery mode. The enumerated values are as follows:
	// backup: indicates restoration from backup files. In this mode, backup_id is mandatory when type is optional.
	// timestamp: indicates point-in-time restoration. In this mode, restore_time is mandatory when type is mandatory.
	Type string `json:"type,omitempty"`
	// Specifies the ID of the backup to be restored.
	// This parameter must be specified when the backup file is used for restoration.
	BackupId string `json:"backup_id,omitempty"`
	// Specifies the time point of data restoration in the UNIX timestamp.
	// The unit is millisecond and the time zone is UTC.
	RestoreTime int `json:"restore_time,omitempty"`
}

func RestoreToNew(client *golangsdk.ServiceClient, opts RestoreOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}
