package securitygrouprules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	SecurityGroupID      string `json:"security_group_id" required:"true"`
	Description          string `json:"description,omitempty"`
	Direction            string `json:"direction" required:"true"`
	EtherType            string `json:"ethertype,omitempty"`
	Protocol             string `json:"protocol,omitempty"`
	PortRangeMin         *int   `json:"port_range_min,omitempty"`
	PortRangeMax         *int   `json:"port_range_max,omitempty"`
	RemoteIPPrefix       string `json:"remote_ip_prefix,omitempty"`
	RemoteGroupID        string `json:"remote_group_id,omitempty"`
	RemoteAddressGroupID string `json:"remote_address_group_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SecurityGroupRule, error) {
	b, err := build.RequestBody(opts, "security_group_rule")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(client.ProjectID, "security-group-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		SecurityGroupRule SecurityGroupRule `json:"security_group_rule"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.SecurityGroupRule, nil
}
