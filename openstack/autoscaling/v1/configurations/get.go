package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Configuration, error) {
	raw, err := client.Get(client.ServiceURL("scaling_configuration", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Configuration
	err = extract.IntoStructPtr(raw.Body, &res, "scaling_configuration")
	return &res, err
}

type Configuration struct {
	// Specifies the AS configuration ID. This parameter is globally unique.
	ID string `json:"scaling_configuration_id"`
	// Specifies the tenant ID.
	Tenant string `json:"tenant"`
	// Specifies the AS configuration name.
	// Supports fuzzy search.
	Name string `json:"scaling_configuration_name"`
	// Specifies the information about instance configurations.
	InstanceConfig InstanceConfig `json:"instance_config"`
	// Specifies the time when AS configurations are created. The time format complies with UTC.
	CreateTime string `json:"create_time"`
	// Specifies the ID of the AS group to which the AS configuration is bound.
	ScalingGroupId string `json:"scaling_group_id,omitempty"`
}

type InstanceConfig struct {
	// Specifies the ECS flavor ID.
	FlavorRef string `json:"flavorRef"`
	// Specifies the image ID. It is same as image_id.
	ImageRef string `json:"imageRef"`
	// Specifies the disk group information.
	Disk []Disk `json:"disk"`
	// Specifies the name of the SSH key pair used to log in to the ECS.
	SSHKey string `json:"key_name"`
	// Specifies the fingerprint of the SSH key pair used to log in to the ECS.
	KeyFingerprint string `json:"key_fingerprint"`
	// This parameter is reserved.
	InstanceName string `json:"instance_name"`
	// This parameter is reserved.
	InstanceID string `json:"instance_id"`
	// This parameter is reserved.
	AdminPass string `json:"adminPass"`
	// Specifies information about the injected file.
	Personality []Personality `json:"personality"`
	// Specifies the EIP of the ECS.
	PublicIp PublicIp `json:"public_ip"`
	// Specifies the Cloud-Init user data, which is encoded using Base64.
	UserData string `json:"user_data"`
	// Specifies the ECS metadata.
	Metadata AdminPassMetadata `json:"metadata"`
	// Specifies the security group information.
	SecurityGroups []SecurityGroup `json:"security_groups"`
	// This parameter is reserved.
	ServerGroupID string `json:"server_group_id"`
	// This parameter is reserved.
	Tenancy string `json:"tenancy"`
	// This parameter is reserved.
	DedicatedHostID string `json:"dedicated_host_id"`
	// This parameter is reserved.
	MarketType string `json:"market_type"`
	// This parameter is reserved.
	MultiFlavorPriorityPolicy string `json:"multi_flavor_priority_policy"`
}

type Disk struct {
	// Specifies the disk size. The unit is GB.
	// The system disk size ranges from 1 to 1024 and must be greater than or equal to the
	// minimum size (min_disk value) of the system disk specified in the image.
	// The data disk size ranges from 10 to 32768.
	Size int `json:"size"`
	// Specifies the ECS system disk type. The disk type must match the available disk type.
	// SATA: common I/O disk type
	// SAS: high I/O disk type
	// SSD: ultra-high I/O disk type
	// co-p1: high I/O (performance-optimized I) disk type
	// uh-l1: ultra-high I/O (latency-optimized) disk type
	// If the specified disk type is not available in the AZ, the disk will fail to create.
	// NOTE:
	// For HANA, HL1, and HL2 ECSs, use co-p1 and uh-l1 disks. For other ECSs, do not use co-p1 or uh-l1 disks.
	VolumeType string `json:"volume_type"`
	// Specifies the ECS system disk type. The disk type must match the available disk type.
	// SATA: common I/O disk type
	// SAS: high I/O disk type
	// SSD: ultra-high I/O disk type
	// co-p1: high I/O (performance-optimized I) disk type
	// uh-l1: ultra-high I/O (latency-optimized) disk type
	// If the specified disk type is not available in the AZ, the disk will fail to create.
	// NOTE:
	// For HANA, HL1, and HL2 ECSs, use co-p1 and uh-l1 disks. For other ECSs, do not use co-p1 or uh-l1 disks.
	DiskType string `json:"disk_type"`
	// Specifies a DSS device ID for creating an ECS disk.
	// NOTE:
	// Specify DSS devices for all disks in an AS configuration or not. If DSS devices are specified,
	// all the data stores must belong to the same AZ, and the disk types supported by a DSS device for a disk
	// must be the same as the volume_type value.
	DedicatedStorageID string `json:"dedicated_storage_id"`
	// Specifies the ID of a data disk image used to export data disks of an ECS.
	DataDiskImageID string `json:"data_disk_image_id"`
	// Specifies the disk backup snapshot ID for restoring the system disk
	// and data disks using a full-ECS backup when a full-ECS image is used.
	// NOTE:
	// Each disk in an AS configuration must correspond to a disk backup in the full-ECS backup by snapshot_id.
	SnapshotID string `json:"snapshot_id"`
	// Specifies the metadata for creating disks.
	Metadata map[string]any `json:"metadata"`
}

type Personality struct {
	// Specifies the path of the injected file.
	// For Linux OSs, specify the path, for example, /etc/foo.txt, for storing the injected file.
	// For Windows, the injected file is automatically stored in the root directory of drive C.
	// You only need to specify the file name, for example, foo. The file name contains only letters and digits.
	Path string `json:"path"`
	// Specifies the content of the injected file.
	// The value must be the information after the content of the injected file is encoded using Base64.
	Content string `json:"content"`
}

type PublicIp struct {
	// Specifies the EIP automatically assigned to the ECS.
	Eip Eip `json:"eip,omitempty"`
}

type Eip struct {
	// Specifies the EIP type.
	// Enumerated value of the IP address type: 5_bgp (indicates dynamic BGP)
	Type string `json:"ip_type"`
	// Specifies the bandwidth of an IP address.
	Bandwidth Bandwidth `json:"bandwidth"`
}

type Bandwidth struct {
	// Specifies the bandwidth (Mbit/s). The value range is 1 to 500.
	// NOTE:
	// The specific range may vary depending on the configuration in each region.
	// You can see the bandwidth range of each region on the management console.
	// The minimum unit for bandwidth varies depending on the bandwidth range.
	// The minimum unit is 1 Mbit/s if the allowed bandwidth size ranges from 0 to 300 Mbit/s (with 300 Mbit/s included).
	// The minimum unit is 50 Mbit/s if the allowed bandwidth size ranges 300 Mbit/s to 500 Mbit/s (with 500 Mbit/s included).
	Size int `json:"size"`
	// Specifies the bandwidth sharing type.
	// Enumerated values of the sharing type:
	// PER: dedicated
	// Only dedicated bandwidth is available.
	ShareType string `json:"share_type"`
	// Specifies the bandwidth billing mode.
	// traffic: billed by traffic.
	// If the parameter value is out of the preceding options, creating the ECS will fail.
	ChargingMode string `json:"charging_mode"`
}

type SecurityGroup struct {
	// Specifies the ID of the security group.
	ID string `json:"id"`
}
