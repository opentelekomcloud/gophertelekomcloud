package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
)

// Get retrieves a particular Configuration based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*IpGroup, error) {
	// GET /v3/{project_id}/elb/ipgroups/{ipgroup_id}
	raw, err := client.Get(client.ServiceURL("ipgroups", id), nil, nil)
	return extra(err, raw)
}

// IpGroup The IP address can contain IP addresses or CIDR blocks.
// 0.0.0.0 will be considered the same as 0.0.0.0/32. If you enter both 0.0.0.0 and 0.0.0.0/32,
// only one will be kept. 0:0:0:0:0:0:0:1 will be considered the same as ::1 and ::1/128.
// If you enter 0:0:0:0:0:0:0:1, ::1 and ::1/128, only one will be kept.
type IpGroup struct {
	// The unique ID for the IpGroup.
	ID string `json:"id"`
	// Specifies the IP address group name.
	Name string `json:"name"`
	// Provides remarks about the IP address group.
	Description string `json:"description"`
	// Specifies the project ID of the IP address group.
	ProjectId string `json:"project_id"`
	// Specifies the IP addresses or CIDR blocks in the IP address group. [] indicates any IP address.
	IpList []IpInfo `json:"ip_list"`
	// Lists the IDs of listeners with which the IP address group is associated.
	Listeners []structs.ResourceRef `json:"listeners"`
	// Specifies the time when the IP address group was created.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the IP address group was updated.
	UpdatedAt string `json:"updated_at"`
}

type IpInfo struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip"`
	// Provides remarks about the IP address group.
	Description string `json:"description"`
}
