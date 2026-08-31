package subnets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name             string         `json:"name" required:"true"`
	Description      string         `json:"description,omitempty"`
	CIDR             string         `json:"cidr" required:"true"`
	GatewayIP        string         `json:"gateway_ip" required:"true"`
	EnableIpv6       *bool          `json:"ipv6_enable,omitempty"`
	EnableDHCP       *bool          `json:"dhcp_enable,omitempty"`
	PrimaryDNS       string         `json:"primary_dns,omitempty"`
	SecondaryDNS     string         `json:"secondary_dns,omitempty"`
	DNSList          []string       `json:"dnsList,omitempty"`
	AvailabilityZone string         `json:"availability_zone,omitempty"`
	VpcID            string         `json:"vpc_id" required:"true"`
	ExtraDHCPOpts    []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Subnet, error) {
	b, err := build.RequestBody(opts, "subnet")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "subnets"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Subnet Subnet `json:"subnet"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.Subnet, err
}
