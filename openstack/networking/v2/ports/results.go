package ports

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a port resource.
func (r commonResult) Extract() (*Port, error) {
	var s Port
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "port")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Port.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Port.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Port.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// AddressPair contains the IP Address and the MAC address.
type AddressPair struct {
	IPAddress  string `json:"ip_address,omitempty"`
	MACAddress string `json:"mac_address,omitempty"`
}

// Port represents a Neutron port. See package documentation for a top-level
// description of what this is.
type Port struct {
	// UUID for the port.
	ID string `json:"id"`

	// Network that this port is associated with.
	NetworkID string `json:"network_id"`

	// Human-readable name for the port. Might not be unique.
	Name string `json:"name"`

	// Administrative state of port. If false (down), port does not forward
	// packets.
	AdminStateUp bool `json:"admin_state_up"`

	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// Mac address to use on this port.
	MACAddress string `json:"mac_address"`

	// Specifies IP addresses for the port thus associating the port itself with
	// the subnets where the IP addresses are picked from
	FixedIPs []IP `json:"fixed_ips"`

	// TenantID is the project owner of the port.
	TenantID string `json:"tenant_id"`

	// ProjectID is the project owner of the port.
	ProjectID string `json:"project_id"`

	// Identifies the entity (e.g.: dhcp agent) using this port.
	DeviceOwner string `json:"device_owner"`

	// Specifies the IDs of any security groups associated with a port.
	SecurityGroups []string `json:"security_groups"`

	// Identifies the device (e.g., virtual server) using this port.
	DeviceID string `json:"device_id"`

	// Identifies the list of IP addresses the port will recognize/accept
	AllowedAddressPairs []AddressPair `json:"allowed_address_pairs"`

	// Specifies the extended option (extended attribute) of DHCP.
	ExtraDhcpOpts []ExtraDhcpOpt `json:"extra_dhcp_opts"`
	// Specifies the VIF details. Parameter ovs_hybrid_plug specifies whether the OVS/bridge hybrid mode is used.
	VifDetails VifDetail `json:"binding:vif_details"`
	// Specifies the custom information configured by users. This is an extended attribute.
	Profile interface{} `json:"binding:profile"`
	// Specifies the type of the bound vNIC. The value can be normal or direct.
	// Parameter normal indicates software switching.
	// Parameter direct indicates SR-IOV PCIe passthrough, which is not supported.
	VnicType string `json:"binding:vnic_type"`
	// Specifies the default private network domain name information of the primary NIC.
	// The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	DnsAssignment []DnsAssignment `json:"dns_assignment"`
	// Specifies the default private network DNS name of the primary NIC.
	// The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	DnsName string `json:"dns_name"`
	// Specifies the ID of the instance to which the port belongs, for example, RDS instance ID.
	// The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	InstanceId string `json:"instance_id"`
	// Specifies the type of the instance to which the port belongs, for example, RDS.
	// The system automatically sets this parameter, and you are not allowed to configure or change the parameter value.
	InstanceType string `json:"instance_type"`
	// Specifies whether the security option is enabled for the port.
	// If the option is not enabled, the security group and DHCP snooping do not take effect.
	PortSecurityEnabled bool `json:"port_security_enabled"`
	// Availability zone to which the port belongs.
	ZoneId string `json:"zone_id"`
	// Whether to enable efi
	EnableEfi bool `json:"enable_efi"`
	// The Shared bandwidth ID bound to IPv6
	Ipv6BandwidthId string `json:"ipv6_bandwidth_id"`
}

// VifDetail is an Object specifying the VIF details.
type VifDetail struct {
	// If the value is true, indicating that it is the main network card of the virtual machine.
	PrimaryInterface bool `json:"primary_interface"`
}

// DnsAssignment is an Object specifying the private network domain information.
type DnsAssignment struct {
	// Specifies the hostname.
	Hostname string `json:"hostname"`
	// Specifies the IP address of the port.
	IpAddress string `json:"ip_address"`
	// Specifies the FQDN.
	Fqdn string `json:"fqdn"`
}

// ExtraDhcpOpt is an Object specifying the DHCP extended properties.
type ExtraDhcpOpt struct {
	// Specifies the DHCP option name.
	// Currently, only '51' is supported to indicate the DHCP lease time.
	OptName string `json:"opt_name,omitempty"`
	// Specifies the DHCP option value.
	// When 'OptName' is '51', the parameter format is 'Xh', indicating that the DHCP lease time is X hours.
	// The value range of 'X' is '1~30000' or '-1', '-1' means the DHCP lease time is infinite.
	OptValue string `json:"opt_value,omitempty"`
}

// PortPage is the page returned by a pager when traversing over a collection
// of network ports.
type PortPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of ports has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r PortPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"ports_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a PortPage struct is empty.
func (r PortPage) IsEmpty() (bool, error) {
	is, err := ExtractPorts(r)
	return len(is) == 0, err
}

// ExtractPorts accepts a Page struct, specifically a PortPage struct,
// and extracts the elements into a slice of Port structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPorts(r pagination.Page) ([]Port, error) {
	var s []Port
	err := ExtractPortsInto(r, &s)
	return s, err
}

func ExtractPortsInto(r pagination.Page, v interface{}) error {
	return r.(PortPage).Result.ExtractIntoSlicePtr(v, "ports")
}
