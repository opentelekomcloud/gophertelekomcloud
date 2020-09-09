package client

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *golangsdk.ServiceClient {
	return &golangsdk.ServiceClient{
		ProviderClient: &golangsdk.ProviderClient{TokenID: TokenID},
		Endpoint:       testhelper.Endpoint(),
	}
}
