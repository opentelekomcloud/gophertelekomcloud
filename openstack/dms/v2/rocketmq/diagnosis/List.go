package diagnosis

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Diagnosis report list.
	Reports []ReportSummary `json:"diagnosis_report_list"`
	// Number of reports.
	TotalNum int `json:"total_num"`
}

// ReportSummary is an entry of the instance diagnosis report list.
type ReportSummary struct {
	// Report ID.
	ReportID string `json:"report_id"`
	// Consumer group name.
	GroupName string `json:"group_name"`
	// Number of consumers.
	ConsumerNums int `json:"consumer_nums"`
	// Report status.
	//    diagnosing
	//    failed
	//    deleted: deleted manually.
	//    finished
	//    normal: no problems found.
	//    abnormal: problems found.
	Status string `json:"status"`
	// Creation time, for example 2025-05-28 11:25:58.477.
	CreatedAt string `json:"created_at"`
	// Number of abnormal items.
	AbnormalItemSum int `json:"abnormal_item_sum"`
	// Number of abnormal nodes.
	FaultedNodeSum int `json:"faulted_node_sum"`
}

// List the diagnosis reports of a RocketMQ instance.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/diagnosis
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "diagnosis").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
