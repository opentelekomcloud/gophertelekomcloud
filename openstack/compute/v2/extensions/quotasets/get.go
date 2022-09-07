package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously created QuotaSet.
func Get(client *golangsdk.ServiceClient, tenantID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", tenantID), nil, nil)
	return
}
