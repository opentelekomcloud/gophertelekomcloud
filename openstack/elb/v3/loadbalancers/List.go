package loadbalancers

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Marker                   string   `q:"marker,omitempty"`
	Limit                    int      `q:"limit,omitempty"`
	PageReverse              *bool    `q:"page_reverse,omitempty"`
	ID                       []string `q:"id,omitempty"`
	Name                     []string `q:"name,omitempty"`
	Description              []string `q:"description,omitempty"`
	AdminStateUp             *bool    `q:"admin_state_up,omitempty"`
	ProvisioningStatus       []string `q:"provisioning_status,omitempty"`
	OperatingStatus          []string `q:"operating_status,omitempty"`
	Guaranteed               *bool    `q:"guaranteed,omitempty"`
	VpcID                    []string `q:"vpc_id,omitempty"`
	VipPortID                []string `q:"vip_port_id,omitempty"`
	VipAddress               []string `q:"vip_address,omitempty"`
	VipSubnetCidrID          []string `q:"vip_subnet_cidr_id,omitempty"`
	IpV6VipPortID            []string `q:"ipv6_vip_port_id,omitempty"`
	IpV6VipAddress           []string `q:"ipv6_vip_address,omitempty"`
	IpV6VipSubnetID          []string `q:"ipv6_vip_virsubnet_id,omitempty"`
	Eips                     []string `q:"eips,omitempty"`
	PublicIps                []string `q:"publicips,omitempty"`
	AvailabilityZoneList     []string `q:"availability_zone_list,omitempty"`
	L4FlavorID               []string `q:"l4_flavor_id,omitempty"`
	L4ScaleFlavorID          []string `q:"l4_scale_flavor_id,omitempty"`
	L7FlavorID               []string `q:"l7_flavor_id,omitempty"`
	L7ScaleFlavorID          []string `q:"l7_scale_flavor_id,omitempty"`
	MemberDeviceID           []string `q:"member_device_id,omitempty"`
	MemberAddress            []string `q:"member_address,omitempty"`
	EnterpriseProjectID      []string `q:"enterprise_project_id,omitempty"`
	IPVersion                []int    `q:"ip_version,omitempty"`
	DeletionProtectionEnable *bool    `q:"deletion_protection_enable,omitempty"`
	ElbSubnetType            []string `q:"elb_virsubnet_type,omitempty"`
	ProtectionStatus         []string `q:"protection_status,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]LoadBalancer, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("loadbalancers").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return LoadbalancerPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractLoadbalancers(pages)
}

type LoadbalancerPage struct {
	pagination.NewPageResult
}

func (p LoadbalancerPage) NewNextPageURL() (string, error) {
	var res struct {
		PageInfo PageInfo `json:"page_info"`
	}
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	if res.PageInfo.NextMarker == "" {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", res.PageInfo.NextMarker)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p LoadbalancerPage) NewIsEmpty() (bool, error) {
	loadbalancers, err := ExtractLoadbalancers(p)
	return len(loadbalancers) == 0, err
}

func ExtractLoadbalancers(page pagination.NewPage) ([]LoadBalancer, error) {
	var res struct {
		Loadbalancers []LoadBalancer `json:"loadbalancers"`
	}
	err := extract.Into(bytes.NewReader(page.(LoadbalancerPage).Body), &res)
	return res.Loadbalancers, err
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
