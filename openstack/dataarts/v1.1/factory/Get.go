package factory

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type GetQuery struct {
	Filter    string `q:"filter"`
	PageNo    int    `q:"page_no"`
	PageSize  int    `q:"page_size"`
	JobType   string `q:"jobType"`
	ClusterId string
	JobName   string
}

// Get returns 500 error, not usable
func Get(client *golangsdk.ServiceClient, opts GetQuery) ([]Job, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	url := "clusters/" + opts.ClusterId + "/cdm/job/" + opts.JobName
	pages, err := pagination.NewPager(client, client.ServiceURL(url)+q.String(),
		func(r pagination.PageResult) pagination.Page {
			return JobPage{pagination.LinkedPageBase{PageResult: r}}
		}).AllPages()
	if err != nil {
		return nil, err
	}

	allJobs, err := ExtractJobs(pages)
	if err != nil {
		return nil, err
	}

	return allJobs, err
}

func ExtractJobs(r pagination.Page) ([]Job, error) {
	var res []Job
	err := extract.IntoSlicePtr(r.(JobPage).Result.BodyReader(), &res, "jobs")
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Job struct {
}

type JobPage struct {
	pagination.LinkedPageBase
}
