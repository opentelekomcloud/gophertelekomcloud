package testing

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createClient() *golangsdk.ServiceClient {
	return &golangsdk.ServiceClient{
		ProviderClient: &golangsdk.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
