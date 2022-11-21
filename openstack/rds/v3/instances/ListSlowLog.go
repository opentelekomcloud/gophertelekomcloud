package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type DbSlowLogOpts struct {
	StartDate string `q:"start_date"`
	EndDate   string `q:"end_date"`
	Offset    string `q:"offset"`
	Limit     string `q:"limit"`
	Level     string `q:"level"`
}

type DbSlowLogBuilder interface {
	ToDbSlowLogListQuery() (string, error)
}

func (opts DbSlowLogOpts) ToDbSlowLogListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func ListSlowLog(client *golangsdk.ServiceClient, opts DbSlowLogBuilder, instanceID string) pagination.Pager {
	url := client.ServiceURL("instances", instanceID, "slowlog")
	if opts != nil {
		query, err := opts.ToDbSlowLogListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SlowLogPage{pagination.SinglePageBase(r)}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}
