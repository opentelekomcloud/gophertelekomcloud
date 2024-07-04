package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ProtectionOpts struct {
	// HSS edition. Its value can be:
	// hss.version.null: protection disabled
	// hss.version.enterprise: enterprise edition
	// hss.version.premium: premium edition
	Version string `json:"version,omitempty"`
	// on_demand: pay-per-use
	ChargingMode string `json:"charging_mode,omitempty"`
	// Instance ID
	ResourceId string `json:"resource_id,omitempty"`
	// Server ID list
	HostIds []string `json:"host_id_list,omitempty"`
	// Resource tag
	Tags []tags.ResourceTag `json:"tags"`
}

func ChangeProtectionStatus(client *golangsdk.ServiceClient, opts ProtectionOpts) (*ProtectionResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v5/{project_id}/host-management/protection
	raw, err := client.Post(client.ServiceURL("host-management", "protection"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"region": client.RegionID},
	})
	if err != nil {
		return nil, err
	}

	var res ProtectionResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ProtectionResp struct {
	// HSS edition.
	Version string `json:"version"`
	// on_demand: pay-per-use
	ChargingMode string `json:"charging_mode"`
	// Instance ID
	ResourceId string `json:"resource_id"`
	// Server ID list
	HostIds []string `json:"host_id_list"`
	// Resource tag
	Tags []tags.ResourceTag `json:"tags"`
}
