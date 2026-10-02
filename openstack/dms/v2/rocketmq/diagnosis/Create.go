package diagnosis

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Consumer group name.
	GroupName string `json:"group_name" required:"true"`
	// Node ID list.
	NodeIDList []string `json:"node_id_list,omitempty"`
}

// CreateResponse is returned by Create.
type CreateResponse struct {
	// Report ID.
	ReportID string `json:"report_id"`
}

// Create an instance diagnosis task.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/diagnosis
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "instances", instanceID, "diagnosis").Build()
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

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
