package response

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Gateway ID
	GatewayID string `json:"-"`
	// Group ID.
	GroupId string `json:"-"`
	// Group name, which can contain 1 to 64 characters, only letters, digits, hyphens (-) and
	// underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Response type definition, which includes a key and value. Options of the key:
	//     AUTH_FAILURE: Authentication failed.
	//     AUTH_HEADER_MISSING: The identity source is missing.
	//     AUTHORIZER_FAILURE: Custom authentication failed.
	//     AUTHORIZER_CONF_FAILURE: There has been a custom authorizer error.
	//     AUTHORIZER_IDENTITIES_FAILURE: The identity source of the custom authorizer is invalid.
	//     BACKEND_UNAVAILABLE: The backend service is unavailable.
	//     BACKEND_TIMEOUT: Communication with the backend service timed out.
	//     THROTTLED: The request was rejected due to request throttling.
	//     UNAUTHORIZED: The app you are using has not been authorized to call the API.
	//     ACCESS_DENIED: Access denied.
	//     NOT_FOUND: No API is found.
	//     REQUEST_PARAMETERS_FAILURE: The request parameters are incorrect.
	//     DEFAULT_4XX: Another 4XX error occurred.
	//     DEFAULT_5XX: Another 5XX error occurred.
	// Each error type is in JSON format.
	Responses map[string]ResponseInfo `json:"responses,omitempty"`
}

type ResponseInfo struct {
	// Response body template.
	Body string `json:"body" required:"true"`
	// HTTP status code of the response. If omitted, the status will be cancelled.
	Status int `json:"status,omitempty"`
	// Indicates whether the response is the default response.
	// Only the response of the API call is supported.
	IsDefault bool `json:"default,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Response, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupId, "gateway-responses"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res Response

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Response struct {
	// Response ID.
	ID string `json:"id"`
	// Response name.
	Name string `json:"name"`
	// Response type definition, which includes a key and value. Options of the key:
	//     AUTH_FAILURE: Authentication failed.
	//     AUTH_HEADER_MISSING: The identity source is missing.
	//     AUTHORIZER_FAILURE: Custom authentication failed.
	//     AUTHORIZER_CONF_FAILURE: There has been a custom authorizer error.
	//     AUTHORIZER_IDENTITIES_FAILURE: The identity source of the custom authorizer is invalid.
	//     BACKEND_UNAVAILABLE: The backend service is unavailable.
	//     BACKEND_TIMEOUT: Communication with the backend service timed out.
	//     THROTTLED: The request was rejected due to request throttling.
	//     UNAUTHORIZED: The app you are using has not been authorized to call the API.
	//     ACCESS_DENIED: Access denied.
	//     NOT_FOUND: No API is found.
	//     REQUEST_PARAMETERS_FAILURE: The request parameters are incorrect.
	//     DEFAULT_4XX: Another 4XX error occurred.
	//     DEFAULT_5XX: Another 5XX error occurred.
	// Each error type is in JSON format.
	Responses map[string]ResponseInfo `json:"responses"`
	// Indicates whether the group response is the default response.
	IsDefault bool `json:"default"`
	// Creation time.
	CreatedAt string `json:"create_time"`
	// Update time.
	UpdatedAt string `json:"update_time"`
}
