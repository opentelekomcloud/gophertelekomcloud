package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePeerOpts struct {
	// Specifies the name of the virtual interface peer.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the virtual interface peer.
	Description string `json:"description,omitempty"`
	// Specifies the remote subnet list, which records the CIDR blocks used in the on-premises data center.
	RemoteEpGroup []string `json:"remote_ep_group,omitempty"`
	// Specifies the ID of the virtual interface corresponding to the virtual interface peer.
}

func UpdatePeer(c *golangsdk.ServiceClient, id string, opts UpdatePeerOpts) (*VifPeer, error) {
	b, err := build.RequestBody(opts, "vif_peer")
	if err != nil {
		return nil, err
	}

	raw, err := c.Put(c.ServiceURL("dcaas", "vif-peers", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}
	var res VifPeer
	err = extract.IntoStructPtr(raw.Body, &res, "vif_peer")
	return &res, err
}
