package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously created QuotaSet.
func Get(client *golangsdk.ServiceClient, tenantID string) (r GetResult) {
	raw, err := client.Get(getURL(client, tenantID), nil, nil)
	return
}
