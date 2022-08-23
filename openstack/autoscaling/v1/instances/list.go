package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToInstancesListQuery() (string, error)
}

type ListOpts struct {
	LifeCycleStatus string `q:"life_cycle_state"`
	HealthStatus    string `q:"health_status"`
}

func (opts ListOpts) ToInstancesListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, groupID string, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("scaling_group_instance", groupID, "list")
	if opts != nil {
		q, err := opts.ToInstancesListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return InstancePage{pagination.SinglePageBase(r)}
	})
}
