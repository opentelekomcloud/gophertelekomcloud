package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOps struct {
	// DCS instance name.
	// An instance name is a string of 4–64 characters
	// that contain letters, digits, underscores (_), and hyphens (-).
	// An instance name must start with letters.
	Name string `json:"name" required:"true"`
	// Brief description of the DCS instance.
	// The description supports up to 1024 characters.
	// The backslash (\) and quotation mark (") are special characters for JSON messages.
	// When using these characters in a parameter value,
	// add the escape character (\) before the characters, for example, \\ and \".
	Description string `json:"description,omitempty"`
	// Cache engine, which is Redis.
	Engine string `json:"engine" required:"true"`
	// Cache engine version. If the cache engine is Redis, the value can be 3.0, 4.0, or 5.0.
	EngineVersion string `json:"engine_version" required:"true"`
	// DCS instance specification code.
	// To obtain the value, go to the instance creation page on the DCS console,
	// and view Flavor Name in the Instance Specification table.
	SpecCode string `json:"spec_code" required:"true"`
	// Cache capacity. Unit: GB.
	// For a single-node or master/standby DCS Redis 3.0 instance,
	// the value can be 2, 4, 8, 16, 32, or 64. For a Proxy Cluster DCS Redis 3.0 instance,
	// the value can be 64, 128, 256, 512, or 1024.
	// For a single-node or master/standby DCS Redis 4.0 or 5.0 instance,
	// the value can be 0.125, 0.25, 0.5, 1, 2, 4, 8, 16, 32, 24, 48, or 64.
	// For a Redis Cluster DCS Redis 4.0 or 5.0 instance,
	// the value can be 4, 8, 16, 24, 32, 48, 64, 96, 128, 192, 256, 384, 512, 768, or 1024.
	Capacity float64 `json:"capacity" required:"true"`
	// Password of a DCS instance.
	// The password of a DCS Redis instance must meet the following complexity requirements:
	// Must be a string consisting of 8 to 32 characters.
	// Must be different from the old password.
	// Contains at least three of the following character types:
	// Lowercase letters
	// Uppercase letters
	// Digits
	// Special characters (`~!@#$^&*()-_=+\|{}:,<.>/?)
	Password string `json:"password,omitempty"`
	// VPC ID.
	// Obtain the value by using either of the following methods:
	// Method 1: Log in to VPC console and view the VPC ID in the VPC details.
	// Method 2: Call the API for querying VPCs.
	// For details, see the "Querying VPCs" section in the Virtual Private Cloud API Reference.
	VPCId string `json:"vpc_id" required:"true"`
	// ID of the security group which the instance belongs to.
	// This parameter is mandatory when the engine is Redis and engine_version is 3.0.
	// DCS Redis 3.0 instances support security group access control.
	// This parameter is optional when the engine is Redis and engine_version is 4.0 or 5.0.
	// DCS Redis 4.0 and 5.0 instances do not support security groups.
	// Obtain the value by using either of the following methods:
	// Method 1: Log in to the VPC console and view the security group ID on the security group details page.
	// Method 2: Call the API for querying security groups. For details,
	// see the "Querying Security Groups" section in the Virtual Private Cloud API Reference.
	SecurityGroupID string `json:"security_group_id,omitempty"`
	// Network ID of the subnet.
	// Obtain the value by using either of the following methods:
	// Method 1: Log in to VPC console and click the target subnet on the Subnets tab page.
	// You can view the network ID on the displayed page.
	// Method 2: Call the API for querying subnets.
	// For details, see the "Querying Subnets" section in the Virtual Private Cloud API Reference.
	SubnetID string `json:"subnet_id" required:"true"`
	// ID of the AZ where the cache node resides and which has available resources.
	// For details on how to obtain the value, see Querying AZ Information.
	// Check whether the AZ has available resources.
	// Master/Standby, Proxy Cluster, and Redis Cluster DCS instances support cross-AZ deployment.
	// You can specify an AZ for the standby node. When specifying AZs for nodes,
	// use commas (,) to separate multiple AZs. For details, see the example request.
	AvailableZones []string `json:"available_zones" required:"true"`
	// Backup policy.
	// This parameter is available for master/standby and cluster DCS instances.
	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty"`
	// An indicator of whether to enable public access for a DCS Redis instance.
	EnablePublicIp *bool `json:"enable_publicip,omitempty"`
	// ID of the elastic IP address bound to a DCS Redis instance.
	// This parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).
	PublicIpId string `json:"publicip_id,omitempty"`
	// IP address that is manually specified for a DCS instance.
	PrivateIps []string `json:"private_ips,omitempty"`
	// An indicator of whether to enable SSL for public access to a DCS Redis instance.
	EnableSsl *bool `json:"enable_ssl,omitempty"`
	// Time at which the maintenance time window starts.
	// Format: hh:mm:ss.
	// The start time and end time of the maintenance time window
	// must indicate the time segment of a supported maintenance time window.
	// For details on how to query the time segments of supported maintenance time windows,
	// see Querying Maintenance Time Window.
	// The start time must be set to 22:00:00, 02:00:00, 06:00:00, 10:00:00, 14:00:00, or 18:00: 00.
	// Parameters maintain_begin and maintain_end must be set in pairs.
	// If parameter maintain_start is left blank, parameter maintain_end is also blank.
	// In this case, the system automatically set the start time to 02:00:00.
	MaintainBegin string `json:"maintain_begin,omitempty"`
	// The end time is four hours later than the start time.
	// For example, if the start time is 22:00:00, the end time is 02:00:00.
	// ...In this case, the system automatically set the end time to 06:00:00.
	MaintainEnd string `json:"maintain_end,omitempty"`
	// Port customization, which is supported only by Redis 4.0 and Redis 5.0 instances and not by Redis 3.0 instances.
	// If this parameter is not sent or is left empty when you create a Redis 4.0 or 5.0 instance,
	// the default port 6379 will be used. To customize a port, specify a port number in the range from 1 to 65535.
	Port int32 `json:"port,omitempty"`
	// Critical command renaming, which is supported only by Redis 4.0 and Redis 5.0 instances and not by Redis 3.0 instances.
	// If this parameter is not sent or is left empty when you create a Redis 4.0 or 5.0 instance, no critical command will be renamed.
	// Currently, only COMMAND, KEYS, FLUSHDB, FLUSHALL, and HGETALL commands can be renamed.
	RenameCommands *interface{} `json:"rename_commands,omitempty"`
	// An indicator of whether a DCS instance can be accessed in password-free mode.
	// true: indicates that a DCS instance can be accessed without a password.
	// false: indicates that a DCS instance can be accessed only after password authentication.
	NoPasswordAccess string `json:"no_password_access"`
}

