package authorizer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID string `json:"-"`
	// Custom authorizer name.
	// It can contain 3 to 64 characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Custom authorizer type.
	// FRONTEND
	// BACKEND
	Type string `json:"type" required:"true"`
	// Enumeration values:
	// FUNC
	AuthorizerType string `json:"authorizer_type" required:"true"`
	// Function URN.
	FunctionUrn string `json:"authorizer_uri" required:"true"`
	// Function version.
	// If both a function alias URN and version are passed, the alias URN will be used and the version will be ignored.
	Version string `json:"authorizer_version,omitempty"`
	// Function alias URN.
	// If both a function alias URN and version are passed, the alias URN will be used and the version will be ignored.
	AliasUrn string `json:"authorizer_alias_uri,omitempty"`
	// Identity source.
	Identities []Identity `json:"identities,omitempty"`
	// Maximum cache age.
	Ttl *int `json:"ttl,omitempty"`
	// User data.
	UserData string `json:"user_data,omitempty"`
	// Custom backend ID.
	// Currently, this parameter is not supported.
	CustomBackendID string `json:"ld_api_id,omitempty"`
	// Indicates whether to send the body.
	NeedBody *bool `json:"need_body,omitempty"`
}

type Identity struct {
	// Parameter name.
	Name string `json:"name" required:"true"`
	// Parameter location.
	// Enumeration values:
	// HEADER
	// QUERY
	Location string `json:"location" required:"true"`
	// Parameter verification expression.
	// The default value is null, indicating that no verification is performed.
	Validation string `json:"validation,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*AuthorizerResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "authorizers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res AuthorizerResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AuthorizerResp struct {
	// Custom authorizer ID.
	ID string `json:"id"`
	// Custom authorizer name.
	Name string `json:"name"`
	// Custom authorizer type.
	Type string `json:"type"`
	// Authorizer type
	AuthorizerType string `json:"authorizer_type"`
	// Function URN.
	FunctionUrn string `json:"authorizer_uri"`
	// Function version.
	Version string `json:"authorizer_version"`
	// Function alias URN.
	AliasUrn string `json:"authorizer_alias_uri"`
	// Identity source.
	Identities []Identity `json:"identities"`
	// Maximum cache age.
	Ttl *int `json:"ttl"`
	// User data.
	UserData string `json:"user_data"`
	// Custom backend ID.
	CustomBackendID string `json:"ld_api_id"`
	// Indicates whether to send the body.
	NeedBody *bool `json:"need_body"`
	// Creation time.
	CreatedAt string `json:"create_time"`
	// ID of the application to which the custom authorizer belongs.
	AppID string `json:"roma_app_id"`
	// Name of the application to which the custom authorizer belongs.
	AppName string `json:"roma_app_name"`
}
