package job

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
	url, err := golangsdk.NewURLBuilder().WithEndpoints("clusters", opts.ClusterId, "cdm", "job", opts.JobName).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.NewPager(client, client.ServiceURL(url.String()),
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
