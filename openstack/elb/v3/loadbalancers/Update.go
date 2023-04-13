package loadbalancers

import "github.com/opentelekomcloud/gophertelekomcloud"

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToLoadBalancerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options' struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name string `json:"name,omitempty"`

	// Human-readable description for the Loadbalancer.
	Description *string `json:"description,omitempty"`

	// The administrative state of the Loadbalancer. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// The IP address of the Loadbalancer.
	VipAddress string `json:"vip_address,omitempty"`

	// The network on which to allocate the Loadbalancer's address.
	VipSubnetCidrID *string `json:"vip_subnet_cidr_id,omitempty"`

	// The V6 network on which to allocate the Loadbalancer's address.
	IpV6VipSubnetID *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// The UUID of a l4 flavor.
	L4Flavor string `json:"l4_flavor_id,omitempty"`

	// The UUID of a l7 flavor.
	L7Flavor string `json:"l7_flavor_id,omitempty"`

	// IPv6 Bandwidth.
	IpV6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	// ELB VirSubnet IDs.
	ElbSubnetIDs []string `json:"elb_virsubnet_ids,omitempty"`

	// IP Target Enable.
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// Specifies whether to enable deletion protection for the load balancer.
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
}

// ToLoadBalancerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "loadbalancer")
}

// Update is an operation which modifies the attributes of the specified
// LoadBalancer.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToLoadBalancerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("loadbalancers", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}
