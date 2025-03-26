package servicegroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddGroupMemberOpts struct {
	// Service group ID.
	SetId string `json:"set_id" required:"true"`
	// Service group member list.
	ServiceItems []ServiceItem `json:"service_items" required:"true"`
}

type ServiceItem struct {
	// Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	// It cannot be left blank when type is set to 0 (manual), and can be left blank when type is set to 1 (automatic).
	Protocol int `json:"protocol" required:"true"`
	// Source port.
	SourcePort string `json:"source_port" required:"true"`
	// Destination port.
	DestPort string `json:"dest_port" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to add a member to a service group.
func AddGroupMember(client *golangsdk.ServiceClient, opts AddGroupMemberOpts) (*ServiceItemsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/service-items
	raw, err := client.Post(client.ServiceURL("service-items"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res AddGroupMemberResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type AddGroupMemberResponse struct {
	Data ServiceItemsResponse `json:"data"`
}

type ServiceItemsResponse struct {
	// List of service group member IDs.
	Items []ServiceItemId `json:"items"`
}

type ServiceItemId struct {
	// ID of an service group member.
	Id string `json:"id"`
}
