package subnets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name          string         `json:"name" required:"true"`
	Description   *string        `json:"description,omitempty"`
	EnableIpv6    *bool          `json:"ipv6_enable,omitempty"`
	EnableDHCP    *bool          `json:"dhcp_enable,omitempty"`
	PrimaryDNS    string         `json:"primary_dns,omitempty"`
	SecondaryDNS  string         `json:"secondary_dns,omitempty"`
	DNSList       []string       `json:"dnsList,omitempty"`
	ExtraDHCPOpts []ExtraDHCPOpt `json:"extra_dhcp_opts,omitempty"`
}

func Update(client *golangsdk.ServiceClient, vpcID, id string, opts UpdateOpts) (*Subnet, error) {
	b, err := build.RequestBody(opts, "subnet")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(client.ProjectID, "vpcs", vpcID, "subnets", id), b, nil, &golangsdk.RequestOpts{
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
