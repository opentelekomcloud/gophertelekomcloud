package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDDosStatusOpts struct {
	// ID of an EIP
	FloatingIpId string
	// If this parameter is not used,
	// the defense statuses of all ECSs are displayed in the Neutron-queried order by default.
	Status string `q:"status"`
	// Limit of number of returned results
	Limit int `q:"limit"`
	// Offset
	Offset int `q:"offset"`
	// IP address. Both IPv4 and IPv6 addresses are supported.
	// For example, if you enter ?ip=192.168,
	// the defense status of EIPs corresponding to 192.168.111.1 and 10.192.168.8 is returned.
	Ip string `q:"ip"`
}

type DDosStatus struct {
	// Floating IP address
	FloatingIpAddress string `json:"floating_ip_address"`
	// ID of an EIP
	FloatingIpId string `json:"floating_ip_id"`
	// EIP type.
	NetworkType string `json:"network_type"`
	// Defense status
	Status string `json:"status"`
}

func ListDDosStatus(client *golangsdk.ServiceClient, opts ListDDosStatusOpts) (*ListStatusResponse, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/antiddos
	raw, err := client.Get(client.ServiceURL("antiddos")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListStatusResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListStatusResponse struct {
	// Total number of EIPs
	Total int `json:"total"`
	// List of defense statuses
	DDosStatus []DDosStatus `json:"ddosStatus"`
}
