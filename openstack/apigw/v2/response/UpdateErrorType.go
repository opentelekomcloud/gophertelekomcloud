package response

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateErrorOpts struct {
	// Gateway ID
	GatewayID string `json:"-"`
	// Group ID.
	GroupId string `json:"-"`
	// Response ID.
	ID string `json:"-"`
	// HTTP status code of the response.
	Status int `json:"status,omitempty"`
	// Response body template.
	Body string `json:"body,omitempty"`
}

func UpdateErrorType(client *golangsdk.ServiceClient, respType string, opts UpdateErrorOpts) (*ResponseInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupId, "gateway-responses", opts.ID, respType),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
		})
	if err != nil {
		return nil, err
	}

	var res ResponseInfo

	err = extract.IntoStructPtr(raw.Body, &res, respType)
	return &res, err
}
