package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// LoadBalancer is the primary load balancing configuration object that
// specifies the virtual IP address on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type LoadBalancer struct {
	// The unique ID for the LoadBalancer.
	ID string `json:"id"`

	// Human-readable description for the Loadbalancer.
	Description string `json:"description"`

	// The provisioning status of the LoadBalancer.
	// This value is ACTIVE, PENDING_CREATE or ERROR.
	ProvisioningStatus string `json:"provisioning_status"`

	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// The name of the provider.
	Provider string `json:"provider"`

	// Pools are the pools related to this Loadbalancer.
	Pools []structs.ResourceRef `json:"pools"`

	// Listeners are the listeners related to this Loadbalancer.
	Listeners []structs.ResourceRef `json:"listeners"`

	// The operating status of the LoadBalancer. This value is ONLINE or OFFLINE.
	OperatingStatus string `json:"operating_status"`

	// The IP address of the Loadbalancer.
	VipAddress string `json:"vip_address"`

	// The UUID of the subnet on which to allocate the virtual IP for the
	// Loadbalancer address.
	VipSubnetCidrID string `json:"vip_subnet_cidr_id"`

	// Human-readable name for the LoadBalancer. Does not have to be unique.
	Name string `json:"name"`

	// Owner of the LoadBalancer.
	ProjectID string `json:"project_id"`

	// The UUID of the port associated with the IP address.
	VipPortID string `json:"vip_port_id"`

	// The UUID of a flavor if set.
	Tags []tags.ResourceTag `json:"tags"`

	// Guaranteed.
	Guaranteed bool `json:"guaranteed"`

	// The VPC ID.
	VpcID string `json:"vpc_id"`

	// EIP Info.
	Eips []EipInfo `json:"eips"`

	// IpV6 Vip Address.
	IpV6VipAddress string `json:"ipv6_vip_address"`

	// IpV6 Vip VirSubnet ID.
	IpV6VipSubnetID string `json:"ipv6_vip_virsubnet_id"`

	// IpV6 Vip Port ID.
	IpV6VipPortID string `json:"ipv6_vip_port_id"`

	// Availability Zone List.
	AvailabilityZoneList []string `json:"availability_zone_list"`

	// L4 Flavor ID.
	L4FlavorID string `json:"l4_flavor_id"`

	// L4 Scale Flavor ID.
	L4ScaleFlavorID string `json:"l4_scale_flavor_id"`

	// L7 Flavor ID.
	L7FlavorID string `json:"l7_flavor_id"`

	// L7 Scale Flavor ID.
	L7ScaleFlavorID string `json:"l7_scale_flavor_id"`

	// Public IP Info.
	PublicIps []PublicIpInfo `json:"publicips"`

	// Elb VirSubnet IDs.
	ElbSubnetIDs []string `json:"elb_virsubnet_ids"`

	// Elb VirSubnet Type.
	ElbSubnetType string `json:"elb_virsubnet_type"`

	// Ip Target Enable.
	IpTargetEnable bool `json:"ip_target_enable"`

	// Frozen Scene.
	FrozenScene string `json:"frozen_scene"`

	// Ipv6 Bandwidth.
	IpV6Bandwidth BandwidthRef `json:"ipv6_bandwidth"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	DeletionProtectionEnable bool                     `json:"deletion_protection_enable"`
	EnterpriseProjectID      string                   `json:"enterprise_project_id"`
	BillingInfo              string                   `json:"billing_info"`
	GlobalEips               []GlobalEipInfo          `json:"global_eips"`
	Autoscaling              AutoscalingRef           `json:"autoscaling"`
	PublicBorderGroup        string                   `json:"public_border_group"`
	WafFailureAction         string                   `json:"waf_failure_action"`
	ChargeMode               string                   `json:"charge_mode"`
	ProtectionStatus         string                   `json:"protection_status"`
	ProtectionReason         string                   `json:"protection_reason"`
	LoadbalancerType         string                   `json:"loadbalancer_type"`
	GwFlavorID               string                   `json:"gw_flavor_id"`
	InstanceType             string                   `json:"instance_type"`
	InstanceID               string                   `json:"instance_id"`
	LogGroupID               string                   `json:"log_group_id"`
	LogTopicID               string                   `json:"log_topic_id"`
	CustomQosLimit           CustomQosLimit           `json:"custom_qos_limit"`
	ServiceLBMode            string                   `json:"service_lb_mode"`
	ProxyProtocolExtensions  []ProxyProtocolExtension `json:"proxy_protocol_extensions"`
}

type EipInfo struct {
	// Eip ID
	EipID string `json:"eip_id"`
	// Eip Address
	EipAddress string `json:"eip_address"`
	// Eip Address
	IpVersion int `json:"ip_version"`
}

type PublicIpInfo struct {
	// Public IP ID
	PublicIpID string `json:"publicip_id"`
	// Public IP Address
	PublicIpAddress string `json:"publicip_address"`
	// IP Version
	IpVersion int `json:"ip_version"`
}

type GlobalEipInfo struct {
	GlobalEipID      string `json:"global_eip_id"`
	GlobalEipAddress string `json:"global_eip_address"`
	IpVersion        int    `json:"ip_version"`
}

type AutoscalingRef struct {
	Enable        bool   `json:"enable"`
	MinL7FlavorID string `json:"min_l7_flavor_id"`
}

type CustomQosLimit struct {
	L4 QosLimit `json:"l4"`
	L7 QosLimit `json:"l7"`
}

type QosLimit struct {
	Connection int `json:"connection"`
	CPS        int `json:"cps"`
}

type ProxyProtocolExtension struct {
	VipAddress     string    `json:"vip_address"`
	IpV6VipAddress string    `json:"ipv6_vip_address"`
	Extension      Extension `json:"extension"`
}

type Extension struct {
	EpID        string `json:"ep_id"`
	EpServiceID string `json:"ep_service_id"`
}

// StatusTree represents the status of a loadbalancer.
type StatusTree struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*LoadBalancer, error) {
	s := new(LoadBalancer)
	err := r.ExtractIntoStructPtr(s, "loadbalancer")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetStatusesResult represents the result of a GetStatuses operation.
// Call its Extract method to interpret it as a StatusTree.
type GetStatusesResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r GetStatusesResult) Extract() (*StatusTree, error) {
	s := new(StatusTree)
	err := r.ExtractIntoStructPtr(s, "statuses")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a LoadBalancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
