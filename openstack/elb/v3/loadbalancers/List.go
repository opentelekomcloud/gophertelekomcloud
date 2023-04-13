package loadbalancers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToLoadbalancerListQuery() (string, error)
}

type ListOpts struct {
	ID                   []string `q:"id"`
	Name                 []string `q:"name"`
	Description          []string `q:"description"`
	ProvisioningStatus   []string `q:"provisioning_status"`
	OperatingStatus      []string `q:"operating_status"`
	VpcID                []string `q:"vpc_id"`
	VipPortID            []string `q:"vip_port_id"`
	VipAddress           []string `q:"vip_address"`
	VipSubnetCidrID      []string `q:"vip_subnet_cidr_id"`
	L4FlavorID           []string `q:"l4_flavor_id"`
	L4ScaleFlavorID      []string `q:"l4_scale_flavor_id"`
	AvailabilityZoneList []string `q:"availability_zone_list"`
	L7FlavorID           []string `q:"l7_flavor_id"`
	L7ScaleFlavorID      []string `q:"l7_scale_flavor_id"`
	Limit                int      `q:"limit"`
	Marker               string   `q:"marker"`
}

// ToLoadbalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadbalancerListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("loadbalancers")
	if opts != nil {
		query, err := opts.ToLoadbalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return LoadbalancerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
