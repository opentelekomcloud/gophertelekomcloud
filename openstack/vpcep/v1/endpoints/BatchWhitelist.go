package endpoints

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchUpdateReq struct {
	Permissions []string `json:"permissions" required:"true"`
	Action      string   `json:"action" required:"true"`
}

func BatchUpdateWhitelist(client *golangsdk.ServiceClient, id string, opts BatchUpdateReq) (*BatchUpdateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/vpc-endpoint-services/{vpc_endpoint_service_id}/permissions/action
	raw, err := client.Post(client.ServiceURL("vpc-endpoint-services", id, "permissions", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchUpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchUpdateResponse struct {
	Permissions []string `json:"permissions"`
}
