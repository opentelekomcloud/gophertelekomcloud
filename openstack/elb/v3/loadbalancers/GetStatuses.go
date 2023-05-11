package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetStatuses will return the status of a particular LoadBalancer.
func GetStatuses(client *golangsdk.ServiceClient, id string) (*LoadBalancer, error) {
	raw, err := client.Get(client.ServiceURL("loadbalancers", id, "statuses"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Loadbalancer LoadBalancer `json:"loadbalancer"`
	}
	err = extract.IntoStructPtr(raw.Body, &res, "statuses")
	return &res.Loadbalancer, nil
}
