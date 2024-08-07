package products

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// endpoint/products
const resourcePath = "products"

func listURL(client *golangsdk.ServiceClient, engineType string) string {
	return client.ServiceURL(engineType, resourcePath)
}
