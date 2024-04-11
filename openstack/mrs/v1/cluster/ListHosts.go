package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type ListHostsOpts struct {
	// Cluster ID
	ClusterId string
	// Maximum number of clusters displayed on a page
	// Value range: [1-2147483646].
	// The default value is 10.
	PageSize string `q:"pageSize,omitempty"`
	// Current page number
	// The default value is 1.
	CurrentPage string `q:"currentPage,omitempty"`
}

func ListHosts(client *golangsdk.ServiceClient, opts ListHostsOpts) (*ListHostsResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("clusters", opts.ClusterId, "hosts").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v1.1/{project_id}/clusters/{cluster_id}/hosts
	raw, err := client.Get(client.ServiceURL(url.String()), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res ListHostsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListHostsResponse struct {
	// Total number of hosts in a list
	Total int `json:"total,omitempty"`
	// Host parameters
	Hosts []Hosts `json:"hosts,omitempty"`
}

type Hosts struct {
	// VM ID
	Id string `json:"id,omitempty"`
	// VM IP address
	Ip string `json:"ip,omitempty"`
	// VM flavor ID
	Flavor string `json:"flavor,omitempty"`
	// VM type
	// Currently, MasterNode, CoreNode, and TaskNode are supported.
	Type string `json:"type,omitempty"`
	// VM name
	Name string `json:"name,omitempty"`
	// Current VM state
	Status string `json:"status,omitempty"`
	// Memory
	Mem string `json:"mem,omitempty"`
	// Number of CPU cores
	Cpu string `json:"cpu,omitempty"`
	// OS disk capacity
	RootVolumeSize string `json:"root_volume_size,omitempty"`
	// Data disk type
	DataVolumeType string `json:"data_volume_type,omitempty"`
	// Data disk capacity
	DataVolumeSize int `json:"data_volume_size,omitempty"`
	// Number of data disks
	DataVolumeCount int `json:"data_volume_count,omitempty"`
}
