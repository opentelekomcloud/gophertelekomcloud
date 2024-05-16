package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const historyEndpoint = "submissions"

type HistoryQuery struct {
	JobName string `json:"jname" required:"true"`
}

// GetHistory is used to query the job status.
// Send request GET /v1.1/{project_id}/clusters/{cluster_id}/cdm/submissions
func GetHistory(client *golangsdk.ServiceClient, clusterId string, opts *HistoryQuery) (*StatusResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(clustersEndpoint, clusterId, cdmEndpoint, historyEndpoint).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res StatusResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type HistoryResp struct {
	// Submissions is an array of StartJobSubmission objects.
	Submissions []StatusJobSubmission `json:"submissions"`
	// Total is a number of historical records for a job.
	Total int `json:"total"`
	// PageNumber is a page number.
	PageNumber int `json:"page_no"`
	// PageSize is a number of records on each page. The default value is 10.
	PageSize int `json:"page_size"`
}
