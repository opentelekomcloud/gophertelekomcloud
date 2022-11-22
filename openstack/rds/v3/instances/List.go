package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	//
	Id string `q:"id"`
	//
	Name string `q:"name"`
	//
	Type string `q:"type"`
	//
	DataStoreType string `q:"datastore_type"`
	//
	VpcId string `q:"vpc_id"`
	//
	SubnetId string `q:"subnet_id"`
	//
	Offset int `q:"offset"`
	//
	Limit int `q:"limit"`
}

type ListRdsBuilder interface {
	ToRdsListDetailQuery() (string, error)
}

func (opts ListOpts) ToRdsListDetailQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

type ListRdsInstanceOpts struct {
	//
	Id string `q:"id"`
	//
	Name string `q:"name"`
	//
	Type string `q:"type"`
	//
	DataStoreType string `q:"datastore_type"`
	//
	VpcId string `q:"vpc_id"`
	//
	SubnetId string `q:"subnet_id"`
	//
	Offset int `q:"offset"`
	//
	Limit int `q:"limit"`
}

func (opts ListRdsInstanceOpts) ToRdsListDetailQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListRdsBuilder) pagination.Pager {
	url := client.ServiceURL("instances")
	if opts != nil {
		query, err := opts.ToRdsListDetailQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RdsPage{pagination.SinglePageBase(r)}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}
