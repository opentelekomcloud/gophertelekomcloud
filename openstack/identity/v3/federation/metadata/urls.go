package metadata

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation"
)

func metadataURL(client *golangsdk.ServiceClient, provider, protocol string) string {
	return client.ServiceURL(
		federation.BaseURL,
		federation.ProvidersURL, provider,
		federation.ProtocolsURL, protocol,
		"metadata",
	)
}
