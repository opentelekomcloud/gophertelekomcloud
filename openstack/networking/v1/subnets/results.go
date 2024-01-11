package subnets

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Subnet struct {
	// ID is the unique identifier for the subnet.
	ID string `json:"id"`

	// Name is the human-readable name for the subnet. It does not have to be
	// unique.
	Name string `json:"name"`

	// Provides supplementary information about the subnet.
	Description string `json:"description"`

	// Specifies the network segment on which the subnet resides.
	CIDR string `json:"cidr"`

	// Specifies the IP address list of DNS servers on the subnet.
	DNSList []string `json:"dnsList"`

	// Status indicates whether a subnet is currently operational.
	Status string `json:"status"`

	// Specifies the gateway of the subnet.
	GatewayIP string `json:"gateway_ip"`

	// Specifies whether the DHCP function is enabled for the subnet.
	EnableDHCP bool `json:"dhcp_enable"`

	// Specifies whether an IPv6 subnet can be created.
	EnableIpv6 bool `json:"ipv6_enable"`

	// Specifies the IP address of DNS server 1 on the subnet.
	PrimaryDNS string `json:"primary_dns"`

	// Specifies the IP address of DNS server 2 on the subnet.
	SecondaryDNS string `json:"secondary_dns"`

	// Identifies the availability zone (AZ) to which the subnet belongs.
	AvailabilityZone string `json:"availability_zone"`

	// Specifies the ID of the VPC to which the subnet belongs.
	VpcID string `json:"vpc_id"`

	// Specifies the subnet ID.
	SubnetID string `json:"neutron_subnet_id"`

	// Specifies the network ID.
	NetworkID string `json:"neutron_network_id"`

	// Specifies the extra dhcp opts.
	ExtraDHCPOpts []ExtraDHCP `json:"extra_dhcp_opts"`
}

type ExtraDHCP struct {
	OptName  string `json:"opt_name"`
	OptValue string `json:"opt_value"`
}

// SubnetPage is the page returned by a pager when traversing over a
// collection of subnets.
type SubnetPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of subnets has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r SubnetPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"subnets_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SubnetPage struct is empty.
func (r SubnetPage) IsEmpty() (bool, error) {
	is, err := ExtractSubnets(r)
	return len(is) == 0, err
}

// ExtractSubnets accepts a Page struct, specifically a SubnetPage struct,
// and extracts the elements into a slice of Subnet structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractSubnets(r pagination.Page) ([]Subnet, error) {
	var s []Subnet
	err := (r.(SubnetPage)).ExtractIntoSlicePtr(&s, "subnets")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a Subnet.
func (r commonResult) Extract() (*Subnet, error) {
	s := new(Subnet)
	err := r.ExtractIntoStructPtr(s, "subnet")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Subnet.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Subnet.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Subnet.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
