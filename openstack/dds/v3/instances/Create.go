package instances

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name             string         `json:"name" required:"true"`
	DataStore        DataStore      `json:"datastore" required:"true"`
	Region           string         `json:"region" required:"true"`
	AvailabilityZone string         `json:"availability_zone" required:"true"`
	VpcId            string         `json:"vpc_id" required:"true"`
	SubnetId         string         `json:"subnet_id" required:"true"`
	SecurityGroupId  string         `json:"security_group_id" required:"true"`
	Port             string         `json:"port,omitempty"`
	Password         string         `json:"password" required:"true"`
	DiskEncryptionId string         `json:"disk_encryption_id,omitempty"`
	Mode             string         `json:"mode" required:"true"`
	Flavor           []Flavor       `json:"flavor" required:"true"`
	BackupStrategy   BackupStrategy `json:"backup_strategy" required:"true"`
	Ssl              string         `json:"ssl_option,omitempty"`
}

type DataStore struct {
	Type          string `json:"type" required:"true"`
	Version       string `json:"version" required:"true"`
	StorageEngine string `json:"storage_engine" required:"true"`
}

type Flavor struct {
	Type     string `json:"type" required:"true"`
	Num      int    `json:"num" required:"true"`
	Storage  string `json:"storage,omitempty"`
	Size     int    `json:"size,omitempty"`
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategy struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  int    `json:"keep_days,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*Instance, error) {
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Instance struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	DataStore        DataStore         `json:"datastore"`
	CreatedAt        string            `json:"created"`
	Status           string            `json:"status"`
	Region           string            `json:"region"`
	AvailabilityZone string            `json:"availability_zone"`
	VpcId            string            `json:"vpc_id"`
	SubnetId         string            `json:"subnet_id"`
	SecurityGroupId  string            `json:"security_group_id"`
	DiskEncryptionId string            `json:"disk_encryption_id"`
	Mode             string            `json:"mode"`
	Flavor           []FlavorOpt       `json:"flavor"`
	BackupStrategy   BackupStrategyOpt `json:"backup_strategy"`
	Ssl              string            `json:"ssl_option"`
	JobId            string            `json:"job_id"`
}

type FlavorOpt struct {
	Type     string `json:"type" required:"true"`
	Num      string `json:"num" required:"true"`
	Storage  string `json:"storage,omitempty"`
	Size     string `json:"size,omitempty"`
	SpecCode string `json:"spec_code" required:"true"`
}

type BackupStrategyOpt struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  string `json:"keep_days,omitempty"`
}
