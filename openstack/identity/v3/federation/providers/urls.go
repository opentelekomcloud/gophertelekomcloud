package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation"
)

const identityProvider = "identity_providers"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(federation.BaseURL, identityProvider)
}

func providerURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(federation.BaseURL, identityProvider, id)
}
