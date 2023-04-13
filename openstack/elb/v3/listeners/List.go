package listeners

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

type ListOpts struct {
	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`

	ProtocolPort            []int      `q:"protocol_port"`
	Protocol                []Protocol `q:"protocol"`
	Description             []string   `q:"description"`
	DefaultTLSContainerRef  []string   `q:"default_tls_container_ref"`
	ClientCATLSContainerRef []string   `q:"client_ca_tls_container_ref"`
	DefaultPoolID           []string   `q:"default_pool_id"`
	ID                      []string   `q:"id"`
	Name                    []string   `q:"name"`
	LoadBalancerID          []string   `q:"loadbalancer_id"`
	TLSCiphersPolicy        []string   `q:"tls_ciphers_policy"`
	MemberAddress           []string   `q:"member_address"`
	MemberDeviceID          []string   `q:"member_device_id"`
	MemberTimeout           []int      `q:"member_timeout"`
	ClientTimeout           []int      `q:"client_timeout"`
	KeepAliveTimeout        []int      `q:"keepalive_timeout"`
}

func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("listeners")
	if opts != nil {
		q, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
