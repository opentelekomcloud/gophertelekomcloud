package loadbalancers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetStatuses(client *golangsdk.ServiceClient, id string) (*StatusTree, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("loadbalancers", id, "statuses").Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}
	var res struct {
		Statuses StatusTree `json:"statuses"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Statuses, nil
}
