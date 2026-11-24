package resolverrules

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// v21URL builds a /v2.1 URL. The DNS client ResourceBase points to /v2/, so ServiceURL can't be used here.
func v21URL(client *golangsdk.ServiceClient, parts ...string) string {
	return client.Endpoint + "v2.1/" + strings.Join(parts, "/")
}

// Get is used to query an endpoint rule.
// Send GET /v2.1/resolverrules/{resolverrule_id}
func Get(client *golangsdk.ServiceClient, ruleID string) (*ResolverRule, error) {
	if ruleID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "ruleID"}
	}

	raw, err := client.Get(v21URL(client, "resolverrules", ruleID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ResolverRule
	err = extract.IntoStructPtr(raw.Body, &res, "resolver_rule")
	return &res, err
}

// ResolverRule is the endpoint rule returned by the endpoint rule APIs.
// IPAddresses is returned only by Get, Routers only by Get and List.
type ResolverRule struct {
	// Endpoint rule ID, which is a UUID used to identify the rule.
	ID string `json:"id"`
	// Rule name.
	Name string `json:"name"`
	// Domain name.
	DomainName string `json:"domain_name"`
	// ID of the endpoint to which the current rule belongs.
	EndpointID string `json:"endpoint_id"`
	// Resource status. The value can be: ACTIVE, PENDING_CREATE, PENDING_DELETE, ERROR.
	Status string `json:"status"`
	// Rule type. This parameter is reserved. The default value is FORWARD.
	RuleType string `json:"rule_type"`
	// Number of IP addresses in the endpoint rule.
	IPAddressCount int `json:"ipaddress_count"`
	// Destination IP address added to a rule.
	IPAddresses []IPInfo `json:"ipaddresses"`
	// VPC associated with the endpoint rule.
	Routers []Router `json:"routers"`
	// Creation time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	CreateTime string `json:"create_time"`
	// Update time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	UpdateTime string `json:"update_time"`
}

// Router is a VPC associated with the endpoint rule.
type Router struct {
	// ID of the associated VPC.
	RouterID string `json:"router_id"`
	// Region where the associated VPC is located.
	RouterRegion string `json:"router_region"`
	// Resource status.
	Status string `json:"status"`
}
