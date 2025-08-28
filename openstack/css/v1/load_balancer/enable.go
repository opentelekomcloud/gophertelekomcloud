package load_balancer

import (
	"encoding/json"
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// EnableLoadBalancerOpts holds options for enabling the load balancer.
type EnableLoadBalancerOpts struct {
	// These parameters are passed to the loadbalancer.EnableLoadBalancer function.
	// ID of the loadbalancer of the css cluster.
	ElbId string `json:"elb_id" required:"true"`
	// These parameters are passed to the loadbalancer.EnableLoadBalancer function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency" required:"true"`
}

// EnableLoadBalancer function is used to enable the loadbalancer switch of a CSS cluster base on EnableLoadBalancerOpts.
func EnableLoadBalancer(client *golangsdk.ServiceClient, clusterID string, opts EnableLoadBalancerOpts) (string, error) {
	body := map[string]interface{}{
		"enable": true,
		"elb_id": opts.ElbId,
		"agency": opts.Agency,
	}

	b, err := build.RequestBody(body, "")
	if err != nil {
		return "", err
	}

	url := client.ServiceURL("clusters", clusterID, "loadbalancers", "es-switch")

	var responseBody map[string]interface{}

	_, err = client.Post(url, b, &responseBody, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return "", err
	}

	// Extract elb_id from response
	elbIDValue, ok := responseBody["elb_id"]
	if !ok {
		return "", fmt.Errorf("elb_id not found in response")
	}

	elbID, ok := elbIDValue.(string)
	if !ok {
		jsonData, _ := json.Marshal(elbIDValue)
		if err := json.Unmarshal(jsonData, &elbID); err != nil {
			return "", fmt.Errorf("elb_id is not a string: %v", err)
		}
	}

	return elbID, nil
}