type InstanceBackupPolicy struct {
	// Retention time.
	// Unit: day.
	// Range: 1–7.
	SaveDays int `json:"save_days"`
	// Backup type. Options:
	// auto: automatic backup.
	// manual: manual backup.
	BackupType string `json:"backup_type"`
	// Backup plan.
	PeriodicalBackupPlan PeriodicalBackupPlan `json:"periodical_backup_plan" required:"true"`
}

type PeriodicalBackupPlan struct {
	// Time at which backup starts.
	// "00:00-01:00" indicates that backup starts at 00:00:00.
	BeginAt string `json:"begin_at" required:"true"`
	// Interval at which backup is performed.
	// Currently, only weekly backup is supported.
	PeriodType string `json:"period_type" required:"true"`
	// Day in a week on which backup starts.
	// Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday.
	BackupAt []int `json:"backup_at" required:"true"`
	// Time zone in which backup is performed.
	// Value range: GMT–12:00 to GMT+12:00. If this parameter is left blank,
	// the current time zone of the DCS-Server VM is used by default.
	TimezoneOffset string `json:"timezone_offset,omitempty"`
}

// Create an instance with given parameters.
func Create(client *golangsdk.ServiceClient, opts CreateOps) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		InstanceID string `json:"instance_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.InstanceID, err
}
