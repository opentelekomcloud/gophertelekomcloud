package configurations

import "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("scaling_configuration", id), &r.Body, nil)
	return
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (Configuration, error) {
	var s Configuration
	err := r.ExtractIntoStructPtr(&s, "scaling_configuration")
	return s, err
}

type Configuration struct {
	ID             string         `json:"scaling_configuration_id"`
	Tenant         string         `json:"tenant"`
	Name           string         `json:"scaling_configuration_name"`
	InstanceConfig InstanceConfig `json:"instance_config"`
	CreateTime     string         `json:"create_time"`
}

type InstanceConfig struct {
	FlavorRef                 string                 `json:"flavorRef"`
	ImageRef                  string                 `json:"imageRef"`
	Disk                      []Disk                 `json:"disk"`
	SSHKey                    string                 `json:"key_name"`
	KeyFingerprint            string                 `json:"key_fingerprint"`
	InstanceName              string                 `json:"instance_name"`
	InstanceID                string                 `json:"instance_id"`
	AdminPass                 string                 `json:"adminPass"`
	Personality               []Personality          `json:"personality"`
	PublicIp                  PublicIp               `json:"public_ip"`
	UserData                  string                 `json:"user_data"`
	Metadata                  map[string]interface{} `json:"metadata"`
	SecurityGroups            []SecurityGroup        `json:"security_groups"`
	ServerGroupID             string                 `json:"server_group_id"`
	Tenancy                   string                 `json:"tenancy"`
	DedicatedHostID           string                 `json:"dedicated_host_id"`
	MarketType                string                 `json:"market_type"`
	MultiFlavorPriorityPolicy string                 `json:"multi_flavor_priority_policy"`
}

type Disk struct {
	Size               int                    `json:"size"`
	VolumeType         string                 `json:"volume_type"`
	DiskType           string                 `json:"disk_type"`
	DedicatedStorageID string                 `json:"dedicated_storage_id"`
	DataDiskImageID    string                 `json:"data_disk_image_id"`
	SnapshotID         string                 `json:"snapshot_id"`
	Metadata           map[string]interface{} `json:"metadata"`
}

type Personality struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type PublicIp struct {
	Eip Eip `json:"eip"`
}

type Eip struct {
	Type      string    `json:"ip_type"`
	Bandwidth Bandwidth `json:"bandwidth"`
}

type Bandwidth struct {
	Size         int    `json:"size"`
	ShareType    string `json:"share_type"`
	ChargingMode string `json:"charging_mode"`
}

type SecurityGroup struct {
	ID string `json:"id"`
}
