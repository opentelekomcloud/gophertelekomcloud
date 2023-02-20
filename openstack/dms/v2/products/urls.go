package products

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// endpoint/products
const resourcePath = "products"

func getURL(client *golangsdk.ServiceClient) string {
	url := strings.Split(client.Endpoint, "/v2/")[0]
	return url + "/v2/products"
}

func listURL(client *golangsdk.ServiceClient, engineType string) string {
	return client.ServiceURL(engineType, resourcePath)
}
