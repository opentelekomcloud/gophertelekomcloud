package instances

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateReplicaOpts struct {
	// Specifies the DB instance name.
	// DB instances of the same type can have same names under the same tenant.
	// The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"  required:"true"`
	// Specifies the DB instance ID, which is used to create a read replica.
	ReplicaOfId string `json:"replica_of_id" required:"true"`
	// Specifies the key ID for disk encryption. The default value is empty.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Specifies the key ID for disk encryption. The default value is empty.
	DiskEncryptionId string `json:"disk_encryption_id,omitempty"`
	// Specifies the specification code. The value cannot be empty.
	FlavorRef string `json:"flavor_ref" required:"true"`
	// Specifies the volume information.
	Volume *Volume `json:"volume" required:"true"`
	// Specifies the region ID. Currently, read replicas can be created only in the same region as that of the primary DB instance.
	// The value cannot be empty.
	Region string `json:"region"`
	// Specifies the AZ ID.
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// Specifies the billing information, which is pay-per-use. By default, pay-per-use is used.
	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
}

type Volume struct {
	// Indicates the volume type.
	// Its value can be any of the following and is case-sensitive:
	// COMMON: indicates the SATA type.
	// ULTRAHIGH: indicates the SSD type.
	Type string `json:"type" required:"true"`
	// Indicates the volume size.
	// Its value range is from 40 GB to 4000 GB. The value must be a multiple of 10.
	Size int `json:"size,omitempty"`
}

type ChargeInfo struct {
	// Indicates the billing information, which is pay-per-use.
	ChargeMode string `json:"charge_mode" required:"true"`

	PeriodType  string `json:"period_type,omitempty"`
	PeriodNum   int    `json:"period_num,omitempty"`
	IsAutoRenew string `json:"is_auto_renew,omitempty"`
	IsAutoPay   string `json:"is_auto_pay,omitempty"`
}

// CreateReplica (Only Microsoft SQL Server 2017 EE supports read replicas and does not support single DB instances.)
func CreateReplica(client *golangsdk.ServiceClient, opts CreateReplicaOpts) (*CreateRds, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*CreateRds, error) {
	if err != nil {
		return nil, err
	}

	var res CreateRds
	err = extract.Into(raw.Body, &res)
	return &res, err
}
