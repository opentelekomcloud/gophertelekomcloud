package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                string             `json:"name" required:"true"`
	Region              string             `json:"region" required:"true"`
	Mode                string             `json:"mode" required:"true"`
	Flavor              string             `json:"flavor_ref" required:"true"`
	VpcId               string             `json:"vpc_id" required:"true"`
	SubnetId            string             `json:"subnet_id" required:"true"`
	SecurityGroupId     string             `json:"security_group_id,omitempty"`
	Password            string             `json:"password" required:"true"`
	TimeZone            string             `json:"time_zone,omitempty"`
	AZMode              string             `json:"availability_zone_mode" required:"true"`
	SlaveCount          int                `json:"slave_count" required:"true"`
	MasterAZ            string             `json:"master_availability_zone,omitempty"`
	ConfigurationId     string             `json:"configuration_id,omitempty"`
	DedicatedResourceId string             `json:"dedicated_resource_id,omitempty"`
	LowerCaseTableNames *int               `json:"lower_case_table_names,omitempty"`
	DataStore           DataStoreOpt       `json:"datastore" required:"true"`
	BackupStrategy      *BackupStrategyOpt `json:"backup_strategy,omitempty"`
	ChargeInfo          *ChargeInfo        `json:"charge_info,omitempty"`
	Volume              *VolumeOpt         `json:"volume,omitempty"`
	Tags                *Tags              `json:"tags,omitempty"`
}

type DataStoreOpt struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  string `json:"keep_days" required:"true"`
}

type ChargeInfo struct {
	ChargingMode string `json:"charge_mode,omitempty"`
	OrderId      string `json:"order_id,omitempty"`
}

type VolumeOpt struct {
	Size int `json:"size" required:"true"`
}

type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201, 202},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	return &res, extract.Into(raw.Body, &res)
}

type CreateResponse struct {
	Instance TaurusDBResponse `json:"instance"`
	JobId    string           `json:"job_id"`
}

type TaurusDBResponse struct {
	Id              string         `json:"id"`
	Name            string         `json:"name"`
	Status          string         `json:"status"`
	Region          string         `json:"region"`
	Mode            string         `json:"mode"`
	Port            string         `json:"port"`
	VpcId           string         `json:"vpc_id"`
	SubnetId        string         `json:"subnet_id"`
	Flavor          string         `json:"flavor_ref"`
	SecurityGroupId string         `json:"security_group_id"`
	ConfigurationId string         `json:"configuration_id"`
	AZMode          string         `json:"availability_zone_mode"`
	MasterAZ        string         `json:"master_availability_zone"`
	SlaveCount      int            `json:"slave_count"`
	DataStore       DataStore      `json:"datastore"`
	BackupStrategy  BackupStrategy `json:"backup_strategy"`
	ChargeInfo      ChargeInfo     `json:"charge_info"`
}

type DataStore struct {
	Type          string `json:"type"`
	Version       string `json:"version"`
	KernelVersion string `json:"kernel_version"`
}

type BackupStrategy struct {
	StartTime string `json:"start_time"`
	KeepDays  string `json:"keep_days"`
}
