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
	url := client.ServiceURL("flavors", dbName)
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

type Flavor struct {
	VCPUs        string            `json:"vcpus"`
	RAM          int               `json:"ram"`
	SpecCode     string            `json:"spec_code"`
	InstanceMode string            `json:"instance_mode"`
	AzStatus     map[string]string `json:"az_status"`
}

type DbFlavorsPage struct {
	pagination.SinglePageBase
}

func (r DbFlavorsPage) IsEmpty() (bool, error) {
	flavors, err := ExtractDbFlavors(r)
	if err != nil {
		return false, err
	}
	return len(flavors) == 0, err
}

func ExtractDbFlavors(r pagination.Page) ([]Flavor, error) {
	var s []Flavor
	err := (r.(DbFlavorsPage)).ExtractIntoSlicePtr(&s, "flavors")
	if err != nil {
		return nil, err
	}
	return s, nil
}
