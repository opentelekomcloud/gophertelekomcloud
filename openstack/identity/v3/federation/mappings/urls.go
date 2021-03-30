package mappings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation"
)

const rootPath = "mappings"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(federation.BaseURL, rootPath)
}

func mappingURL(client *golangsdk.ServiceClient, mappingID string) string {
	return client.ServiceURL(federation.BaseURL, rootPath, mappingID)
}
