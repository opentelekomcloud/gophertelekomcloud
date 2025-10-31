package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListJobsOpts struct {
	// Task ID
	Id string `q:"id"`
	// Query start time in "yyyy-mm-ddThh:mm:ssZ" format
	StartTime string `q:"start_time"`
	// Query end time in "yyyy-mm-ddThh:mm:ssZ" format
	EndTime string `q:"end_time"`
	// Task status: Running, Completed, Failed
	Status string `q:"status"`
	// Task name
	Name string `q:"name"`
	// Index offset
	Offset int `q:"offset"`
	// Number of records (10, 20, or 50)
	Limit int `q:"limit"`
}

func ListJobs(client *golangsdk.ServiceClient, opts ListJobsOpts) (*ListJobsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("jobs").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListJobsResponse
	return &res, extract.Into(raw.Body, &res)
}

type ListJobsResponse struct {
	Jobs       []JobDetail `json:"jobs"`
	TotalCount int         `json:"total_count"`
}

type JobDetail struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Status     string          `json:"status"`
	StartTime  string          `json:"start_time"`
	EndTime    string          `json:"end_time"`
	Progress   string          `json:"progress"`
	Instance   JobInstanceInfo `json:"instance"`
	FailReason string          `json:"fail_reason"`
}

type JobInstanceInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
