package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// EnableLoadBalancerOpts holds options for enabling the load balancer.
type EnableLoadBalancerOpts struct {
	ElbId  string `json:"elb_id" required:"true"`
	Agency string `json:"agency" required:"true"`
}

// EnableLoadBalancer enables the load balancer switch of a CSS cluster.
func EnableLoadBalancer(client *golangsdk.ServiceClient, clusterID string, opts EnableLoadBalancerOpts) (*LoadbalancerSwitchResp, error) {
	return setLoadBalancerSwitch(client, clusterID, true, opts.ElbId, opts.Agency)
}

// DisableLoadBalancer disables the load balancer switch of a CSS cluster.
func DisableLoadBalancer(client *golangsdk.ServiceClient, clusterID string) (*LoadbalancerSwitchResp, error) {
	return setLoadBalancerSwitch(client, clusterID, false, "", "")
}

// internal function to handle both enable and disable
func setLoadBalancerSwitch(client *golangsdk.ServiceClient, clusterID string, enable bool, elbId, agency string) (*LoadbalancerSwitchResp, error) {
	body := map[string]interface{}{
		"enable": enable,
	}

	if enable {
		body["elb_id"] = elbId
		body["agency"] = agency
	}

	b, err := build.RequestBody(body, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("clusters", clusterID, "loadbalancers", "es-switch")

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return nil, err
	}

	var res LoadbalancerSwitchResp
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return &res, err
}

type LoadbalancerSwitchResp struct {
	ElbId string `json:"elb_id"`
}
