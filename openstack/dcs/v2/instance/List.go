package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListDcsInstanceOpts struct {
	InstanceId     string `q:"instance_id"`
	IncludeFailure string `q:"include_failure"`
	IncludeDelete  string `q:"include_delete"`
	Name           string `q:"name"`
	Offset         int    `q:"offset"`
	Limit          int    `q:"limit"`
	Status         string `q:"status"`
	NameEqual      string `q:"name_equal"`
	Tags           string `q:"tags"`
	Ip             string `q:"ip"`
	Capacity       string `q:"capacity"`
}

func List(client *golangsdk.ServiceClient, opts ListDcsInstanceOpts) (*ListDcsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListDcsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListDcsResponse struct {
	InstanceNum int                   `json:"instance_num"`
	Instances   []ListDcsInstanceResp `json:"instances"`
}

type ListDcsInstanceResp struct {
	PublicIpId         string             `json:"publicip_id"`
	VpcName            string             `json:"vpc_name"`
	ChargingMode       int                `json:"charging_mode"`
	VpcId              string             `json:"vpc_id"`
	SubnetId           string             `json:"subnet_id"`
	SecurityGroupId    string             `json:"security_group_id"`
	CreatedAt          string             `json:"created_at"`
	UpdatedAt          string             `json:"updated_at"`
	EnableSsl          bool               `json:"enable_ssl"`
	MaxMemory          int                `json:"max_memory"`
	UsedMemory         int                `json:"used_memory"`
	PublicIpAddress    string             `json:"publicip_address"`
	Capacity           float64            `json:"capacity"`
	CapacityMinor      string             `json:"capacity_minor"`
	OrderId            string             `json:"order_id"`
	MaintainBegin      string             `json:"maintain_begin"`
	MaintainEnd        string             `json:"maintain_end"`
	Engine             string             `json:"engine"`
	EngineVersion      string             `json:"engine_version"`
	ServiceUpgrade     bool               `json:"service_upgrade"`
	NoPasswordAccess   string             `json:"no_password_access"`
	ServiceTaskId      string             `json:"service_task_id"`
	Ip                 string             `json:"ip"`
	AccessUser         string             `json:"access_user"`
	InstanceId         string             `json:"instance_id"`
	EnablePublicIp     bool               `json:"enable_publicip"`
	Port               int                `json:"port"`
	UserId             string             `json:"user_id"`
	UserName           string             `json:"user_name"`
	DomainName         string             `json:"domain_name"`
	ReadOnlyDomainName string             `json:"readonly_domain_name"`
	Name               string             `json:"name"`
	SpecCode           string             `json:"spec_code"`
	Status             string             `json:"status"`
	Tags               []tags.ResourceTag `json:"tags"`
	Description        string             `json:"description"`
	CpuType            string             `json:"cpu_type"`
	AzCodes            []string           `json:"az_codes"`
	Features           Features           `json:"features"`
	SubStatus          string             `json:"sub_status"`
}

type Features struct {
	SupportAcl                 bool `json:"support_acl"`
	SupportTransparentClientIp bool `json:"support_transparent_client_ip"`
	SupportSsl                 bool `json:"support_ssl"`
	SupportAuditLog            bool `json:"support_audit_log"`
}
