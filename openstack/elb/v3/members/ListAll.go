package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllOpts struct {
	Marker              string   `q:"marker,omitempty"`
	Limit               int      `q:"limit,omitempty"`
	PageReverse         bool     `q:"page_reverse,omitempty"`
	Name                []string `q:"name,omitempty"`
	Weight              []int    `q:"weight,omitempty"`
	AdminStateUp        *bool    `q:"admin_state_up,omitempty"`
	SubnetCIDRID        []string `q:"subnet_cidr_id,omitempty"`
	Address             []string `q:"address,omitempty"`
	ProtocolPort        []int    `q:"protocol_port,omitempty"`
	ID                  []string `q:"id,omitempty"`
	OperatingStatus     []string `q:"operating_status,omitempty"`
	EnterpriseProjectID []string `q:"enterprise_project_id,omitempty"`
	IPVersion           []string `q:"ip_version,omitempty"`
	PoolID              []string `q:"pool_id,omitempty"`
	LoadbalancerID      []string `q:"loadbalancer_id,omitempty"`
}

func ListAll(client *golangsdk.ServiceClient, opts ListAllOpts) ([]Member, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("members").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return MemberPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractMembers(pages)
}
