package sni

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Number of records displayed on each page.
	Limit int `q:"limit,omitempty"`
	// Start resource ID of pagination query.
	// If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker,omitempty"`
	// ID of the supplementary network interface. Multiple IDs can be specified for filtering.
	Id []string `q:"id,omitempty"`
	// VPC name, which can be used to filter VPCs.
	Name []string `q:"name,omitempty"`
	// Description of the supplementary network interface.
	// Multiple descriptions can be specified for filtering.
	Description []string `q:"description,omitempty"`
	// MAC address of the supplementary network interface.
	// Multiple MAC addresses can be specified for filtering.
	MacAddress []string `q:"mac_address,omitempty"`
	// ID of elastic network interface that the supplementary network interface is attached to.
	// Multiple IDs can be specified for filtering.
	ParentId []string `q:"parent_id,omitempty"`
	// Private IPv4 address of the supplementary network interface.
	// Multiple private IPv4 addresses can be specified for filtering.
	PrivateIpAddress []string `q:"private_ip_address,omitempty"`
	// Virtual subnet ID of the supplementary network interface.
	// Multiple IDs can be specified for filtering.
	VirSubnetId []string `q:"virsubnet_id,omitempty"`
	// VPC ID of the supplementary network interface.
	// Multiple IDs can be specified for filtering.
	VpcId []string `q:"vpc_id,omitempty"`
}

type SubNetworkInterface struct {
	// Unique identifier of the supplementary network interface
	// The value is in UUID format with hyphens (-).
	Id string `json:"id"`
	// Virtual subnet ID
	// The value must be in standard UUID format.
	VirSubnetId string `json:"virsubnet_id"`
	// Private IPv4 address of the supplementary network interface
	// The value must be within the virtual subnet.
	// If this parameter is left blank, an IP address will be randomly assigned.
	PrivateIpAddress string `json:"private_ip_address"`
	// IPv6 address of the supplementary network interface
	Ipv6IpAddress string `json:"ipv6_ip_address"`
	// MAC address of the supplementary network interface
	// The value is a valid MAC address assigned by the system randomly.
	MacAddress string `json:"mac_address"`
	// Device ID
	// The value must be in standard UUID format.
	ParentDeviceId string `json:"parent_device_id"`
	// ID of the elastic network interface
	// The value must be in standard UUID format.
	ParentId string `json:"parent_id"`
	// Description of the supplementary network interface
	// The value can contain no more than 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description"`
	// VPC ID of the supplementary network interface
	// The value must be in standard UUID format.
	VpcId string `json:"vpc_id"`
	// VLAN ID of the supplementary network interface
	// The value can be from 1 to 4094.
	// Each supplementary network interface of an elastic network interface has a unique VLAN ID.
	VlanId int `json:"vlan_id"`
	// Security group IDs, for example, "security_groups": ["a0608cbf-d047-4f54-8b28-cd7b59853fff"]
	// The default value is the default security group.
	SecurityGroups []string `json:"security_groups"`
	// Tags of the supplementary network interface
	Tags []string `json:"tags"`
	// Project ID of the supplementary network interface
	ProjectId string `json:"project_id"`
	// Creation time of the supplementary network interface
	// The value is a UTC time in the format of yyyy-MM-ddTHH:mmss.
	CreatedAt string `json:"created_at"`
}

// List is used to query supplementary network interfaces.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]SubNetworkInterface, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v3/{project_id}/vpc/sub-network-interfaces
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("sub-network-interfaces") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return InterfacePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSubInterfaces(pages)
}

// InterfacePage is the page returned by a pager when traversing over a
// collection of supplementary network interfaces.
type InterfacePage struct {
	pagination.NewSinglePageBase
}

// ExtractSubInterfaces accepts a Page struct, specifically a InterfacePage struct,
// and extracts the elements into a slice of supplementary network interfaces structs.
func ExtractSubInterfaces(r pagination.NewPage) ([]SubNetworkInterface, error) {
	var s struct {
		SubNetworkInterfaces []SubNetworkInterface `json:"sub_network_interfaces"`
	}
	err := extract.Into(bytes.NewReader((r.(InterfacePage)).Body), &s)
	return s.SubNetworkInterfaces, err
}
