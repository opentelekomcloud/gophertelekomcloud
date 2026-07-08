package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Whether to only send the check request.
	// The value can be:
	// true: A check request will be sent and no IP address group will be created.
	// 	Check items include mandatory parameters, request format, and constraints.
	// 	If the check fails, an error will be returned.
	// 	If the check succeeds, response code 202 will be returned.
	// false: A request will be sent and an IP address group will be created.
	DryRun bool `json:"dry_run,omitempty"`
	// Request body for creating an IP address group.
	AddressGroup AddressGroupOptions `json:"address_group" required:"true"`
}

type AddressGroupOptions struct {
	// IP Address group name.
	// The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name" required:"true"`
	// Description about the IP Address group.
	// The value can contain up to 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// IP address version of an IP address group.
	// 4: IPv4 address group
	// 6: IPv6 address group
	IpVersion int `json:"ip_version" required:"true"`
	// Enterprise project ID.
	// When creating an IP address group, associate an enterprise project ID with the IP address group.
	// The project ID can be 0 or a string that contains a maximum of 36 characters in UUID format with hyphens (-).
	// 0 indicates the default enterprise project.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Maximum number of IP address entries in an IP address group.
	// Range: 0 to 20
	// Default: 0
	MaxCapacity int `json:"max_capacity,omitempty"`
	// IP address entries in an IP address group. Both IPv4 and IPv6 address entries are supported.
	// The default maximum number of IP address sets, including IP addresses, IP address ranges, and CIDR blocks, in an IP address group is 20.
	// The ip_set and ip_extra_set parameters cannot be both specified in a request.
	// If you need to add remarks for IP address entries, use the ip_extra_set parameter.
	// An IP address entry in an IP address set can be:
	// 		A single IP address, for example, 192.168.21.25
	// 		An IP address range, for example, 192.168.21.25-192.168.21.30
	// 		A CIDR block, for example, 192.168.21.0/24
	IpSet []string `json:"ip_set,omitempty"`
	// IP address entries in an IP address group and their remarks.
	// There are multiple IP address entry objects in an IP address group.
	// Each IP address entry object contains one IP address entry and its remarks.
	// An IP address entry is the same as that in ip_set.
	// The remarks can describe the usage of the IP address entry.
	IpExtraSet []IpExtraSetOption `json:"ip_extra_set,omitempty"`
}

// This function is used to create an IP address group.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AddressGroup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vpc/address-groups
	raw, err := client.Post(client.ServiceURL("address-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res AddressGroupResponse
	return &res.AddressGroup, extract.Into(raw.Body, &res)
}
