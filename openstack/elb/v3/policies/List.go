package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

type ListOpts struct {
	Marker             string   `q:"marker"`
	Limit              int      `q:"limit"`
	PageReverse        bool     `q:"page_reverse"`
	ID                 []string `q:"id"`
	Name               []string `q:"name"`
	Description        []string `q:"description"`
	ListenerID         []string `q:"listener_id"`
	Action             []string `q:"action"`
	RedirectPoolID     []string `q:"redirect_pool_id"`
	RedirectListenerID []string `q:"redirect_listener_id"`
	ProvisioningStatus []string `q:"provisioning_status"`
	DisplayAllRules    bool     `q:"display_all_rules"`
}

func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("l7policies")
	if opts != nil {
		q, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}
