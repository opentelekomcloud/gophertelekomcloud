package management

import (
	"errors"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetQueryParameters represents the query parameters for the firewall instance list.
type GetQueryParameters struct {
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset int `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
	// Service type. Currently, only 0 (Internet protection) is supported.
	ServiceType int `q:"service_type" required:"true"`
	// Enterprise project ID
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
	// Firewall instance ID. This field is required if Name is not provided.
	FwInstanceID string `q:"fw_instance_id,omitempty"`
	// Firewall name. This field is required if FwInstanceID is not provided.
	Name string `q:"name,omitempty"`
}

// This function is used to query details about a Firewall instance.
func Get(client *golangsdk.ServiceClient, opts GetQueryParameters) (*GetFirewallInstanceResponseRecord, error) {
	if opts.FwInstanceID == "" && opts.Name == "" {
		return nil, errors.New("one of the two i.e name or firewall instance id, is required in opts")
	}
	// GET /v1/{project_id}/firewall/exist
	url, err := golangsdk.NewURLBuilder().WithEndpoints("firewall", "exist").WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}
	firewallInstance := res.Data.Records[0]
	return &firewallInstance, nil
}

type GetResponse struct {
	// Contains the data for the firewall instance response.
	Data GetFirewallInstanceData `json:"data"`
}

// GetFirewallInstanceData represents the data for the firewall instance response.
type GetFirewallInstanceData struct {
	// The maximum number of records to return.
	Limit int `json:"limit"`
	// The offset of the first record to return.
	Offset int `json:"offset"`
	// The total number of firewall instances.
	Total int `json:"total"`
	// The list of firewall instance records.
	Records []GetFirewallInstanceResponseRecord `json:"records"`
}

// GetFirewallInstanceResponseRecord represents the structure of an individual firewall instance record.
type GetFirewallInstanceResponseRecord struct {
	// The unique identifier of the firewall instance
	FwInstanceID string `json:"fw_instance_id"`
	// The name of the firewall instance
	Name string `json:"name"`
	// Cluster type: 0 (active/standby), 1 (cluster). In active/standby mode, there are four nodes.
	// Two active nodes form a cluster, and the other two are the standby of the active nodes.
	// In cluster mode, only two nodes are started to form a cluster.
	HAType int `json:"ha_type"`
	// Billing mode: 0 (yearly/monthly), 1 (pay-per-use).
	ChargeMode string `json:"charge_mode"`
	// Firewall protection type. Currently, its value can only be 0 (Internet protection).
	ServiceType int `json:"service_type"`
	// Engine type. Its value can only be 1 (Hillstone engine).
	EngineType int `json:"engine_type"`
	// Firewall specifications.
	Flavor Flavor `json:"flavor"`
	// Protected object list.
	ProtectObjects []ProtectObjectVO `json:"protect_objects"`
	// Firewall status: -1 (waiting for payment), 0 (creating),
	// 1 (deleting), 2 (running), 3 (upgrading), 4 (deleted),
	// 5 (frozen), 6 (creation failed), 7 (deletion failed),
	// 8 (freezing failed), or 9 (being stored), 10 (storage failed),
	// or 11 (upgrade failed).
	Status int `json:"status"`
	// Whether an engine old: true (yes), false (no).
	IsOldFirewallInstance bool `json:"is_old_firewall_instance"`
	// Whether OBS is supported: true (yes), false (no).
	IsAvailableObs bool `json:"is_available_obs"`
	// Whether threat intelligence tags are supported: true (yes), false (no).
	IsSupportThreatTags bool `json:"is_support_threat_tags"`
	// Whether IPv6 is supported: true (yes), false (no).
	SupportIpv6 bool `json:"support_ipv6"`
	// Whether a feature is enabled: true (yes), false (no).
	FeatureToggle map[string]bool `json:"feature_toggle"`
	// Firewall resource list.
	Resources []FirewallInstanceResource `json:"resources"`
	// The enterprise project ID of the firewall instance
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// The resource ID of the firewall instance
	ResourceID string `json:"resource_id"`
	// Whether website filtering is supported: true (yes), false (no).
	SupportUrlFiltering bool `json:"support_url_filtering"`
	// The list of tags associated with the firewall instance
	Tags string `json:"tags"`
}

type ProtectObjectVO struct {
	// Protected object ID. It is used to distinguish Internet border protection from VPC border protection after a CFW instance is created.
	ObjectID string `json:"object_id"`
	// Protected object name.
	ObjectName string `json:"object_name"`
	// Project type: 0 (north-south), 1 (east-west).
	Type int `json:"type"`
}

// Resource represents a cloud resource with its ID, type, and specifications.
type FirewallInstanceResource struct {
	// Resource ID. It can be the firewall ID, bandwidth ID, EIP ID, VPC ID, or the ID returned after CBC callback.
	ResourceID string `json:"resource_id"`
	// Service type, which is used by CBC. The value is otc.service.type.cfw.
	CloudServiceType string `json:"cloud_service_type"`
	// Resource type. Enumeration values:
	// - otc.resource.type.cfw (cloud firewall)
	// - otc.resource.type.cfw.exp.eip (EIP)
	// - otc.resource.type.cfw.exp.bandwidth (bandwidth)
	// - otc.resource.type.cfw.exp (VPC)
	ResourceType string `json:"resource_type"`
	// Inventory unit code:
	// - cfw.standard (firewall standard edition)
	// - cfw.professional (firewall professional edition)
	// - cfw.expack.eip.standard (EIP standard edition)
	// - cfw.expack.eip.professional (EIP professional edition)
	// - cfw.expack.bandwidth.standard (bandwidth basic edition)
	// - cfw.expack.bandwidth.professional (bandwidth professional edition)
	// - cfw.expack.vpc.professional (VPC professional edition)
	ResourceSpecCode string `json:"resource_spec_code"`
	// Resource quantity.
	ResourceSize int `json:"resource_size"`
	// Resource unit.
	ResourceSizeMeasureID int `json:"resource_size_measure_id"`
}
