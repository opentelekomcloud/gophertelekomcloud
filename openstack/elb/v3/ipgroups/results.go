package ipgroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/structs"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

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

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a IpGroup.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a IpGroup.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

type BatchResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

type IpGroupPage struct {
	pagination.PageWithInfo
}

func (p IpGroupPage) IsEmpty() (bool, error) {
	l, err := ExtractIpGroups(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

func ExtractIpGroups(r pagination.Page) ([]IpGroup, error) {
	var s []IpGroup
	err := (r.(IpGroupPage)).ExtractIntoSlicePtr(&s, "ipgroups")
	if err != nil {
		return nil, err
	}
	return s, nil
}
