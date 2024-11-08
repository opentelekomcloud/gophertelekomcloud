package lifecycle

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOps is a struct that contains all the parameters.
type CreateDeprOpts struct {
	// Indicates the name of an instance.
	// An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports
	// only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name" required:"true"`

	// Indicates the description of an instance.
	// It is a character string containing not more than 1024 characters.
	Description string `json:"description,omitempty"`

	// Indicates a message engine.
	Engine string `json:"engine" required:"true"`

	// Indicates the version of a message engine.
	EngineVersion string `json:"engine_version" required:"true"`

	// BrokerNum is a number of brokers.
	BrokerNum *int `json:"broker_num,omitempty"`

	// Indicates the message storage space.
	StorageSpace int `json:"storage_space" required:"true"`

	// Indicates the baseline bandwidth of a Kafka instance, that is,
	// the maximum amount of data transferred per unit time. Unit: byte/s.
	Specification string `json:"specification,omitempty"`

	// Indicates the maximum number of topics in a Kafka instance.
	PartitionNum int `json:"partition_num,omitempty"`

	// Indicates a username.
	// A username consists of 1 to 64 characters
	// and supports only letters, digits, and hyphens (-).
	AccessUser string `json:"access_user,omitempty"`

	// Indicates the password of an instance.
	// An instance password must meet the following complexity requirements:
	// Must be 6 to 32 characters long.
	// Must contain at least two of the following character types:
	// Lowercase letters
	// Uppercase letters
	// Digits
	// Special characters (`~!@#$%^&*()-_=+\|[{}]:'",<.>/?)
	Password string `json:"password,omitempty"`

	// Indicates the ID of a VPC.
	VpcID string `json:"vpc_id" required:"true"`

	// Indicates the ID of a security group.
	SecurityGroupID string `json:"security_group_id" required:"true"`

	// Indicates the ID of a subnet.
	SubnetID string `json:"subnet_id" required:"true"`

	// Indicates the ID of an AZ.
	// The parameter value can be left blank or an empty array.
	AvailableZones []string `json:"available_zones" required:"true"`

	// Indicates a product ID.
	ProductID string `json:"product_id" required:"true"`

	// Indicates the time at which a maintenance time window starts.
	// Format: HH:mm:ss
	MaintainBegin string `json:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm:ss
	MaintainEnd string `json:"maintain_end,omitempty"`

	// Indicates whether to open the public network access function. Default to false.
	EnablePublicIP bool `json:"enable_publicip,omitempty"`

	// Indicates the bandwidth of the public network.
	PublicBandWidth int `json:"public_bandwidth,omitempty"`

	// Indicates the ID of the Elastic IP address bound to the instance.
	PublicIpID string `json:"publicip_id,omitempty"`

	// Indicates whether to enable SSL-encrypted access.
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// Indicates the action to be taken when the memory usage reaches the disk capacity threshold. Options:
	// time_base: Automatically delete the earliest messages.
	// produce_reject: Stop producing new messages.
	RetentionPolicy string `json:"retention_policy,omitempty"`

	// Indicates whether to enable automatic topic creation.
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	// Indicates the storage I/O specification. For details on how to select a disk type
	StorageSpecCode string `json:"storage_spec_code,omitempty"`

	// Indicates whether disk encryption is enabled.
	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	// Disk encryption key. If disk encryption is not enabled, this parameter is left blank.
	DiskEncryptedKey string `json:"disk_encrypted_key,omitempty"`

	// Indicates the tags of the instance
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

// Create an instance with given parameters.
func CreateDepr(client *golangsdk.ServiceClient, opts CreateDeprOpts) (*InstanceCreate, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res InstanceCreate
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// InstanceCreate response
type InstanceCreate struct {
	InstanceID string `json:"instance_id"`
}
