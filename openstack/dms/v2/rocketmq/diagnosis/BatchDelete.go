package diagnosis

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteOpts struct {
	// Diagnosis report ID list.
	ReportIDList []string `json:"report_id_list" required:"true"`
}

// BatchDeleteResponse is returned by BatchDelete.
type BatchDeleteResponse struct {
	// Diagnosis report ID list.
	ReportIDList []string `json:"report_id_list"`
}

// BatchDelete deletes diagnosis reports of an instance in batches.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/diagnosis/batch-delete
func BatchDelete(client *golangsdk.ServiceClient, instanceID string, opts BatchDeleteOpts) (*BatchDeleteResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "instances", instanceID, "diagnosis", "batch-delete").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchDeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
