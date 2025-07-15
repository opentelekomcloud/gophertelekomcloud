package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetOpts struct {
	// Filter When job_name is all, this parameter is used for fuzzy job filtering.
	Filter string `json:"filter,omitempty"`
	// PageNumber is a page number
	PageNumber int `json:"page_no,omitempty"`
	// PageSize is a number of jobs on each page. The value ranges from 10 to 100.
	PageSize int `json:"page_size,omitempty"`
	// JobType is a type of the jobs to be queried.
	JobType string `json:"jobType,omitempty"`
}

// Get is used to query a job.
// Send request GET /v1.1/{project_id}/clusters/{cluster_id}/cdm/job/{job_name}
func Get(client *golangsdk.ServiceClient, clusterId, jobName string, opts *GetOpts) (*GetQueryResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetQueryResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetQueryResp struct {
	// Total is a number of jobs.
	Total int `json:"total"`
	// Jobs is a Job list. For details, see the descriptions of jobs parameters.
	Jobs []*Job `json:"jobs"`
	// PageNumber is a page number. Jobs on the specified page will be returned.
	PageNumber int `json:"page_no"`
	// PageSize is a number of jobs on each page.
	PageSize int `json:"page_size,"`
}
