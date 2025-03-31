package host_groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListHostOpts struct {
	// List of host group IDs. The host type must be the same as the host group type.
	HostIdList []string `json:"host_id_list,omitempty"`
	// Filters other than host IDs.
	Filter *HostFilter `json:"filter,omitempty"`
}

type HostFilter struct {
	// Host name list.
	HostNameList []string `json:"host_name_list,omitempty"`
	// List of host IP addresses. You can filter hosts by host IP address.
	HostIpList []string `json:"host_ip_list,omitempty"`
	// Host status. You can filter hosts by host status.
	// uninstall: not installed.
	// running: running.
	// offline: offline.
	// error: abnormal.
	// plugin error: plug-in error.
	// installing: installing.
	// install-fail: Installation failed.
	// upgrading: upgrading.
	// upgrade failed: Upgrade failed.
	// uninstalling: being uninstalled.
	// authentication error: Authentication failed.
	Status string `json:"host_status,omitempty"`
	// Host version. You can filter hosts by host version.
	HostVersion string `json:"host_version,omitempty"`
}

func ListHost(client *golangsdk.ServiceClient, opts ListHostOpts) (*ListHostResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /v3/{project_id}/lts/host-list
	raw, err := client.Post(client.ServiceURL("lts", "host-list"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListHostResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListHostResult struct {
	// Host group details.
	Result []HostResponse `json:"result"`
	// Number of deleted host groups.
	Total int64 `json:"total"`
}

type HostResponse struct {
	// Host ID.
	ID string `json:"host_id"`
	// Host IP.
	IP string `json:"host_ip"`
	// Host name.
	HostName string `json:"host_name"`
	// Host status.
	HostStatus string `json:"host_status"`
	// Host type.
	Type string `json:"host_type"`
	// Host version.
	HostVersion string `json:"host_version"`
	// Update time.
	UpdatedAt int64 `json:"update_time"`
}
