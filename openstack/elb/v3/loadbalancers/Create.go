package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToLoadBalancerCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options' struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name string `json:"name,omitempty"`

	// Human-readable description for the Loadbalancer.
	Description string `json:"description,omitempty"`

	// The IP address of the Loadbalancer.
	VipAddress string `json:"vip_address,omitempty"`

	// The network on which to allocate the Loadbalancer's address.
	VipSubnetCidrID string `json:"vip_subnet_cidr_id,omitempty"`

	// The V6 network on which to allocate the Loadbalancer's address.
	IpV6VipSubnetID string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// The UUID of a l4 flavor.
	L4Flavor string `json:"l4_flavor_id,omitempty"`

	// Guaranteed.
	Guaranteed *bool `json:"guaranteed,omitempty"`

	// The VPC ID.
	VpcID string `json:"vpc_id,omitempty"`

	// Availability Zone List.
	AvailabilityZoneList []string `json:"availability_zone_list" required:"true"`

	// The tags of the Loadbalancer.
	Tags []tags.ResourceTag `json:"tags,omitempty"`

	// The administrative state of the Loadbalancer. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// The UUID of a l7 flavor.
	L7Flavor string `json:"l7_flavor_id,omitempty"`

	// IPv6 Bandwidth.
	IPV6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	// Public IP IDs.
	PublicIpIDs []string `json:"publicip_ids,omitempty"`

	// Public IP.
	PublicIp *PublicIp `json:"publicip,omitempty"`

	// ELB VirSubnet IDs.
	ElbSubnetIDs []string `json:"elb_virsubnet_ids,omitempty"`

	// IP Target Enable.
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
}

// ToLoadBalancerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "loadbalancer")
}

// Create is an operation which provisions a new loadbalancer based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToLoadBalancerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("loadbalancers"), b, &r.Body, nil)
	return
}
