package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Whether to only send the check request.
	// The value can be:
	// true: A check request will be sent and the IP address group will not be updated.
	// 	Check items include mandatory parameters, request format, and constraints.
	// 	If the check fails, an error will be returned.
	// 	If the check succeeds, response code 202 will be returned.
	// false: A request will be sent and an IP address group will be created.
	DryRun bool `json:"dry_run,omitempty"`
	// Request body for creating an IP address group.
	AddressGroup UpdateAddressGroupOptions `json:"address_group" required:"true"`
}

type UpdateAddressGroupOptions struct {
	// IP Address group name.
	// The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Description about the IP Address group.
	// The value can contain up to 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
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

// This function is used to update an IP address group.
func Update(client *golangsdk.ServiceClient, groupId string, opts UpdateOpts) (*AddressGroup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/vpc/address-groups/{address_group_id}
	raw, err := client.Put(client.ServiceURL("address-groups", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res AddressGroupResponse
	return &res.AddressGroup, extract.Into(raw.Body, &res)
}
