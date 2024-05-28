package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	Name             string                    `json:"name" required:"true"`
	Engine           string                    `json:"engine" required:"true"`
	EngineVersion    string                    `json:"engine_version,omitempty"`
	Capacity         float64                   `json:"capacity" required:"true"`
	SpecCode         string                    `json:"spec_code" required:"true"`
	AzCodes          []string                  `json:"az_codes" required:"true"`
	VpcId            string                    `json:"vpc_id" required:"true"`
	SubnetId         string                    `json:"subnet_id" required:"true"`
	SecurityGroupId  string                    `json:"security_group_id,omitempty"`
	PublicIpId       string                    `json:"publicip_id,omitempty"`
	Description      string                    `json:"description,omitempty"`
	EnableSsl        *bool                     `json:"enable_ssl,omitempty"`
	PrivateIp        string                    `json:"private_ip,omitempty"`
	InstanceNum      int                       `json:"instance_num,omitempty"`
	MaintainBegin    string                    `json:"maintain_begin,omitempty"`
	MaintainEnd      string                    `json:"maintain_end,omitempty"`
	Password         string                    `json:"password,omitempty"`
	NoPasswordAccess *bool                     `json:"no_password_access,omitempty"`
	BssParam         DcsBssParam               `json:"bss_param,omitempty"`
	BackupPolicy     *InstanceBackupPolicyOpts `json:"instance_backup_policy,omitempty"`
	Tags             []tags.ResourceTag        `json:"tags,omitempty"`
	AccessUser       string                    `json:"access_user,omitempty"`
	EnablePublicIp   *bool                     `json:"enable_publicip,omitempty"`
	Port             int                       `json:"port,omitempty"`
	RenameCommands   RenameCommand             `json:"rename_commands,omitempty"`
	TemplateId       string                    `json:"template_id,omitempty"`
}

type DcsBssParam struct {
	ChargingMode string `json:"charging_mode" required:"true"`
	PeriodType   string `json:"period_type,omitempty"`
	PeriodNum    int    `json:"period_num,omitempty"`
	IsAutoRenew  string `json:"is_auto_renew,omitempty"`
	IsAutoPay    string `json:"is_auto_pay,omitempty"`
}

type InstanceBackupPolicyOpts struct {
	BackupType           string      `json:"backup_type" required:"true"`
	SaveDays             int         `json:"save_days,omitempty"`
	PeriodicalBackupPlan *BackupPlan `json:"periodical_backup_plan,omitempty"`
}

type BackupPlan struct {
	TimezoneOffset string `json:"timezone_offset,omitempty"`
	BackupAt       []int  `json:"backup_at" required:"true"`
	PeriodType     string `json:"period_type" required:"true"`
	BeginAt        string `json:"begin_at" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) ([]DcsCreateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res []DcsCreateResp

	err = extract.IntoSlicePtr(raw.Body, &res, "instances")
	return res, err
}

type DcsCreateResp struct {
	InstanceID   string `json:"instance_id"`
	InstanceName string `json:"instance_name"`
}
