package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// IngressAccessOpts allows binding and updating the eip associated with an existing APIG dedicated instance with the
// given parameters.
type IngressAccessOpts struct {
	// EIP ID
	EipId string `json:"eip_id,omitempty"`
}

// EnableIngressAccess is a method to bind and update the eip associated with an existing APIG dedicated instance.
func EnableIngressAccess(client *golangsdk.ServiceClient, instanceId string, opts IngressAccessOpts) (*GatewayResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	var r GatewayResp
	_, err = client.Post(client.ServiceURL("apigw", "instances", instanceId, "eip"), b, &r, &golangsdk.RequestOpts{})
	return &r, err
}
