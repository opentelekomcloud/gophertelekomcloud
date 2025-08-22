package instance

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Instance ID.
	Id string `q:"id"`
	// Instance name.
	Name string `q:"name"`
	// Instance type.
	Type string `q:"type"`
	// Datastore type.
	DataStoreType string `q:"datastore_type"`
	// VPC ID.
	VpcId string `q:"vpc_id"`
	// Subnet ID.
	SubnetId string `q:"subnet_id"`
	// Offset from which the query starts.
	Offset int `q:"offset"`
	// Number of items displayed on each page.
	Limit int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]ListTaurusDBInstance, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return TaurusDBPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractInstances(pages)
}

func ExtractInstances(r pagination.NewPage) ([]ListTaurusDBInstance, error) {
	var s struct {
		Instances []ListTaurusDBInstance `json:"instances"`
	}
	err := extract.Into(bytes.NewReader((r.(TaurusDBPage)).Body), &s)
	return s.Instances, err
}

type TaurusDBPage struct {
	pagination.NewSinglePageBase
}

type ListTaurusDBInstance struct {
	Id                  string         `json:"id"`
	Name                string         `json:"name"`
	Status              string         `json:"status"`
	Type                string         `json:"type"`
	Port                string         `json:"port"`
	VpcId               string         `json:"vpc_id"`
	SubnetId            string         `json:"subnet_id"`
	SecurityGroupId     string         `json:"security_group_id"`
	Flavor              string         `json:"flavor_ref"`
	FlavorInfo          FlavorInfo     `json:"flavor_info"`
	ConfigurationId     string         `json:"configuration_id"`
	AZMode              string         `json:"az_mode"`
	MasterAZ            string         `json:"master_az_code"`
	TimeZone            string         `json:"time_zone"`
	ProjectId           string         `json:"project_id"`
	DbUserName          string         `json:"db_user_name"`
	PublicIps           []string       `json:"public_ips"`
	PrivateIps          []string       `json:"private_ips"`
	Created             string         `json:"created"`
	Updated             string         `json:"updated"`
	Volume              Volume         `json:"volume"`
	Nodes               []Nodes        `json:"nodes"`
	DataStore           DataStore      `json:"datastore"`
	BackupStrategy      BackupStrategy `json:"backup_strategy"`
	DedicatedResourceId string         `json:"dedicated_resource_id"`
	Tags                []Tags         `json:"tags"`
}

type FlavorInfo struct {
	Vcpus string `json:"vcpus"`
	Ram   string `json:"ram"`
}
