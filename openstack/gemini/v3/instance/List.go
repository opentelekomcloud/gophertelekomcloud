package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListGeminiDBOpts struct {
	Id            string `q:"id"`
	Name          string `q:"name"`
	DataStoreType string `q:"datastore_type"`
	Mode          string `q:"mode"`
	VpcId         string `q:"vpc_id"`
	SubnetId      string `q:"subnet_id"`
	Offset        int    `q:"offset"`
	Limit         int    `q:"limit"`
}

func ListGeminiDB(client *golangsdk.ServiceClient, opts ListGeminiDBOpts) (*ListGeminiDBResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListGeminiDBResponse
	return &res, extract.Into(raw.Body, &res)
}

type ListGeminiDBResponse struct {
	Instances  []ListResult `json:"instances"`
	TotalCount int          `json:"total_count"`
}

type ListResult struct {
	Id                  string                   `json:"id"`
	Name                string                   `json:"name"`
	Status              string                   `json:"status"`
	Port                string                   `json:"port"`
	Mode                string                   `json:"mode"`
	Region              string                   `json:"region"`
	DataStore           ListDatastoreResult      `json:"datastore"`
	Engine              string                   `json:"engine"`
	Created             string                   `json:"created"`
	Updated             string                   `json:"updated"`
	DbUserName          string                   `json:"db_user_name"`
	VpcId               string                   `json:"vpc_id"`
	SubnetId            string                   `json:"subnet_id"`
	SecurityGroupId     string                   `json:"security_group_id"`
	BackupStrategy      ListBackupStrategyResult `json:"backup_strategy"`
	PayMode             string                   `json:"pay_mode"`
	MaintenanceWindow   string                   `json:"maintenance_window"`
	Groups              []ListGroupResult        `json:"groups"`
	EnterpriseProjectId string                   `json:"enterprise_project_id"`
	TimeZone            string                   `json:"time_zone"`
	Actions             []string                 `json:"actions"`
}

type ListDatastoreResult struct {
	Type         string `json:"type"`
	Version      string `json:"version"`
	WholeVersion string `json:"whole_version"`
}

type ListBackupStrategyResult struct {
	StartTime string `json:"start_time"`
	KeepDays  int    `json:"keep_days"`
}

type ListGroupResult struct {
	Id     string           `json:"id"`
	Status string           `json:"status"`
	Volume Volume           `json:"volume"`
	Nodes  []ListNodeResult `json:"nodes"`
}

type Volume struct {
	Size string `json:"size"`
	Used string `json:"used"`
}

type ListNodeResult struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	SubnetId         string `json:"subnet_id"`
	PrivateIp        string `json:"private_ip"`
	PublicIp         string `json:"public_ip"`
	SpecCode         string `json:"spec_code"`
	AvailabilityZone string `json:"availability_zone"`
	SupportReduce    bool   `json:"support_reduce"`
}
