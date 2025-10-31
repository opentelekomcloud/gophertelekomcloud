package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                string             `json:"name" required:"true"`
	DataStore           DataStoreOpt       `json:"datastore" required:"true"`
	Region              string             `json:"region" required:"true"`
	AvailabilityZone    string             `json:"availability_zone" required:"true"`
	VpcId               string             `json:"vpc_id" required:"true"`
	SubnetId            string             `json:"subnet_id" required:"true"`
	SecurityGroupId     string             `json:"security_group_id" required:"true"`
	Password            string             `json:"password" required:"true"`
	Mode                string             `json:"mode" required:"true"`
	Flavor              []FlavorOpt        `json:"flavor" required:"true"`
	ConfigurationId     string             `json:"configuration_id,omitempty"`
	BackupStrategy      *BackupStrategyOpt `json:"backup_strategy,omitempty"`
	EnterpriseProjectId *string            `json:"enterprise_project_id,omitempty"`
	SslOption           *string            `json:"ssl_option,omitempty"`
}

type DataStoreOpt struct {
	Type          string `json:"type" required:"true"`
	Version       string `json:"version" required:"true"`
	StorageEngine string `json:"storage_engine" required:"true"`
}

type FlavorOpt struct {
	Num      string `json:"num" required:"true"`
	Size     string `json:"size" required:"true"`
	Storage  string `json:"storage" required:"true"`
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  string `json:"keep_days,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CreateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResp
	return &res, extract.Into(raw.Body, &res)
}

type CreateResp struct {
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	DataStore        DataStore      `json:"datastore"`
	Created          string         `json:"created"`
	Status           string         `json:"status"`
	Region           string         `json:"region"`
	AvailabilityZone string         `json:"availability_zone"`
	VpcId            string         `json:"vpc_id"`
	SubnetId         string         `json:"subnet_id"`
	SecurityGroupId  string         `json:"security_group_id"`
	Mode             string         `json:"mode"`
	Flavor           []Flavor       `json:"flavor"`
	BackupStrategy   BackupStrategy `json:"backup_strategy"`
	SslOption        string         `json:"ssl_option"`
	JobId            string         `json:"job_id"`
}

type DataStore struct {
	Type          string `json:"type"`
	Version       string `json:"version"`
	StorageEngine string `json:"storage_engine"`
}

type Flavor struct {
	Num      string `json:"num"`
	Size     string `json:"size"`
	Storage  string `json:"storage"`
	SpecCode string `json:"spec_code"`
}

type BackupStrategy struct {
	StartTime string `json:"start_time"`
	KeepDays  string `json:"keep_days"`
}
