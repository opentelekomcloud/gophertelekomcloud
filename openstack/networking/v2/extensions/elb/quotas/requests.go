package quotas

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(c *golangsdk.ServiceClient) (r GetResult) {
	_, r.Err = c.Get(rootURL(c), &r.Body, nil)
	return
}
