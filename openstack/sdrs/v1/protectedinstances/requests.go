package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToInstanceListQuery() (string, error)
}

type ListOpts struct {
	ServerGroupID        string   `q:"server_group_id"`
	ServerGroupIDs       []string `q:"server_group_ids"`
	ProtectedInstanceIDs []string `q:"protected_instance_ids"`
	Limit                int      `q:"limit"`
	Offset               int      `q:"offset"`
	Status               string   `q:"status"`
	Name                 string   `q:"name"`
	QueryType            string   `q:"query_type"`
	AvailabilityZone     string   `q:"availability_zone"`
}

func (opts ListOpts) ToInstanceListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		q, err := opts.ToInstanceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return InstancePage{pagination.SinglePageBase(r)}
	})
}
