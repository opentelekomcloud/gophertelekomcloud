package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type CreateRdsOpts struct {
	//
	Name string `json:"name"  required:"true"`
	//
	Datastore *Datastore `json:"datastore" required:"true"`
	//
	Ha *Ha `json:"ha,omitempty"`
	//
	ConfigurationId string `json:"configuration_id,omitempty"`
	//
	Port string `json:"port,omitempty"`
	//
	Password string `json:"password" required:"true"`
	//
	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	//
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	//
	FlavorRef string `json:"flavor_ref" required:"true"`
	//
	Volume *Volume `json:"volume" required:"true"`
	//
	Region string `json:"region" required:"true"`
	//
	AvailabilityZone string `json:"availability_zone" required:"true"`
	//
	VpcId string `json:"vpc_id" required:"true"`
	//
	SubnetId string `json:"subnet_id" required:"true"`
	//
	SecurityGroupId string `json:"security_group_id" required:"true"`
	//
	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
	//
	TimeZone string `json:"time_zone,omitempty"`
}

type Datastore struct {
	//
	Type string `json:"type" required:"true"`
	//
	Version string `json:"version" required:"true"`
}

type Ha struct {
	//
	Mode string `json:"mode" required:"true"`
	//
	ReplicationMode string `json:"replication_mode,omitempty"`
}

type BackupStrategy struct {
	//
	StartTime string `json:"start_time" required:"true"`
	//
	KeepDays int `json:"keep_days,omitempty"`
}

type CreateRdsBuilder interface {
	ToInstancesCreateMap() (map[string]interface{}, error)
}

func (opts CreateRdsOpts) ToInstancesCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateRdsBuilder) (r CreateResult) {
	b, err := opts.ToInstancesCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
