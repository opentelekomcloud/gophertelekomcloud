package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	ProjectID                string             `json:"project_id,omitempty"`
	Name                     string             `json:"name,omitempty"`
	Description              string             `json:"description,omitempty"`
	VipAddress               string             `json:"vip_address,omitempty"`
	VipSubnetCidrID          string             `json:"vip_subnet_cidr_id,omitempty"`
	IpV6VipSubnetID          string             `json:"ipv6_vip_virsubnet_id,omitempty"`
	Provider                 string             `json:"provider,omitempty"`
	L4Flavor                 string             `json:"l4_flavor_id,omitempty"`
	L7Flavor                 string             `json:"l7_flavor_id,omitempty"`
	Guaranteed               *bool              `json:"guaranteed,omitempty"`
	VpcID                    string             `json:"vpc_id,omitempty"`
	AvailabilityZoneList     []string           `json:"availability_zone_list" required:"true"`
	EnterpriseProjectID      string             `json:"enterprise_project_id,omitempty"`
	Tags                     []tags.ResourceTag `json:"tags,omitempty"`
	AdminStateUp             *bool              `json:"admin_state_up,omitempty"`
	IPV6Bandwidth            *BandwidthRef      `json:"ipv6_bandwidth,omitempty"`
	PublicIpIDs              []string           `json:"publicip_ids,omitempty"`
	PublicIp                 *PublicIp          `json:"publicip,omitempty"`
	ElbSubnetIDs             []string           `json:"elb_virsubnet_ids,omitempty"`
	IpTargetEnable           *bool              `json:"ip_target_enable,omitempty"`
	DeletionProtectionEnable *bool              `json:"deletion_protection_enable,omitempty"`
	WafFailureAction         string             `json:"waf_failure_action,omitempty"`
	ChargeMode               string             `json:"charge_mode,omitempty"`
	ProtectionStatus         string             `json:"protection_status,omitempty"`
	ProtectionReason         string             `json:"protection_reason,omitempty"`
}

type BandwidthRef struct {
	ID string `json:"id" required:"true"`
}

type PublicIp struct {
	IpVersion   int       `json:"ip_version,omitempty"`
	NetworkType string    `json:"network_type" required:"true"`
	BillingInfo string    `json:"billing_info,omitempty"`
	Description string    `json:"description,omitempty"`
	Bandwidth   Bandwidth `json:"bandwidth" required:"true"`
}

type Bandwidth struct {
	Name        string `json:"name,omitempty"`
	Size        int    `json:"size,omitempty"`
	ChargeMode  string `json:"charge_mode,omitempty"`
	ShareType   string `json:"share_type,omitempty"`
	BillingInfo string `json:"billing_info,omitempty"`
	ID          string `json:"id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*LoadBalancer, error) {
	b, err := build.RequestBody(opts, "loadbalancer")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("loadbalancers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		LoadBalancer LoadBalancer `json:"loadbalancer"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.LoadBalancer, nil
}
