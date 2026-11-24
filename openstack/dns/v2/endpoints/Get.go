package endpoints

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// v21URL builds a /v2.1 URL. The DNS client ResourceBase points to /v2/, so ServiceURL can't be used here.
func v21URL(client *golangsdk.ServiceClient, parts ...string) string {
	return client.Endpoint + "v2.1/" + strings.Join(parts, "/")
}

// Get is used to query an endpoint.
// Send GET /v2.1/endpoints/{endpoint_id}
func Get(client *golangsdk.ServiceClient, endpointID string) (*Endpoint, error) {
	if endpointID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "endpointID"}
	}

	raw, err := client.Get(v21URL(client, "endpoints", endpointID), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Endpoint
	err = extract.IntoStructPtr(raw.Body, &res, "endpoint")
	return &res, err
}

// Endpoint is the EndpointResp object returned by the endpoint APIs.
type Endpoint struct {
	// Endpoint ID, which is a UUID used to identify the endpoint.
	ID string `json:"id"`
	// Endpoint name.
	Name string `json:"name"`
	// Endpoint direction. The value can be:
	//   inbound: indicates an inbound endpoint.
	//   outbound: indicates an outbound endpoint.
	Direction string `json:"direction"`
	// Resource status. The value can be: ACTIVE, PENDING_CREATE, PENDING_DELETE, ERROR.
	Status string `json:"status"`
	// ID of the VPC to which the endpoint belongs.
	VpcID string `json:"vpc_id"`
	// Number of IP addresses of the endpoint.
	IPAddressCount int `json:"ipaddress_count"`
	// Number of endpoint rules in the endpoint.
	// This parameter is returned only for outbound endpoints.
	ResolverRuleCount int `json:"resolver_rule_count"`
	// Creation time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	CreateTime string `json:"create_time"`
	// Update time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS
	UpdateTime string `json:"update_time"`
}
