package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

const resourcePath = "scaling_policy"

// createURL will build the rest query url of creation
// the create url is endpoint/scaling_policy
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func singleURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
