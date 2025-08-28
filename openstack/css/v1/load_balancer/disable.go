package load_balancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// DisableLoadBalancer disables the loadbalancer switch of a CSS cluster.
func DisableLoadBalancer(client *golangsdk.ServiceClient, clusterID string) error {
	body := map[string]interface{}{
		"enable": false,
	}

	b, err := build.RequestBody(body, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "loadbalancers", "es-switch")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
