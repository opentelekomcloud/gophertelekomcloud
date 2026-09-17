package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*LoadBalancer, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("loadbalancers", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		LoadBalancer LoadBalancer `json:"loadbalancer"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.LoadBalancer, nil
}
