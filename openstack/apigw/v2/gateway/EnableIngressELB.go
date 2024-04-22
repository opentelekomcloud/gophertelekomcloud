package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// ElbIngressAccessOpts is the structure that used to bind ingress EIP to instance when loadbalancer_provider is set to elb.
type ElbIngressAccessOpts struct {
	// The APIG dedicated instance ID.
	InstanceId string `json:"-"`
	// Public inbound access bandwidth.
	IngressBandwithSize int `json:"bandwidth_size" required:"true"`
	// Billing type of the public inbound access bandwidth.
	// + bandwidth: billed by bandwidth.
	// + traffic: billed by traffic.
	IngressBandwithChargingMode string `json:"bandwidth_charging_mode" required:"true"`
}

// EnableElbIngressAccess is a method to bind the ingress eip associated with an existing APIG dedicated instance.
// Supported only when loadbalancer_provider is set to elb.
func EnableElbIngressAccess(client *golangsdk.ServiceClient, opts ElbIngressAccessOpts) (*GatewayResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	var r GatewayResp
	_, err = client.Post(client.ServiceURL("apigw", "instances", opts.InstanceId, "ingress-eip"), b, &r, &golangsdk.RequestOpts{})
	return &r, err
}
