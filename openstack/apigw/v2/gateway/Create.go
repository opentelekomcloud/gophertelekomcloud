package gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	Description                  string             `json:"description,omitempty"`
	MaintainBegin                string             `json:"maintain_begin,omitempty"`
	MaintainEnd                  string             `json:"maintain_end,omitempty"`
	InstanceName                 string             `json:"instance_name" required:"true"`
	InstanceID                   string             `json:"instance_id,omitempty"`
	SpecID                       string             `json:"spec_id" required:"true"`
	VpcID                        string             `json:"vpc_id" required:"true"`
	SubnetID                     string             `json:"subnet_id" required:"true"`
	SecGroupID                   string             `json:"security_group_id" required:"true"`
	AvailableZoneIDs             []string           `json:"available_zone_ids" required:"true"`
	BandwidthSize                *int               `json:"bandwidth_size,omitempty"`
	BandwidthChargingMode        string             `json:"bandwidth_charging_mode,omitempty"`
	LoadBalancerProvider         string             `json:"loadbalancer_provider" required:"true"`
	Tags                         []tags.ResourceTag `json:"tags"`
	VpcepServiceName             string             `json:"vpcep_service_name,omitempty"`
	IngressBandwidthSize         *int               `json:"ingress_bandwidth_size,omitempty"`
	IngressBandwidthChargingMode string             `json:"ingress_bandwidth_charging_mode,omitempty"`
	EipId                        string             `json:"eip_id,omitempty"`
	Ipv6Enable                   bool               `json:"ipv6_enable,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*GatewayResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res GatewayResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GatewayResp struct {
	InstanceID string `json:"instance_id"`
	Message    string `json:"message"`
	JobID      string `json:"job_id"`
}
