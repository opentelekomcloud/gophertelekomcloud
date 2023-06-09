package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	// GET /v1/{project_id}/premium-waf/instance
	raw, err := client.Get(client.ServiceURL("instance", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Instance
	return &res, extract.Into(raw.Body, &res)
}

type Instance struct {
	// ID of the dedicated WAF engine.
	ID string `json:"id"`
	// Name of the dedicated WAF engine.
	Name string `json:"instance_name"`
	// Region where a dedicated engine is to be created. Its value is EU-DE.
	Region string `json:"region"`
	// Az ID.
	AvailabilityZone string `json:"zone"`
	// CPU architecture.
	Architecture string `json:"arch"`
	// ECS specification ID.
	Flavor string `json:"cpu_flavor"`
	// ID of the VPC where the dedicated engine is located.
	VpcID string `json:"vpc_id"`
	// Subnet ID of the VPC where the dedicated engine is located.
	SubnetId string `json:"subnet_id"`
	// Service plane IP address of the dedicated engine.
	ServiceIp string `json:"service_ip"`
	// IPv6 address of the service plane of the dedicated engine.
	ServiceIpv6 string `json:"service_ipv6"`
	// Security groups bound to the dedicated engine ECS.
	SecurityGroups []string `json:"security_group_ids"`
	// Billing status of dedicated WAF engine. The value can be 0, 1, or 2.
	// 0: The billing is normal.
	// 1: The billing account is frozen. Resources
	// and data will be retained, but the cloud
	// services cannot be used by the account.
	// 2: The billing is terminated. Resources and data will be cleared.
	BillingStatus int `json:"status"`
	// Running status of the dedicated engine. The value can be:
	// 0 (creating)
	// 1 (running)
	// 2 (deleting)
	// 3 (deleted)
	// 4 (creation failed)
	// 5 (frozen)
	// 6 (abnormal)
	// 7 (updating)
	// 8 (update failed).
	Status int `json:"run_status"`
	// Access status of the dedicated engine. The value can be 0 or 1.
	// 0: the dedicated engine is not connected.
	// 1: the dedicated engine is connected.
	AccessStatus int `json:"access_status"`
	// Whether the dedicated engine can be upgraded.
	// The value can be 0 for no or 1 for yes.
	Upgradable int `json:"upgradable"`
	// Dedicated engine ECS specification for example,
	// 8 vCPUs | 16 GB.
	Specification string `json:"specification"`
	// Domain name protected by the dedicated engine.
	Hosts []HostEntry `json:"hosts"`
	// ID of the ECS hosting the dedicated engine.
	ServerId string `json:"server_id"`
	// Timestamp when the dedicated WAF engine was created.
	CreatedAt int `json:"create_time"`
}

type HostEntry struct {
	// ID of the protected domain name. This is a
	// unique ID automatically generated by the system.
	ID string `json:"id"`
	// Protected domain name
	Hostname string `json:"hostname"`
}
