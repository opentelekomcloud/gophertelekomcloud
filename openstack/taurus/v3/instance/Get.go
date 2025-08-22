package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, instanceId string) (*TaurusDBInstance, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TaurusDBInstance
	err = extract.IntoStructPtr(raw.Body, &res, "instance")
	return &res, err
}

type TaurusDBInstance struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Alias             string   `json:"alias"`
	Status            string   `json:"status"`
	Type              string   `json:"type"`
	Port              string   `json:"port"`
	NodeCount         int      `json:"node_count"`
	VpcId             string   `json:"vpc_id"`
	SubnetId          string   `json:"subnet_id"`
	SecurityGroupId   string   `json:"security_group_id"`
	ConfigurationId   string   `json:"configuration_id"`
	AZMode            string   `json:"az_mode"`
	MasterAZ          string   `json:"master_az_code"`
	TimeZone          string   `json:"time_zone"`
	ProjectId         string   `json:"project_id"`
	DbUserName        string   `json:"db_user_name"`
	PublicIps         string   `json:"public_ips"`
	PrivateIps        []string `json:"private_write_ips"`
	Created           string   `json:"created"`
	Updated           string   `json:"updated"`
	MaintenanceWindow string   `json:"maintenance_window"`
	BackupUsedSpace   float64  `json:"backup_used_space"`

	Nodes          []Nodes        `json:"nodes"`
	DataStore      DataStore      `json:"datastore"`
	BackupStrategy BackupStrategy `json:"backup_strategy"`
	ChargeInfo     ChargeInfo     `json:"charge_info"`
	Tags           []Tags         `json:"tags"`
	Proxies        []ProxyInfo    `json:"proxies"`

	DedicatedResourceId string `json:"dedicated_resource_id"`
}

type Volume struct {
	Type string `json:"type"`
	Used string `json:"used"`
}

type Nodes struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Status           string     `json:"status"`
	PrivateIps       []string   `json:"private_read_ips"`
	Port             int        `json:"port"`
	Flavor           string     `json:"flavor_ref"`
	FlavorId         string     `json:"flavor_id"`
	Region           string     `json:"region_code"`
	AvailabilityZone string     `json:"az_code"`
	Volume           NodeVolume `json:"volume"`
	Created          string     `json:"created"`
	Updated          string     `json:"updated"`
	MaxConnections   string     `json:"max_connections"`
	Vcpus            string     `json:"vcpus"`
	Ram              string     `json:"ram"`
	NeedRestart      bool       `json:"need_restart"`
	Priority         int        `json:"priority"`
}

type NodeVolume struct {
	Type string `json:"type"`
	Size int64  `json:"size"`
	Used string `json:"used"`
}

type ProxyInfo struct {
	PoolId  string `json:"pool_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
