package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToLoadbalancerListQuery() (string, error)
}

type ListOpts struct {
	ID                   []string `q:"id"`
	Name                 []string `q:"name"`
	Description          []string `q:"description"`
	ProvisioningStatus   []string `q:"provisioning_status"`
	OperatingStatus      []string `q:"operating_status"`
	VpcID                []string `q:"vpc_id"`
	VipPortID            []string `q:"vip_port_id"`
	VipAddress           []string `q:"vip_address"`
	VipSubnetCidrID      []string `q:"vip_subnet_cidr_id"`
	L4FlavorID           []string `q:"l4_flavor_id"`
	L4ScaleFlavorID      []string `q:"l4_scale_flavor_id"`
	AvailabilityZoneList []string `q:"availability_zone_list"`
	L7FlavorID           []string `q:"l7_flavor_id"`
	L7ScaleFlavorID      []string `q:"l7_scale_flavor_id"`
	Limit                int      `q:"limit"`
	Marker               string   `q:"marker"`
}

// ToLoadbalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadbalancerListQuery() (string, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(client)
	if opts != nil {
		query, err := opts.ToLoadbalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.Pager{
		Client:     client,
		InitialURL: url,
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return LoadbalancerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
		},
	}
}

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

type BandwidthRef struct {
	// Share Bandwidth ID
	ID string `json:"id" required:"true"`
}

type PublicIp struct {
	// IP Version.
	IpVersion int `json:"ip_version,omitempty"`

	// Network Type
	NetworkType string `json:"network_type" required:"true"`

	// Billing Info.
	BillingInfo string `json:"billing_info,omitempty"`

	// Description.
	Description string `json:"description,omitempty"`

	// Bandwidth
	Bandwidth Bandwidth `json:"bandwidth" required:"true"`
}

type Bandwidth struct {
	// Name
	Name string `json:"name" required:"true"`

	// Size
	Size int `json:"size" required:"true"`

	// Charge Mode
	ChargeMode string `json:"charge_mode" required:"true"`

	// Share Type
	ShareType string `json:"share_type" required:"true"`
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
	_, r.Err = client.Post(rootURL(client), b, &r.Body, nil)
	return
}

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(resourceURL(client, id), &r.Body, nil)
	return
}

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
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular LoadBalancer based on its
// unique ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(resourceURL(client, id), nil)
	return
}

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(client *golangsdk.ServiceClient, id string) (r GetStatusesResult) {
	_, r.Err = client.Get(statusURL(client, id), &r.Body, nil)
	return
}
