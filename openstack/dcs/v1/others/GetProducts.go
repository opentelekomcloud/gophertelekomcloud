package others

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get products
func GetProducts(client *golangsdk.ServiceClient) (r GetResult) {
	// remove projectid from endpoint
	_, r.Err = client.Get(strings.Replace(client.ServiceURL("products"), "/"+client.ProjectID, "", -1), &r.Body, nil)
	return
}
