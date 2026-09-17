package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name                     string        `json:"name,omitempty"`
	AdminStateUp             *bool         `json:"admin_state_up,omitempty"`
	Description              *string       `json:"description,omitempty"`
	IpV6VipSubnetID          *string       `json:"ipv6_vip_virsubnet_id,omitempty"`
	VipSubnetCidrID          *string       `json:"vip_subnet_cidr_id,omitempty"`
	VipAddress               string        `json:"vip_address,omitempty"`
	L4Flavor                 string        `json:"l4_flavor_id,omitempty"`
	L7Flavor                 string        `json:"l7_flavor_id,omitempty"`
	IpV6Bandwidth            *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	IpTargetEnable           *bool         `json:"ip_target_enable,omitempty"`
	ElbSubnetIDs             []string      `json:"elb_virsubnet_ids,omitempty"`
	DeletionProtectionEnable *bool         `json:"deletion_protection_enable,omitempty"`
	WafFailureAction         string        `json:"waf_failure_action,omitempty"`
	ProtectionStatus         string        `json:"protection_status,omitempty"`
	ProtectionReason         string        `json:"protection_reason,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*LoadBalancer, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("loadbalancers", id).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "loadbalancer")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
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
