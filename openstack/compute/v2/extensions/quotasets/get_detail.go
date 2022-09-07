package quotasets

import "github.com/opentelekomcloud/gophertelekomcloud"

// GetDetail returns detailed public data about a previously created QuotaSet.
func GetDetail(client *golangsdk.ServiceClient, tenantID string) (r GetDetailResult) {
	raw, err := client.Get(client.ServiceURL("os-quota-sets", tenantID, "detail"), nil, nil)
	return
}
