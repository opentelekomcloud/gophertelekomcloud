package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	VersionName string `q:"version_name"`
	SpecCode    string `q:"spec_code"`
}

type ListOptsBuilder interface {
	ToListOptsQuery() (string, error)
}

func (opts ListOpts) ToListOptsQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder, dbName string) pagination.Pager {
	url := listURL(client, dbName)
	if opts != nil {
		query, err := opts.ToListOptsQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pagerRDS := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DbFlavorsPage{pagination.SinglePageBase(r)}
	})

	pagerRDS.Headers = map[string]string{"Content-Type": "application/json"}
	return pagerRDS
}
