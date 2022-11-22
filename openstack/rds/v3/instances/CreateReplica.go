package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type CreateReplicaOpts struct {
	//
	Name string `json:"name"  required:"true"`
	//
	ReplicaOfId string `json:"replica_of_id" required:"true"`
	//
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	//
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	//
	FlavorRef string `json:"flavor_ref" required:"true"`
	//
	Volume *Volume `json:"volume" required:"true"`
	//
	Region string `json:"region,omitempty"`
	//
	AvailabilityZone string `json:"availability_zone" required:"true"`
	//
	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
}

type Volume struct {
	//
	Type string `json:"type" required:"true"`
	//
	Size int `json:"size,omitempty"`
}

type ChargeInfo struct {
	//
	ChargeMode string `json:"charge_mode" required:"true"`
	//
	PeriodType string `json:"period_type,omitempty"`
	//
	PeriodNum int `json:"period_num,omitempty"`
	//
	IsAutoRenew string `json:"is_auto_renew,omitempty"`
	//
	IsAutoPay string `json:"is_auto_pay,omitempty"`
}

type CreateReplicaBuilder interface {
	ToCreateReplicaMap() (map[string]interface{}, error)
}

func (opts CreateReplicaOpts) ToCreateReplicaMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateReplica(client *golangsdk.ServiceClient, opts CreateReplicaBuilder) (r CreateResult) {
	b, err := opts.ToCreateReplicaMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})

	return
}
