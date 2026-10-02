package specification

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResizeOpts struct {
	// Change type.
	//    storage: expand the storage without changing the broker quantity.
	//    horizontal: scale up the RocketMQ 5.x instance specification.
	OperType string `json:"oper_type" required:"true"`
	// New storage space, in GB. Mandatory when oper_type is storage:
	// the storage space of each broker must be expanded by at least 100 GB.
	NewStorageSpace int `json:"new_storage_space,omitempty"`
	// New product ID. Mandatory when oper_type is horizontal.
	// Specification changes are only for cluster instances.
	NewProductID string `json:"new_product_id,omitempty"`
	// IDs of the EIPs bound to the instance, separated by commas (,).
	// Mandatory when oper_type is horizontal and public access is enabled.
	PublicIPID string `json:"publicip_id,omitempty"`
}

// ResizeResponse is returned by Resize.
type ResizeResponse struct {
	// ID of the specification modification task.
	JobID string `json:"job_id"`
}

// Resize changes the specifications of a RocketMQ instance.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/extend
func Resize(client *golangsdk.ServiceClient, instanceID string, opts ResizeOpts) (*ResizeResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "instances", instanceID, "extend").Build()
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

	var res ResizeResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
