package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type DbErrorlogOpts struct {
	//
	StartDate string `q:"start_date"`
	//
	EndDate string `q:"end_date"`
	//
	Offset string `q:"offset"`
	//
	Limit string `q:"limit"`
	//
	Level string `q:"level"`
}

type DbErrorlogBuilder interface {
	DbErrorlogQuery() (string, error)
}

func (opts DbErrorlogOpts) DbErrorlogQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func ListErrorLog(client *golangsdk.ServiceClient, opts DbErrorlogBuilder, instanceID string) pagination.Pager {
	url := client.ServiceURL("instances", instanceID, "errorlog")
	if opts != nil {
		query, err := opts.DbErrorlogQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ErrorLogPage{pagination.SinglePageBase(r)}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}
