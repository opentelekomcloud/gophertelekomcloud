package instances

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts is a struct that contains all the parameters for a creation.
type CreateOpts struct {
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
	BrokerNum int `json:"broker_num" required:"true"`

	// Indicates the message storage space.
	StorageSpace int `json:"storage_space" required:"true"`

	// // Indicates the baseline bandwidth of a Kafka instance, that is,
	// // the maximum amount of data transferred per unit time. Unit: byte/s.
	// Specification string `json:"specification,omitempty"`
	//
	// // Indicates the maximum number of topics in a Kafka instance.
	// PartitionNum int `json:"partition_num,omitempty"`

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
	// Format: HH:mm
	MaintainBegin string `json:"maintain_begin,omitempty"`

	// Indicates the time at which a maintenance time window ends.
	// Format: HH:mm
	MaintainEnd string `json:"maintain_end,omitempty"`

	// Indicates whether to open the public network access function. Default to false.
	EnablePublicIP bool `json:"enable_publicip,omitempty"`

	// // Indicates the bandwidth of the public network.
	// PublicBandWidth int `json:"public_bandwidth,omitempty"`

	// Indicates the ID of the Elastic IP address bound to the instance.
	PublicIpID string `json:"publicip_id,omitempty"`

	// Indicates whether to enable SSL-encrypted access.
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// Security protocol to use after SASL is enabled. This parameter is mandatory if SASL authentication is enabled (ssl_enable=true).
	// If this parameter is left blank, SASL_SSL authentication is enabled by default.
	// This setting is fixed once the instance is created.
	//    SASL_SSL: Data is encrypted with SSL certificates for high-security transmission.
	//    SASL_PLAINTEXT: Data is transmitted in plaintext with username and password authentication. This protocol uses the SCRAM-SHA-512 mechanism and delivers high performance.
	KafkaSecurityProtocol string `json:"kafka_security_protocol,omitempty"`

	// Authentication mechanism to use after SASL is enabled. This parameter is mandatory if SASL authentication is enabled (ssl_enable=true).
	// If this parameter is left blank, PLAIN authentication is enabled by default.
	// Select both or either of the following mechanisms for SASL authentication. Options:
	//    PLAIN: simple username and password verification.
	//    SCRAM-SHA-512: user credential verification, which is more secure than PLAIN.
	SaslEnabledMechanisms []string `json:"sasl_enabled_mechanisms,omitempty"`

	// Indicates the action to be taken when the memory usage reaches the disk capacity threshold. Options:
	// time_base: Automatically delete the earliest messages.
	// produce_reject: Stop producing new messages.
	RetentionPolicy string `json:"retention_policy,omitempty"`

	// Indicates whether to enable IPv6. This parameter is available only when the VPC supports IPv6.
	// Default: false
	IPv6Enable bool `json:"ipv6_enable,omitempty"`

	// Indicates whether disk encryption is enabled.
	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	// Disk encryption key. If disk encryption is not enabled, this parameter is left blank.
	DiskEncryptedKey string `json:"disk_encrypted_key,omitempty"`

	// Indicates whether to enable automatic topic creation.
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	// Indicates the storage I/O specification. For details on how to select a disk type
	StorageSpecCode string `json:"storage_spec_code,omitempty"`

	// Enterprise project ID. This parameter is mandatory for an enterprise project account.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`

	// Indicates the tags of the instance
	Tags []tags.ResourceTag `json:"tags,omitempty"`

	// CPU architecture. Currently, only the x86 architecture is supported.
	// Options:
	//    X86
	ArchType string `json:"arch_type" required:"true"`

	// Intra-VPC plaintext access.
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

	// Parameter related to the yearly/monthly billing mode.
	// If this parameter is left blank, the billing mode is pay-per-use by default. If this parameter is not left blank, the billing mode is yearly/monthly.
	BSSParam *BSSParam `json:"bss_param,omitempty"`
}

type BSSParam struct {
	// Whether auto renewal is enabled.
	// Options:
	//    true: Auto renewal is enabled.
	//    false: Auto renewal is not enabled.
	// By default, auto renewal is disabled.
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// Billing mode.
	// This parameter specifies a payment mode.
	// Options:
	//    prePaid: yearly/monthly billing.
	//    postPaid: pay-per-use billing.
	// The default value is postPaid.
	ChargingMode string `json:"charging_mode,omitempty"`

	// Specifies whether the order is automatically or manually paid.
	// Options:
	//    true: The order will be automatically paid.
	//    false: The order must be manually paid.
	// The default payment mode is manual.
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// Subscription period type.
	// Options:
	//    month
	//    year:
	// This parameter is valid and mandatory only when chargingMode is set to prePaid. **
	PeriodType string `json:"period_type,omitempty"`

	// Subscribed periods.
	//
	// Options:
	//    If periodType is month, the value ranges from 1 to 9.
	//    If periodType is year, the value ranges from 1 to 3.
	// **This parameter is valid and mandatory only when chargingMode is set to prePaid. **
	PeriodNum int `json:"period_num,omitempty"`
}

// InstanceIDResp response
type InstanceIDResp struct {
	InstanceID string `json:"instance_id"`
}

// Create an instance with given parameters.
// Send POST to /v2/{engine}/{project_id}/instances
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*InstanceIDResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// Here we should patch a client, because for a creation url path is different.
	// For all requests we use schema /v2/{project_id}/instances
	// But for a creation             /v2/{engine}/{project_id}/instances
	paths := strings.SplitN(client.Endpoint, "v2", 2)
	url := fmt.Sprintf("%sv2/%s%s%s", paths[0], opts.Engine, paths[1], resourcePath)

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res InstanceIDResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
