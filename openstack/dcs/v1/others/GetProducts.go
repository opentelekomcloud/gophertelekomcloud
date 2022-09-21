package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get products
func GetProducts(client *golangsdk.ServiceClient) (r GetResult) {
	// remove projectid from endpoint
	raw, err := client.Get(strings.Replace(client.ServiceURL("products"), "/"+client.ProjectID, "", -1), nil, nil)
	return
}
