package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListInstanceOpts struct {
	Id            string `q:"id"`
	Name          string `q:"name"`
	Mode          string `q:"mode"`
	DataStoreType string `q:"datastore_type"`
	VpcId         string `q:"vpc_id"`
	SubnetId      string `q:"subnet_id"`
	Offset        int    `q:"offset"`
	Limit         int    `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListInstanceOpts) (*ListResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Instances  []InstanceResponse `json:"instances"`
	TotalCount int                `json:"total_count"`
}

type InstanceResponse struct {
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	Status            string         `json:"status"`
	Port              int            `json:"port,string"`
	Mode              string         `json:"mode"`
	Region            string         `json:"region"`
	DataStore         DataStore      `json:"datastore"`
	Engine            string         `json:"engine"`
	Created           string         `json:"created"`
	Updated           string         `json:"updated"`
	DbUserName        string         `json:"db_user_name"`
	Ssl               int            `json:"ssl"`
	VpcId             string         `json:"vpc_id"`
	SubnetId          string         `json:"subnet_id"`
	SecurityGroupId   string         `json:"security_group_id"`
	BackupStrategy    BackupStrategy `json:"backup_strategy"`
	MaintenanceWindow string         `json:"maintenance_window"`
	Groups            []Group        `json:"groups"`
	DiskEncryptionId  string         `json:"disk_encryption_id"`
	TimeZone          string         `json:"time_zone"`
	Actions           []string       `json:"actions"`
	PayMode           string         `json:"pay_mode"`
}

type Group struct {
	Type   string  `json:"type"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Volume Volume  `json:"volume"`
	Nodes  []Nodes `json:"nodes"`
}

type Volume struct {
	Size string `json:"size"`
	Used string `json:"used"`
}

type Nodes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	Role             string `json:"role"`
	PrivateIP        string `json:"private_ip"`
	PublicIP         string `json:"public_ip"`
	SpecCode         string `json:"spec_code"`
	AvailabilityZone string `json:"availability_zone"`
}
